package callback

import (
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"maps"
	"reflect"
	"slices"
)

func BeforeUpdateMapDeletePkAndSetPkToClause(tx *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(tx); ok {
		mapValue := deps.IndI(tx.Statement.Dest).MapValues()
		slices.All(sch.PrimaryFields)(func(_ int, pF *schema.Field) bool {
			if mapElem, ok := util.MapElemOk(mapValue, pF.DBName); ok {
				mapElemRv := reflect.ValueOf(mapElem)
				if mapElemRv.IsValid() && !mapElemRv.IsZero() {
					tx.Statement.AddClause(clauses.ClauseFieldEq(pF, mapElem))
				}
				deps.IndI(tx.Statement.Dest).MapSetField(map[string]any{pF.DBName: nil})
			}
			return true
		})
	}
}

func BeforeUpdateMapDeleteUnknownColumn(tx *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(tx); ok {
		dbNames := sets.New(sch.DBNames...)
		notInDBNames := make([]string, 0, len(dbNames)/4)
		switch mapValue := tx.Statement.Dest.(type) {
		case map[string]any:
			maps.All(mapValue)(func(dbName string, _ any) bool {
				if !dbNames.Has(dbName) {
					notInDBNames = append(notInDBNames, dbName)
				}
				return true
			})
			slices.All(notInDBNames)(func(_ int, dbName string) bool {
				delete(mapValue, dbName)
				return true
			})

		case *map[string]any:
			maps.All(*mapValue)(func(dbName string, _ any) bool {
				if !dbNames.Has(dbName) {
					notInDBNames = append(notInDBNames, dbName)
				}
				return true
			})
			slices.All(notInDBNames)(func(_ int, dbName string) bool {
				delete(*mapValue, dbName)
				return true
			})

		default: // ignore
		}
	}
}

func BeforeUpdateMapDeleteZeroValueColumn(tx *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(tx); ok {
		mapValue := deps.IndI(tx.Statement.Dest).MapValues()
		slices.All(sch.Fields)(func(_ int, field *schema.Field) bool {
			if mapElem, ok := util.MapElemOk(mapValue, field.DBName); ok {
				mapElemRv := reflect.ValueOf(mapElem)
				if mapElemRv.IsValid() && mapElemRv.IsZero() {
					deps.IndI(tx.Statement.Dest).MapSetField(map[string]any{field.DBName: nil})
				}
			}
			return true
		})
	}
}

// no need to call Model for Hooks,
// gorm will do: callbacks.SetupUpdateReflectValue
// we only need to do:
//  1. db.Statement.Model and set before
//  2. set Config.BeforeUpdate before callbacks.SetupUpdateReflectValue
//
// detail in: Config.Initialize
// referred from: callbacks.BeforeUpdate
func BeforeUpdateMapCallHook(db *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(db); ok &&
		!db.Statement.SkipHooks && sch.BeforeUpdate {
		setUpDestMapStmtModel(db, sch)
		/*CallHooks(db, func(v any, tx *gorm.DB) bool {
			if beforeUpdateI, ok := v.(callbacks.BeforeUpdateInterface); ok {
				_ = db.AddError(beforeUpdateI.BeforeUpdate(tx))
				return true
			}
			return false
		})*/
	}
}

// no need to call Model for Hooks, gorm will do: callbacks.SetupUpdateReflectValue
// referred from: callbacks.AfterUpdate
func AfterUpdateMapCallHook(db *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(db); ok &&
		!db.Statement.SkipHooks && sch.AfterUpdate {
		/*CallHooks(db, func(v any, tx *gorm.DB) bool {
			if afterUpdateI, ok := v.(callbacks.AfterUpdateInterface); ok {
				_ = db.AddError(afterUpdateI.AfterUpdate(tx))
				return true
			}
			return false
		})*/
	}
}

func BeforeUpdateGetClausePk(modelRv reflect.Value, stmt *gorm.Statement) (clausePk clause.Expression, oks bool) {
	clauseList := make([]clause.Interface, 0, 4)
	if sch := stmt.Schema; sch != nil {
		if modelRv.IsValid() && modelRv.Kind() == reflect.Struct {
			slices.All(sch.PrimaryFields)(func(_ int, field *schema.Field) bool {
				if value, isZero := field.ValueOf(stmt.Context, modelRv); !isZero {
					clauseList = append(clauseList, clauses.ClauseFieldEq(field, value))
				}
				return true
			})
		} else if modelRv.Kind() == reflect.Map {
			mapValue := deps.Ind(stmt.ReflectValue).MapValues()
			slices.All(sch.PrimaryFields)(func(_ int, pF *schema.Field) bool {
				if mapElem, ok := util.MapElemOk(mapValue, pF.DBName); ok {
					mapElemRv := reflect.ValueOf(mapElem)
					if mapElemRv.IsValid() && !mapElemRv.IsZero() {
						clauseList = append(clauseList, clauses.ClauseFieldEq(pF, mapElem))
					}
				}
				return true
			})
		}
	}
	if len(clauseList) > 0 {
		return clause.And(lo.Map(clauseList, func(item clause.Interface, _ int) clause.Expression {
			return clause.Expression(item)
		})...), true
	}
	return clausePk, false
}
