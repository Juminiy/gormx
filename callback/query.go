package callback

import (
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/schema"
	"slices"
)

func BeforeQueryOmit(tx *gorm.DB) {
	// replaced by gorm tag `->:false` is an AfterQuery Set Fields To Zero
	// QueryOmit is Omit query column
	if sch := tx.Statement.Schema; sch != nil {
		slices.All(sch.Fields)(func(_ int, field *schema.Field) bool {
			if !field.Readable {
				tx.Statement.Omit(field.DBName)
			}
			return true
		})
	}
}

func AfterFindMapCallHook(db *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(db); ok &&
		!db.Statement.SkipHooks && sch.AfterFind {
		/*setUpDestMapStmtModel(db, sch)*/
		CallHooks(db, func(v any, tx *gorm.DB) bool {
			if afterFindI, ok := v.(callbacks.AfterFindInterface); ok {
				_ = db.AddError(afterFindI.AfterFind(tx))
				return true
			}
			return false
		})
		scanModelToDestMap(db)
	}
}
