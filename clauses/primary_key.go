package clauses

import (
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"reflect"
	"slices"
)

func StmtPrimaryKeyClause(stmt *gorm.Statement) (clauseI clause.Expression, ok bool) {
	if clauseI, ok = stmtPrimaryKeyClause(stmt); ok {
		return
	} else if clauseI, ok = destKindIsMapPrimaryKeyClause(stmt); ok {
		return
	} else if clauseI, ok = modelKindIsStructPrimaryKeyClause(stmt); ok {
		return
	} else if clauseI, ok = destKindIsStructPrimaryKeyClause(stmt); ok {
		return
	}
	return
}

func StmtHasPrimaryKeyNotZero(stmt *gorm.Statement) bool {
	switch stmt.ReflectValue.Kind() {
	case reflect.Map:
		return destKindIsMapAndHasPrimaryKeyNotZero(stmt) ||
			modelKindIsStructAndHasPrimaryKeyNotZero(stmt) ||
			stmtHasPrimaryKeyNotZero(stmt)

	case reflect.Struct, reflect.Array, reflect.Slice:
		return destKindIsStructAndHasPrimaryKeyNotZero(stmt) ||
			modelKindIsStructAndHasPrimaryKeyNotZero(stmt) ||
			stmtHasPrimaryKeyNotZero(stmt)

	default:
		return false
	}
}

func destKindIsMapAndHasPrimaryKeyNotZero(stmt *gorm.Statement) bool {
	_, ok := destKindIsMapPrimaryKeyClause(stmt)
	return ok
}

func destKindIsMapPrimaryKeyClause(stmt *gorm.Statement) (clauseI clause.Expression, ok bool) {
	if stmt.SQL.Len() == 0 && stmt.Schema != nil {
		if mapRv := deps.IndI(stmt.Dest); mapRv.Value.Kind() == reflect.Map {
			mapValue := mapRv.MapValues()
			var columns []string
			var values []any
			for _, pF := range stmt.Schema.PrimaryFields {
				if mapElem, ok := util.MapElemOk(mapValue, pF.DBName); ok && !deps.ItemValueIsZero(mapElem) {
					columns = append(columns, pF.DBName)
					values = append(values, mapElem)
				}
			}
			if len(columns) > 0 {
				return clause.IN{Column: columns, Values: []any{values}}, true
			}
		}
	}
	return
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
			if column, values, ok = filterZeroValuesColumn(column, values); ok {
				return clause.IN{Column: column, Values: values}, true
			}
		}

		if stmt.ReflectValue.CanAddr() && stmt.Dest != stmt.Model && stmt.Model != nil {
			_, queryValues = schema.GetIdentityFieldValuesMap(stmt.Context, deps.IndI(stmt.Model).Value, stmt.Schema.PrimaryFields)
			column, values = schema.ToQueryValues(stmt.Table, stmt.Schema.PrimaryFieldDBNames, queryValues)
			if len(values) > 0 {
				if column, values, ok = filterZeroValuesColumn(column, values); ok {
					return clause.IN{Column: column, Values: values}, true
				}
			}
		}
	}
	return
}

func filterZeroValuesColumn(column any, values []any) (any, []any, bool) {
	if col, ok := column.([]clause.Column); ok {
		if vals, ok := values[0].([]any); ok {
			var colOf []clause.Column
			var valOf []any
			slices.All(vals)(func(idx int, val any) bool {
				if !deps.ItemValueIsZero(val) {
					colOf = append(colOf, col[idx])
					valOf = append(valOf, val)
				}
				return true
			})
			if len(colOf) > 0 {
				return colOf, []any{valOf}, true
			}
		}
	}
	return nil, nil, false
}

// referred from: callbacks.BuildQuerySQL
func destKindIsStructAndHasPrimaryKeyNotZero(stmt *gorm.Statement) bool {
	_, ok := destKindIsStructPrimaryKeyClause(stmt)
	return ok
}

func destKindIsStructPrimaryKeyClause(stmt *gorm.Statement) (clauseI clause.Expression, ok bool) {
	return kindIsStructPrimaryKeyClause(stmt, stmt.Dest)
}

func modelKindIsStructAndHasPrimaryKeyNotZero(stmt *gorm.Statement) bool {
	_, ok := modelKindIsStructPrimaryKeyClause(stmt)
	return ok
}

func modelKindIsStructPrimaryKeyClause(stmt *gorm.Statement) (clauseI clause.Expression, ok bool) {
	return kindIsStructPrimaryKeyClause(stmt, stmt.Model)
}

func kindIsStructPrimaryKeyClause(stmt *gorm.Statement, stmtValue any) (clauseI clause.Expression, ok bool) {
	if stmt.SQL.Len() == 0 {
		if modelRv := deps.IndI(stmtValue); modelRv.T.Kind() == reflect.Struct &&
			modelRv.Type == stmt.Schema.ModelType {
			var columns []string
			var values []any
			for _, primaryField := range stmt.Schema.PrimaryFields {
				if val, isZero := primaryField.ValueOf(stmt.Context, modelRv.Value); !isZero {
					columns = append(columns, primaryField.DBName)
					values = append(values, val)
				}
			}
			if len(columns) > 0 {
				return clause.IN{Column: columns, Values: []any{values}}, true
			}
		}
	}
	return nil, false
}
