package clauses

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
)

func ModifyOrderByClause(tx *gorm.DB, columnOk func(clause.OrderByColumn) bool) (orderBy clause.OrderBy, ok bool) {
	orderBy, ok = OrderByClause(tx.Statement)
	if !ok {
		return
	}

	columns := make([]clause.OrderByColumn, 0, len(orderBy.Columns))
	slices.All(orderBy.Columns)(func(_ int, column clause.OrderByColumn) bool {
		if columnOk(column) {
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
func OrderByClause(stmt *gorm.Statement) (orderByClause clause.OrderBy, ok bool) {
	orderBy, ook := util.MapElemOk(stmt.Clauses, OrderBy)
	if !ook {
		return
	}
	if orderByClause, ok = orderBy.Expression.(clause.OrderBy); ok {
		ok = len(orderByClause.Columns) > 0
	}
	return
}

func LegalOrderByColumn(column clause.OrderByColumn) bool {
	return len(column.Column.Name) > 0
}

func KnownOrderByColumn(column clause.OrderByColumn) bool {
	return len(column.Column.Name) > 0
}
