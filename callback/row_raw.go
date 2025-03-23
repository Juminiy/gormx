package callback

import (
	"github.com/Juminiy/gormx/clauses"
	"gorm.io/gorm"
)

// WriteToRowOrRaw
// effect on raw and row callbacks
// will append
// clause.Where, clause.OrderBy, clause.Limit
// to sql that before gorm/finisher_api.go
func WriteToRowOrRaw(tx *gorm.DB) {
	if where, ok := clauses.ModifyWhereClause(tx, clauses.LegalExpr); ok {
		_, _ = tx.Statement.WriteString(" WHERE ")
		where.Build(tx.Statement)
	}

	if orderBy, ok := clauses.ModifyOrderByClause(tx, clauses.LegalOrderByColumn); ok {
		_, _ = tx.Statement.WriteString(" ORDER BY ")
		orderBy.Build(tx.Statement)
	}

	if limit, ok := clauses.LimitClause(tx); ok {
		_ = tx.Statement.WriteByte(' ')
		limit.Build(tx.Statement)
	}
}
