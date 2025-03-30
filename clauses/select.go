package clauses

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SelectClause
// SELECT columns
func SelectClause(tx *gorm.DB) (selectClause clause.Select, ok bool) {
	slctC, sok := util.MapElemOk(tx.Statement.Clauses, Select)
	if !sok {
		return
	}
	if selectClause, ok = slctC.Expression.(clause.Select); ok {
		ok = len(selectClause.Columns) > 0
	}
	return
}
