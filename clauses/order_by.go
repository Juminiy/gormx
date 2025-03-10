package clauses

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
)

func ModifyOrderByClause(tx *gorm.DB) (orderBy clause.OrderBy, ok bool) {
	orderBy, ok = OrderByClause(tx)
	if !ok {
		return
	}

	columns := make([]clause.OrderByColumn, 0, len(orderBy.Columns))
	slices.All(orderBy.Columns)(func(_ int, column clause.OrderByColumn) bool {
		if len(column.Column.Name) > 0 {
			columns = append(columns, column)
		}
		return true
	})
	orderClause := tx.Statement.Clauses[OrderBy]
	orderBy.Columns = columns
	orderClause.Expression = orderBy
	tx.Statement.Clauses[OrderBy] = orderClause
	return orderBy, ok
}

// OrderByClause
// ORDER BY column or ORDER BY columnList
func OrderByClause(tx *gorm.DB) (orderByClause clause.OrderBy, ok bool) {
	orderBy, ook := util.MapElemOk(tx.Statement.Clauses, OrderBy)
	if !ook {
		return
	}
	if orderByClause, ok = orderBy.Expression.(clause.OrderBy); ok {
		ok = len(orderByClause.Columns) > 0
	}
	return
}

func omitOrderByClauseEmptyOrNotKnownColumn(tx *gorm.DB) {}
