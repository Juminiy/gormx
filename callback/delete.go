package callback

import (
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
)

func DoQueryBeforeDelete(tx *gorm.DB) {
	ntx := tx.Session(&gorm.Session{NewDB: true})

	/*ntx = _SkipQueryCallback.Set(ntx)*/

	if schema := tx.Statement.Schema; schema != nil {
		slices.All(schema.QueryClauses)(func(_ int, c clause.Interface) bool {
			ntx.Statement.AddClause(c)
			return true
		})
	}

	if txClause, ok := clauses.WhereClause(tx); ok {
		ntx.Where(txClause)
	}

	if returning, ok := util.MapElemOk(tx.Statement.Clauses, "RETURNING"); ok {
		if returningClause, ok := returning.Expression.(clause.Returning); ok {
			slices.All(returningClause.Columns)(func(_ int, column clause.Column) bool {
				ntx.Statement.Selects = append(ntx.Statement.Selects, column.Name)
				return true
			})
		}
	} else if len(tx.Statement.Selects) != 0 {
		ntx.Statement.Selects = tx.Statement.Selects
	}

	err := ntx.Find(tx.Statement.Dest).Error
	if err != nil {
		tx.Logger.Error(tx.Statement.Context, "before delete, do query, error: %s", err.Error())
	}
}
