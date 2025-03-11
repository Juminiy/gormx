package tenants

import (
	"github.com/Juminiy/gormx/clauses"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (t *Tenant) AddClause(tx *gorm.DB) {
	tx.Statement.AddClause(t)
}

func (t *Tenant) Name() string { return "Tenant" }

func (t *Tenant) Build(_ clause.Builder) {}

func (t *Tenant) MergeClause(_ *clause.Clause) {}

// referred from: gorm.SoftDeleteQueryClause
func (t *Tenant) ModifyStatement(stmt *gorm.Statement) {
	if c, ok := stmt.Clauses[clauses.Where]; ok {
		if where, ok := c.Expression.(clause.Where); ok && len(where.Exprs) >= 1 {
			for _, expr := range where.Exprs {
				if orCond, ok := expr.(clause.OrConditions); ok && len(orCond.Exprs) == 1 {
					where.Exprs = []clause.Expression{clause.And(where.Exprs...)}
					c.Expression = where
					stmt.Clauses[clauses.Where] = c
					break
				}
			}
		}
	}

	stmt.AddClause(clause.Where{Exprs: []clause.Expression{
		t.Field.Clause(),
	}})
}
