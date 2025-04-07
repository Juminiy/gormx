package clauses

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GroupByClause
// GROUP BY and HAVING
func GroupByClause(stmt *gorm.Statement) (groupByClause clause.GroupBy, ok bool) {
	groupBy, ok := util.MapElemOk(stmt.Clauses, GroupBy)
	if !ok {
		return
	}
	if groupByClause, ok = groupBy.Expression.(clause.GroupBy); ok {
		ok = len(groupByClause.Columns) > 0
	}
	return
}
