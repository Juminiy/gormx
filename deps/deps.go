package deps

import (
	rv "github.com/Juminiy/kube/pkg/util/safe_reflect"
	rv2 "github.com/Juminiy/kube/pkg/util/safe_reflect/v2"
	rv3 "github.com/Juminiy/kube/pkg/util/safe_reflect/v3"
	"reflect"
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
	return !itemRvalue.IsValid() ||
		(itemRvalue.IsValid() && Comp(itemRvalue.Type) && itemRvalue.IsZero())
}
