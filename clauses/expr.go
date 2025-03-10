package clauses

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"strings"
)

func LegalExpr(exprI clause.Expression) bool {
	return legalExpr(exprI)
}

// implicit recursive
func legalExpr(exprI clause.Expression) bool {
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
		return !util.AssertZero(exprV.Column) &&
			!util.AssertZero(exprV.Value)

	case clause.Expr:
		return len(exprV.SQL) != 0 &&
			strings.Count(exprV.SQL, "?") == len(exprV.Vars)

	case clause.NamedExpr:
		return len(exprV.SQL) != 0 &&
			strings.Count(exprV.SQL, "?") == len(exprV.Vars)

	case clause.AndConditions:
		return legalExprList(exprV.Exprs)

	case clause.OrConditions:
		return legalExprList(exprV.Exprs)

	case clause.NotConditions:
		return legalExprList(exprV.Exprs)

	default:
		return true
	}
}

func legalExprList(exprIList []clause.Expression) bool {
	return len(exprIList) > 0 &&
		lo.CountBy(exprIList, func(item clause.Expression) bool {
			return legalExpr(item)
		}) == len(exprIList)
}

func NotZeroValueExpr(exprI clause.Expression) bool {
	return notZeroValueExpr(exprI)
}

// implicit recursive
func notZeroValueExpr(exprI clause.Expression) bool {
	switch exprV := exprI.(type) {
	case clause.Eq:
		return !util.AssertZero(exprV.Column) &&
			!util.AssertZero(exprV.Value)
	case clause.Neq:
		return !util.AssertZero(exprV.Column) &&
			!util.AssertZero(exprV.Value)
	case clause.Gt:
		return !util.AssertZero(exprV.Column) &&
			!util.AssertZero(exprV.Value)
	case clause.Gte:
		return !util.AssertZero(exprV.Column) &&
			!util.AssertZero(exprV.Value)
	case clause.Lt:
		return !util.AssertZero(exprV.Column) &&
			!util.AssertZero(exprV.Value)
	case clause.Lte:
		return !util.AssertZero(exprV.Column) &&
			!util.AssertZero(exprV.Value)
	case clause.Like:
		return !util.AssertZero(exprV.Column) &&
			!util.AssertZero(exprV.Value)

	case clause.Expr:
		return len(exprV.SQL) != 0 &&
			strings.Count(exprV.SQL, "?") == len(exprV.Vars)

	case clause.NamedExpr:
		return len(exprV.SQL) != 0 &&
			strings.Count(exprV.SQL, "?") == len(exprV.Vars)

	case clause.AndConditions:
		return notZeroValueExprList(exprV.Exprs)

	case clause.OrConditions:
		return notZeroValueExprList(exprV.Exprs)

	case clause.NotConditions:
		return notZeroValueExprList(exprV.Exprs)

	default:
		return true
	}
}

func notZeroValueExprList(exprIList []clause.Expression) bool {
	return len(exprIList) > 0 &&
		lo.CountBy(exprIList, func(item clause.Expression) bool {
			return notZeroValueExpr(item)
		}) == len(exprIList)
}

func TrueExpr() clause.NamedExpr {
	return clause.NamedExpr{
		SQL: "1=1",
	}
}

func FalseExpr() clause.NamedExpr {
	return clause.NamedExpr{
		SQL: "1!=1",
	}
}

func ClauseFieldEq(field *schema.Field, value any) clause.Interface {
	return clause.Where{Exprs: []clause.Expression{
		clause.Eq{
			Column: clause.Column{
				Table: field.Schema.Table,
				Name:  field.DBName,
			},
			Value: value,
		},
	}}
}

func ClauseColumnEq(column string, value any) clause.Interface {
	return clause.Where{Exprs: []clause.Expression{
		clause.Eq{
			Column: column,
			Value:  value,
		},
	}}
}
