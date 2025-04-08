package coverindex

import (
	"fmt"
	"github.com/Juminiy/gormx/clauses"
	"gorm.io/gorm"
)

type DelayJoin struct {
	OffsetUpperBound int64
}

// ModifyStatement
// DelayJoin impl gorm.StatementModifier
//
// Modify srcSQL:
/*
	SELECT <cols> FROM tbl_xxx WHERE <clauses>
	ORDER BY <orders> LIMIT <limit> OFFSET <offset>
*/
// that SQL offset >= DelayJoin.OffsetUpperBound
//
// To dstSQL:
/*
	SELECT <cols> FROM tbl_xxx
	INNER JOIN (
		SELECT <pk_cols> FROM tbl_xxx WHERE <clauses>
		ORDER BY <orders> LIMIT <limit> OFFSET <offset>
	) tbl_xxx_pk USING (pk)
*/
func (d DelayJoin) ModifyStatement(stmt *gorm.Statement) {
	sch := stmt.Schema
	if d.OffsetUpperBound == 0 || sch == nil {
		return
	}

	orderBy, hasOrderBy := clauses.OrderByClause(stmt)
	if !hasOrderBy {
		return
	}
	limit, hasLimitOffset := clauses.LimitClause(stmt)
	if !hasLimitOffset || int64(limit.Offset) < d.OffsetUpperBound {
		return
	}

	queryPk := stmt.DB.Session(&gorm.Session{NewDB: true, SkipDefaultTransaction: true}).
		Model(stmt.Model).
		Select(sch.PrimaryFieldDBNames).
		Clauses(orderBy, limit)
	where, hasWhere := clauses.WhereClause(stmt)
	if hasWhere {
		queryPk.Clauses(where)
		delete(stmt.Clauses, clauses.Where)
	}
	delete(stmt.Clauses, clauses.OrderBy)
	delete(stmt.Clauses, clauses.Limit)

	/*clauseJoin := clause.Join{
		Type: clause.InnerJoin,
		Table: clause.Table{
			Name:  stmt.DB.ToSQL(func(db *gorm.DB) *gorm.DB { return queryPk }),
			Alias: sch.Table + "_pk",
			Raw:   true,
		},
		Using: sch.PrimaryFieldDBNames,
	}*/
	stmt.DB.Joins(
		fmt.Sprintf("INNER JOIN (?) AS %s_pk USING %s",
			sch.Table, stmt.Quote(sch.PrimaryFieldDBNames)),
		queryPk,
	)

}
