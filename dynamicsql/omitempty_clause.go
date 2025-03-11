package dynamicsql

import (
	"github.com/Juminiy/gormx/clauses"
	"gorm.io/gorm"
)

// TODO: fix
// OmitEmptyClause
// more to expansion
func OmitEmptyClause(tx *gorm.DB) {
	clauses.ModifyWhereClause(tx, clauses.NotZeroValueExpr)
	clauses.ModifyOrderByClause(tx, clauses.KnownColumn)
}
