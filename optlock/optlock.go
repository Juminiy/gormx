package optlock

import (
	"fmt"
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
	"slices"
)

type Config struct {
	Name       string // default: gormx:optlock
	TagKey     string // default: gormx
	TagOptLock string // default: version
}

func (cfg *Config) OptimisticLock(tx *gorm.DB) bool {
	// when update not in transaction
	if sch := tx.Statement.Schema; tx.SkipDefaultTransaction && sch != nil {
		// field has updated_at or tag gormx:"version"
		var verField, byUpdateTime, byVerTag *schema.Field
		slices.Values(sch.Fields)(func(field *schema.Field) bool {
			if field.AutoUpdateTime > 0 {
				byUpdateTime = field
			} else if tag, ok := field.Tag.Lookup(cfg.TagKey); ok && tag == cfg.TagOptLock {
				byVerTag = field
			}
			return true
		})
		if byUpdateTime == nil && byVerTag == nil {
			return false
		} else if byUpdateTime != nil {
			verField = byUpdateTime
		} else if !util.ElemIn(byVerTag.DataType, schema.Int, schema.Uint) {
			return false
		} else {
			verField = byVerTag
		}

		// must update by primaryKey to determine one row update
		// primaryKey value must in
		// 1. .Updates(Struct) 			-> Dest  Struct PkValue
		// 2. .Updates(Map)    			-> Dest  Map 	PkValue
		// 3. .Model(Struct).Updates() 	-> Model Struct PkValue
		// 4. .Model(Struct).Update() 	-> Model Struct PkValue
		// !!!will not check of where clause Pk!!!
		clausePk, clauseCan := clauses.StmtPrimaryKeyClause(tx.Statement)
		if !clauseCan {
			return false
		}

		var verValue = verField.NewValuePool.Get()
		defer verField.NewValuePool.Put(verValue)
		err := tx.Session(&gorm.Session{NewDB: true}).
			Table(sch.Table).
			Where(clausePk).
			Pluck(verField.DBName, &verValue).Error
		if err != nil {
			tx.Logger.Error(tx.Statement.Context, "optimistic lock query version field value error: %s", err.Error())
			return false
		}

		// origin tx addClause
		tx.Statement.AddClause(clauses.ClauseFieldEq(verField, verValue))
		if verField.AutoUpdateTime == 0 { // not updateTime; must be int or uint
			switch tx.Statement.ReflectValue.Kind() {
			case reflect.Map:
				deps.Ind(tx.Statement.ReflectValue).MapSetField(map[string]any{
					verField.DBName: gorm.Expr(fmt.Sprintf("%s + 1", verField.DBName)),
				})

			case reflect.Struct:
				if verField.DataType == schema.Int {
					deps.Ind(tx.Statement.ReflectValue).StructSet(map[string]any{
						verField.Name: cast.ToInt(verValue) + 1,
					})
				} else if verField.DataType == schema.Uint {
					deps.Ind(tx.Statement.ReflectValue).StructSet(map[string]any{
						verField.Name: cast.ToUint(verValue) + 1,
					})
				}

			default: // ignore
				return false
			}
		}
		return true // optimistic lock is success
	}
	return false
}
