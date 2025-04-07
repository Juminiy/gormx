package clauses

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// LimitClause
// LIMIT and OFFSET
func LimitClause(stmt *gorm.Statement) (limitClause clause.Limit, ok bool) {
	limit, ook := util.MapElemOk(stmt.Clauses, Limit)
	if !ook {
		return
	}
	if limitClause, ok = limit.Expression.(clause.Limit); ok {
		ok = limitClause.Limit != nil && *limitClause.Limit > 0
	}
	return
}
