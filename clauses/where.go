package clauses

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"slices"
)

func ModifyWhereClause(tx *gorm.DB, exprOk func(clause.Expression) bool) (where clause.Where, ok bool) {
	where, ok = WhereClause(tx)
	if !ok {
		return
	}

	exprIList := make([]clause.Expression, 0, len(where.Exprs))
	slices.All(where.Exprs)(func(_ int, exprI clause.Expression) bool {
		if exprOk(exprI) {
			exprIList = append(exprIList, exprI)
		}
		return true
	})
	whereClause := tx.Statement.Clauses[Where]
	where.Exprs = exprIList
	whereClause.Expression = where
	tx.Statement.Clauses[Where] = whereClause
	return where, ok
}

// WhereClause
// Expr or ExprList
func WhereClause(tx *gorm.DB) (whereClause clause.Where, ok bool) {
	where, wok := util.MapElemOk(tx.Statement.Clauses, Where)
	if !wok {
		return
	}
	if whereClause, ok = where.Expression.(clause.Where); ok {
		ok = len(whereClause.Exprs) > 0
	}
	return
}

func NoWhereClause(tx *gorm.DB) bool {
	_, ok := WhereClause(tx)
	return !ok && !StmtHasPrimaryKeyNotZero(tx.Statement)
}
