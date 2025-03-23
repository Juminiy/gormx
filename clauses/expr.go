package clauses

import (
	"github.com/Juminiy/gormx/deps"
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
		return !LegalColumn(exprV.Column)
	case clause.Neq:
		return !LegalColumn(exprV.Column)
	case clause.Gt:
		return !LegalColumn(exprV.Column)
	case clause.Gte:
		return !LegalColumn(exprV.Column)
	case clause.Lt:
		return !LegalColumn(exprV.Column)
	case clause.Lte:
		return !LegalColumn(exprV.Column)
	case clause.Like:
		return !LegalColumn(exprV.Column) &&
			!deps.ItemValueIsZero(exprV.Value)

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
		return !LegalColumn(exprV.Column) &&
			!deps.ItemValueIsZero(exprV.Value)
	case clause.Neq:
		return !LegalColumn(exprV.Column) &&
			!deps.ItemValueIsZero(exprV.Value)
	case clause.Gt:
		return !LegalColumn(exprV.Column) &&
			!deps.ItemValueIsZero(exprV.Value)
	case clause.Gte:
		return !LegalColumn(exprV.Column) &&
			!deps.ItemValueIsZero(exprV.Value)
	case clause.Lt:
		return !LegalColumn(exprV.Column) &&
			!deps.ItemValueIsZero(exprV.Value)
	case clause.Lte:
		return !LegalColumn(exprV.Column) &&
			!deps.ItemValueIsZero(exprV.Value)
	case clause.Like:
		return !LegalColumn(exprV.Column) &&
			!deps.ItemValueIsZero(exprV.Value)

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

// LegalColumn
// referred to clause.Statement.QuoteTo
func LegalColumn(column any) bool {
	switch columnV := column.(type) {
	case string:
		return len(columnV) > 0
	case []string:
		_, columnLen0 := lo.Find(columnV, func(item string) bool {
			return len(item) == 0
		})
		return !columnLen0
	case clause.Column:
		return len(columnV.Name) > 0
	case []clause.Column:
		_, columnLen0 := lo.Find(columnV, func(item clause.Column) bool {
			return len(item.Name) == 0
		})
		return !columnLen0
	default:
		return false
	}
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
