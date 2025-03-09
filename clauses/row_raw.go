package clauses

import (
	"github.com/Juminiy/gormx/callback"
	"gorm.io/gorm"
)

func (cfg *Config) RowRawClause(tx *gorm.DB) {
	if tx.Error != nil || callback.SkipRawRow.OK(tx) {
		return
	}

	if cfg.WriteClauseToRawOrRow {
		cfg.WhereClause(tx)
		if where, ok := WhereClause(tx); ok {
			_, _ = tx.Statement.WriteString(" WHERE ")
			where.Build(tx.Statement)
		}

		cfg.OrderByClause(tx)
		if orderBy, ok := OrderByClause(tx); ok {
			_ = tx.Statement.WriteByte(' ')
			orderBy.Build(tx.Statement)
		}

		if limit, ok := LimitClause(tx); ok {
			_ = tx.Statement.WriteByte(' ')
			limit.Build(tx.Statement)
		}
	}
}
