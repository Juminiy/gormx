package optlock

import (
	"github.com/Juminiy/gormx/clauses"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
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
		if verField.AutoUpdateTime == 0 {
			tx.Statement.AddClause(clause.Assignments(map[string]any{
				verField.DBName: verValue,
			}))
		}
		return true // optimistic lock is success
	}
	return false
}
