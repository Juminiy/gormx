package clauseslite

import "gorm.io/gorm/clause"

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
