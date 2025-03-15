package clauses

import (
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"reflect"
)

func StmtGetPrimaryKeyNotZeroClause(stmt *gorm.Statement) (clauseI clause.Expression, ok bool) {
	return stmtPrimaryKeyClause(stmt)
}

func destHasPrimaryKeyNotZero(stmt *gorm.Statement) bool {
	switch stmt.ReflectValue.Kind() {
	case reflect.Map:
		return destKindIsMapAndHasPrimaryKeyNotZero(stmt) ||
			stmtHasPrimaryKeyNotZero(stmt)

	case reflect.Struct, reflect.Array, reflect.Slice:
		return stmtHasPrimaryKeyNotZero(stmt)

	default:
		return false
	}
}

func destKindIsMapAndHasPrimaryKeyNotZero(stmt *gorm.Statement) bool {
	if stmt.SQL.Len() == 0 && stmt.Schema != nil {
		if mapRv := deps.IndI(stmt.Dest); mapRv.Value.Kind() == reflect.Map {
			mapValue := mapRv.MapValues()
			for _, pF := range stmt.Schema.PrimaryFields {
				if mapElem, ok := util.MapElemOk(mapValue, pF.DBName); ok && !deps.ItemValueIsZero(mapElem) {
					return true
				}
			}
		}
	}
	return false
}

func stmtHasPrimaryKeyNotZero(stmt *gorm.Statement) bool {
	_, ok := stmtPrimaryKeyClause(stmt)
	return ok
}

// referred from: callbacks.Delete, callbacks.Update
func stmtPrimaryKeyClause(stmt *gorm.Statement) (clauseI clause.Expression, ok bool) {
	if stmt.SQL.Len() == 0 && stmt.Schema != nil {
		_, queryValues := schema.GetIdentityFieldValuesMap(stmt.Context, stmt.ReflectValue, stmt.Schema.PrimaryFields)
		column, values := schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)
		if len(values) > 0 {
			return clause.IN{Column: column, Values: values}, true
		}

		if /*stmt.ReflectValue.CanAddr() && */ stmt.Dest != stmt.Model && stmt.Model != nil {
			_, queryValues = schema.GetIdentityFieldValuesMap(stmt.Context, reflect.ValueOf(stmt.Model), stmt.Schema.PrimaryFields)
			column, values = schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)
			if len(values) > 0 {
				return clause.IN{Column: column, Values: values}, true
			}
		}
	}
	return
}

// referred from: callbacks.BuildQuerySQL
func destKindIsStructAndHasPrimaryKeyNotZero(stmt *gorm.Statement) bool {
	if stmt.SQL.Len() == 0 {
		if stmt.ReflectValue.Kind() == reflect.Struct &&
			stmt.ReflectValue.Type() == stmt.Schema.ModelType {
			for _, primaryField := range stmt.Schema.PrimaryFields {
				if _, isZero := primaryField.ValueOf(stmt.Context, stmt.ReflectValue); !isZero {
					return true
				}
			}
		}
	}
	return false
}
