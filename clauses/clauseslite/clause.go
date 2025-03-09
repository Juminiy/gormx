package clauseslite

import (
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

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
