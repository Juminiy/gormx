package clauses

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ReturningClause(tx *gorm.DB) (returningClause clause.Returning, ok bool) {
	returning, rok := util.MapElemOk(tx.Statement.Clauses, Returning)
	if !rok {
		return
	}
	returningClause, ok = returning.Expression.(clause.Returning)
	return
}
