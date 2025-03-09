package schemas

import (
	"github.com/Juminiy/gormx/clauses/clauseslite"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type Field struct {
	Name    string
	DBTable string
	DBName  string
	Value   any
	Values  []any
}

func FieldFromSchema(field *schema.Field) Field {
	return Field{
		Name:    field.Name,
		DBTable: field.Schema.Table,
		DBName:  field.DBName,
	}
}

func (f Field) Clause() clause.Expression {
	var expr clause.Expression = clauseslite.TrueExpr()
	if f.Value != nil {
		expr = f.ClauseEq()
	} else if len(f.Values) > 0 {
		expr = f.ClauseIn()
	}
	return expr
}

func (f Field) ClauseEq() clause.Eq {
	return clause.Eq{
		Column: clause.Column{
			Table: f.DBTable,
			Name:  f.DBName,
		},
		Value: f.Value,
	}
}

func (f Field) ClauseIn() clause.IN {
	return clause.IN{
		Column: clause.Column{
			Table: f.DBTable,
			Name:  f.DBName,
		},
		Values: f.Values,
	}
}
