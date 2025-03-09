package uniques

import (
	"github.com/Juminiy/gormx/clauses/clauseslite"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"k8s.io/apimachinery/pkg/util/sets"
	"maps"
	"reflect"
	"slices"
	"strings"
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
	} /*else if len(d.ColumnValue) == 0 {
		d.ColumnValue = lo.MapKeys(d.FieldValue, func(_ any, name string) string {
			return d.FieldColumn[name]
		})
	}*/

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
	orExpr = clauseslite.FalseExpr()
	noExpr = true
	slices.All(lo.MapToSlice(d.Groups, func(_ string, names []string) clause.Expression {
		var andExpr clause.Expression = clauseslite.TrueExpr()
		slices.All(names)(func(_ int, name string) bool {
			fieldValue, ok := d.FieldValue[name]
			if !ok || deps.ItemValueIsZero(fieldValue) {
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
	sliceMap := rval.Values()

	if fdErr := d.listCheck(sliceMap); fdErr != nil {
		_ = tx.AddError(fdErr)
		return
	}

	d.List = lo.Map(sliceMap, func(item map[string]any, _ int) rowValues {
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
	orExpr = clauseslite.FalseExpr()
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

func (d *rowsValues) listCheck(sliceMap []map[string]any) (fdErr error) {
	groupSets := lo.MapValues(d.Groups, func(_ []string, groupName string) sets.Set[string] {
		return sets.New[string]()
	})

	setErr := func(fields []string, values []any) {
		fdErr = fieldDupListCheckErr{
			fieldDupCountErr: fieldDupCountErr{
				dbTable: d.DBTable,
				dbName: lo.Map(fields, func(field string, _ int) string {
					return d.FieldColumn[field]
				}),
			},
			dupValues: values,
		}
	}

	accStr := func(fields []string, mapValues map[string]any) (string, []any, bool) {
		isComplete, values, buf := true, make([]any, 0, len(fields)), strings.Builder{}
		slices.Values(fields)(func(field string) bool {
			if val, ok := util.MapElemOk(mapValues, field); ok {
				if strVal := cast.ToString(val); !deps.ItemValueIsZero(val) && len(strVal) > 0 {
					buf.WriteString(strVal)
					values = append(values, val)
					return true
				}
			}
			isComplete = false
			return false
		})
		return buf.String(), values, isComplete
	}

	slices.Values(sliceMap)(func(mapValues map[string]any) bool {
		if fdErr != nil {
			return false
		}
		maps.All(d.Groups)(func(groupName string, fields []string) bool {
			if bufStr, values, ok := accStr(fields, mapValues); ok {
				if groupSets[groupName].Has(bufStr) {
					setErr(fields, values)
					return false
				} else {
					groupSets[groupName].Insert(bufStr)
				}
			}
			return true
		})
		return true
	})

	return
}
