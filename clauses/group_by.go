package clauses

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GroupByClause
// GROUP BY and HAVING
func GroupByClause(tx *gorm.DB) (groupByClause clause.GroupBy, ok bool) {
	groupBy, ok := util.MapElemOk(tx.Statement.Clauses, GroupBy)
	if !ok {
		return
	}
	if groupByClause, ok = groupBy.Expression.(clause.GroupBy); ok {
		ok = len(groupByClause.Columns) > 0
	}
	return
}
