package deps

import (
	"database/sql"
	"github.com/Juminiy/kube/pkg/util"
	rv "github.com/Juminiy/kube/pkg/util/safe_reflect"
	rv2 "github.com/Juminiy/kube/pkg/util/safe_reflect/v2"
	rv3 "github.com/Juminiy/kube/pkg/util/safe_reflect/v3"
	"github.com/samber/lo"
	"reflect"
	"time"
)

func Ind(rvalue reflect.Value) rv3.Tv {
	return rv3.WrapI(rvalue)
}

func Dir(rvalue reflect.Value) rv3.Tv { return rv3.Wrap(rvalue) }

func IndI(i any) rv3.Tv {
	return rv3.Indirect(i)
}

func DirI(i any) rv3.Tv {
	return rv3.Direct(i)
}

func IndISet(i any) rv2.Value {
	return rv2.Indirect(i)
}

func Tag(s string) rv3.Tag {
	return rv3.ParseTagValue(s)
}

func AS(i any) []any {
	return rv3.ToAnySlice(i)
}

func Comp(rtype reflect.Type) bool { return rv.CanDirectCompare(rtype) }

func ItemValueIsZero(i any) bool {
	itemRvalue := IndI(i)
	if !itemRvalue.IsValid() { // nil
		return true
	} else if itemRvKind := itemRvalue.Type.Kind(); itemRvKind == reflect.Bool { // bool: not judge
		return false
	} else if Comp(itemRvalue.Type) { // int,uint,string,float
		return itemRvalue.IsZero()
	} else if util.ElemIn(itemRvKind, reflect.Array, reflect.Slice, reflect.Map) { // len = 0 is zero
		return itemRvalue.Value.Len() == 0
	} else if itemRvKind == reflect.Struct && itemRvalue.Value.NumField() == 0 { // empty struct is zero
		return true
	} else if itemRvKind == reflect.Struct && IsSqlNullType(itemRvalue.Type) { // sql.* Type: for special judge
		return !itemRvalue.Value.FieldByName("Valid").Bool()
	} else if itemRvKind == reflect.Struct && IsStdTimeType(itemRvalue.Type) { // time.Time Type: for special judge
		rets, called := itemRvalue.CallMethod("IsZero", nil)
		if called && len(rets) == 1 {
			if retBool, assetOk := rets[0].(bool); retBool && assetOk {
				return true
			}
		}
	}
	return false
}

func IsStdTimeType(rt reflect.Type) bool {
	return rt == _timeTimeTypeRType || rt.ConvertibleTo(_timeTimeTypeRType)
}

func IsSqlNullType(rt reflect.Type) bool {
	_, ok := lo.Find(_sqlNullTypeRType, func(item reflect.Type) bool {
		if rt == item || rt.ConvertibleTo(item) {
			return true
		}
		return false
	})
	return ok
}

var _timeTimeTypeRType = reflect.TypeOf((*time.Time)(nil)).Elem()

var _sqlNullTypeRType = []reflect.Type{
	reflect.TypeOf((*sql.NullString)(nil)).Elem(),
	reflect.TypeOf((*sql.NullInt64)(nil)).Elem(),
	reflect.TypeOf((*sql.NullInt32)(nil)).Elem(),
	reflect.TypeOf((*sql.NullInt16)(nil)).Elem(),
	reflect.TypeOf((*sql.NullByte)(nil)).Elem(),
	reflect.TypeOf((*sql.NullFloat64)(nil)).Elem(),
	reflect.TypeOf((*sql.NullBool)(nil)).Elem(),
	reflect.TypeOf((*sql.NullTime)(nil)).Elem(),
}
