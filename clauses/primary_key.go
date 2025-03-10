package clauses

import (
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"reflect"
)

// referred from: callbacks.BuildQuerySQL
// has at least one primaryKey value is not zero
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
