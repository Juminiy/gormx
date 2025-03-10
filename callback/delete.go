package callback

import (
	"github.com/Juminiy/gormx/clauses"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
)

// TODO: returning before delete
func BeforeDeleteReturning(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	ntx := tx.Session(&gorm.Session{NewDB: true})

	//ntx = callback.SkipQuery.Set(ntx)

	if schema := tx.Statement.Schema; schema != nil {
		slices.All(schema.QueryClauses)(func(_ int, c clause.Interface) bool {
			ntx.Statement.AddClause(c)
			return true
		})
	}

	if txClause, ok := clauses.WhereClause(tx); ok {
		ntx.Statement.AddClause(txClause)
	}

	if returning, ok := clauses.ReturningClause(tx); ok && len(returning.Columns) > 0 {
		slices.All(returning.Columns)(func(_ int, column clause.Column) bool {
			ntx.Statement.Selects = append(ntx.Statement.Selects, column.Name)
			return true
		})
	} else if len(tx.Statement.Selects) != 0 {
		ntx.Statement.Selects = tx.Statement.Selects
	}

	err := ntx.Find(tx.Statement.Dest).Error
	if err != nil {
		ntx.Logger.Error(ntx.Statement.Context, "before delete, do query, error: %s", err.Error())
	}
}
