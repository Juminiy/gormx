package coverindex

import (
	"github.com/Juminiy/gormx/clauses"
	"gorm.io/gorm"
)

type DelayJoin struct {
	OffsetUpperBound int64
}

// ModifyStatement
// TODO: DelayJoin to impl gorm.StatementModifier
//
// Modify srcSQL:
/*
	SELECT <cols> FROM tbl_xxx WHERE (clauses)
	ORDER BY <orders> LIMIT <limit> OFFSET <offset>
*/
// that SQL offset >= DelayJoin.OffsetUpperBound
//
// To dstSQL:
/*
	SELECT <cols> FROM tbl_xxx
	INNER JOIN (
		SELECT <pk_cols> FROM tbl_xxx WHERE (clauses)
		ORDER BY <orders> LIMIT <limit> OFFSET <offset>
	) tbl_xxx_pk
	ON tbl_xxx.pk0 = tbl_xxx_pk.pk0 AND tbl_xxx.pk1 = tbl_xxx_pk.pk1 AND ...
*/
func (d DelayJoin) ModifyStatement(stmt *gorm.Statement) {
	sch := stmt.Schema
	if sch == nil {
		return
	}
	orderBy, ok := clauses.OrderByClause(stmt)
	if !ok {
		return
	}
	limit, ok := clauses.LimitClause(stmt)
	if !ok || int64(limit.Offset) < d.OffsetUpperBound {
		return
	}
	where, ok := clauses.WhereClause(stmt)
	if ok {
		delete(stmt.Clauses, clauses.Where)
	}

	stmt.InnerJoins(
		sch.Table+"_pk",
		stmt.DB.Model(stmt.Model).Select(sch.PrimaryFieldDBNames).
			Where(where).Order(orderBy).Limit(*limit.Limit).Offset(limit.Offset))
}
