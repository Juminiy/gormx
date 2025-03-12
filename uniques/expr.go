package uniques

import (
	"github.com/Juminiy/gormx/clauses"
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

// rowValuesExpr
// support one group, multiple groups
// each group one field, multiple fields
// each group if one or more fields reflect.Value.IsZero(), the group will be omitted
// if no groups, the count will be omitted
type rowValues struct {
	FieldValue  map[string]any
	ColumnValue map[string]any
	*FieldDup
	ForUpdate bool
}

func (d *rowValues) simple(tx *gorm.DB) {
	if len(d.FieldValue) == 0 && len(d.ColumnValue) == 0 {
		return
	} else if len(d.FieldValue) == 0 {
		d.FieldValue = lo.MapKeys(d.ColumnValue, func(_ any, column string) string {
			return d.ColumnField[column]
		})
	}

	if exprI, ok := d.expr(); ok {
		d.doCount(tx, exprI, d.ForUpdate)
	}
}

func (d *rowValues) expr() (exprI clause.Expression, ok bool) {
	return d.exprIn()
}

func (d *rowValues) exprIn() (exprIn clause.Expression, ok bool) {
	exprIn, ok = clauses.FalseExpr(), false
	slices.Values(lo.MapToSlice(d.Groups, func(_ string, names []string) clause.Expression {
		if inExprValid(names, d.FieldValue) {
			return clause.IN{
				Column: lo.Map(names, func(name string, _ int) string { return d.FieldColumn[name] }),
				Values: []any{lo.Map(names, func(name string, _ int) any { return d.FieldValue[name] })},
			}
		}
		return nil
	}))(func(inExpr clause.Expression) bool {
		if inExpr != nil {
			exprIn, ok = clause.Or(exprIn, inExpr), true
		}
		return true
	})
	return exprIn, ok
}

func inExprValid(fields []string, fieldValues map[string]any) bool {
	if _, noFieldOrValZero := lo.Find(fields, func(field string) bool {
		fieldVal, hasField := fieldValues[field]
		return !hasField || deps.ItemValueIsZero(fieldVal)
	}); !noFieldOrValZero {
		return true
	}
	return false
}

func (d *rowValues) exprOr() (orExpr clause.Expression, orOk bool) {
	orExpr, orOk = clauses.FalseExpr(), false
	slices.All(lo.MapToSlice(d.Groups, func(_ string, names []string) clause.Expression {
		var andExpr clause.Expression = clauses.TrueExpr()
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
	}))(func(_ int, andExpr clause.Expression) bool {
		if andExpr == nil {
			return true
		}
		orExpr, orOk = clause.Or(orExpr, andExpr), true
		return true
	})
	return
}

type rowsValues struct {
	Value []map[string]any
	//List  []rowValues
	*FieldDup
}

func (d *rowsValues) complex(tx *gorm.DB) {
	rval := deps.Ind(tx.Statement.ReflectValue)
	if !util.ElemIn(rval.T.Indirect().Kind(), reflect.Struct, reflect.Map) {
		return
	}
	d.Value = rval.Values()

	if fdErr := d.listCheck(); fdErr != nil {
		_ = tx.AddError(fdErr)
		return
	}

	/*d.List = lo.Map(sliceMap, func(item map[string]any, _ int) rowValues {
		return rowValues{
			FieldValue: item,
			FieldDup:   d.FieldDup,
		}
	})*/
	if exprI, ok := d.expr(); ok {
		d.doCount(tx, exprI, false)
	}
}

func (d *rowsValues) expr() (exprI clause.Expression, ok bool) {
	return d.exprIn()
}

func (d *rowsValues) exprIn() (exprIn clause.Expression, ok bool) {
	exprIn, ok = clauses.FalseExpr(), false
	slices.Values(lo.MapToSlice(d.Groups, func(_ string, names []string) clause.Expression {
		return clause.IN{
			Column: lo.Map(names, func(name string, _ int) string { return d.FieldColumn[name] }),
			Values: lo.FilterMap(d.Value, func(mapValue map[string]any, _ int) (any, bool) {
				groupValues := lo.Map(names, func(name string, _ int) any { return mapValue[name] })
				if inExprValidV2(groupValues) {
					return groupValues, true
				}
				return nil, false
			}),
		}
	}))(func(inExpr clause.Expression) bool {
		exprIn, ok = clause.Or(exprIn, inExpr), true
		return true
	})
	return
}

func inExprValidV2(sliceValues []any) bool {
	if _, valZero := lo.Find(sliceValues, func(value any) bool {
		return deps.ItemValueIsZero(value)
	}); !valZero {
		return true
	}
	return false
}

/*func (d *rowsValues) exprOr() (orExpr clause.Expression, orOk bool) {
	orExpr, orOk = clauses.FalseExpr(), false
	slices.All(d.List)(func(_ int, values rowValues) bool {
		subOrExpr, subOrOk := values.expr()
		if !subOrOk {
			return true
		}
		orExpr, orOk = clause.Or(orExpr, subOrExpr), true
		return true
	})
	return
}*/

func (d *rowsValues) listCheck() (fdErr error) {
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

	slices.Values(d.Value)(func(mapValues map[string]any) bool {
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
