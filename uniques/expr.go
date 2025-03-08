package uniques

import (
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"slices"
)

type rowValues struct {
	FieldValue  map[string]any
	ColumnValue map[string]any
	*FieldDup
	ForUpdate bool
}

func (d *rowValues) simple(tx *gorm.DB) {
	if len(d.Groups) == 0 {
		return
	}
	if len(d.FieldValue) == 0 &&
		len(d.ColumnValue) == 0 {
		return
	} else if len(d.FieldValue) == 0 {
		d.FieldValue = lo.MapKeys(d.ColumnValue, func(_ any, column string) string {
			return d.ColumnField[column]
		})
	} else if len(d.ColumnValue) == 0 {
		d.ColumnValue = lo.MapKeys(d.FieldValue, func(_ any, name string) string {
			return d.FieldColumn[name]
		})
	}

	orExpr, noExpr := d.expr()
	if noExpr {
		return
	}
	d.doCount(tx, orExpr, d.ForUpdate)
}

// rowValuesExpr
// support one group, multiple groups
// each group one field, multiple fields
// each group if one or more fields reflect.Value.IsZero(), the group will be omitted
// if no groups, the count will be omitted
func (d *rowValues) expr() (orExpr clause.Expression, noExpr bool) {
	orExpr = clauses.FalseExpr()
	noExpr = true
	slices.All(lo.MapToSlice(d.Groups, func(_ string, names []string) clause.Expression {
		var andExpr clause.Expression = clauses.TrueExpr()
		slices.All(names)(func(_ int, name string) bool {
			fieldValue, ok := d.FieldValue[name]
			if !ok || deps.IndI(fieldValue).Value.IsZero() {
				andExpr = nil
				return false
			}
			andExpr = clause.And(andExpr, clause.Eq{
				Column: d.FieldColumn[name],
				Value:  fieldValue,
			})
			return true
		})
		return andExpr
	}))(func(_ int, expression clause.Expression) bool {
		if expression == nil {
			return true
		}
		noExpr = false
		orExpr = clause.Or(orExpr, expression)
		return true
	})
	return
}

type rowsValues struct {
	List []rowValues
	*FieldDup
}

func (d *rowsValues) complex(tx *gorm.DB) {
	if len(d.Groups) == 0 {
		return
	}
	rval := deps.Ind(tx.Statement.ReflectValue)
	if !util.ElemIn(rval.T.Indirect().Kind(), reflect.Struct, reflect.Map) {
		return
	}
	d.List = lo.Map(rval.Values(), func(item map[string]any, _ int) rowValues {
		return rowValues{
			FieldValue: item,
			FieldDup:   d.FieldDup,
		}
	})

	orExpr, noExpr := d.expr()
	if noExpr {
		return
	}
	d.doCount(tx, orExpr, false)
}

func (d *rowsValues) expr() (orExpr clause.Expression, noExpr bool) {
	orExpr = clauses.FalseExpr()
	noExpr = true
	slices.All(d.List)(func(_ int, values rowValues) bool {
		subOrExpr, noOK := values.expr()
		if noOK {
			return true
		}
		noExpr = false
		orExpr = clause.Or(orExpr, subOrExpr)
		return true
	})
	return
}
