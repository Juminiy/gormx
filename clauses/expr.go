package clauses

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
	"strings"
)

func checkExprI(exprI clause.Expression) bool {
	switch exprV := exprI.(type) {
	case clause.Eq:
		return !util.AssertZero(exprV.Column)
	case clause.Neq:
		return !util.AssertZero(exprV.Column)
	case clause.Gt:
		return !util.AssertZero(exprV.Column)
	case clause.Gte:
		return !util.AssertZero(exprV.Column)
	case clause.Lt:
		return !util.AssertZero(exprV.Column)
	case clause.Lte:
		return !util.AssertZero(exprV.Column)
	case clause.Like:
		if util.AssertZero(exprV.Column) ||
			util.AssertZero(exprV.Value) {
			return false
		} /*else if strings.HasPrefix(cast.ToString(exprV.Value), "%") {

		}*/

	case clause.Expr:
		if len(exprV.SQL) == 0 ||
			strings.Count(exprV.SQL, "?") != len(exprV.Vars) {
			return false
		}

	case clause.NamedExpr:
		if len(exprV.SQL) == 0 ||
			strings.Count(exprV.SQL, "?") != len(exprV.Vars) {
			return false
		}

	case clause.AndConditions:
		return checkExprIList(exprV.Exprs)

	case clause.OrConditions:
		return checkExprIList(exprV.Exprs)

	case clause.NotConditions:
		return checkExprIList(exprV.Exprs)

	default:

	}
	return true
}

func checkExprIList(exprIList []clause.Expression) bool {
	return lo.CountBy(exprIList, func(item clause.Expression) bool {
		return checkExprI(item)
	}) == len(exprIList)
}
