package schemas

/*
gorm support Dest Type
1. Struct
 * reflect.Kind -> T: Struct,Int,UInt,Float,Bool,Byte,...
 * Struct  	   ->	--(indirect)--> T
 * SliceStruct ->	--(indirect)--> []T, []*T
	> internal side at most one level pointer to: []T or []*T
	> outer side no restricted levels pointer to: *[]T,*[]*T, **...***[]T,**...***[]*T
 * ArrayStruct ->	--(indirect)--> [N]T, [N]*T

2. Map
 * Map -> map[string]any, *map[string]any
 * SliceMap -> *[]map[string]any; []map[string]any(Test Bugs)
*/
