package schemas

/*
gorm support Dest Type
1. Struct
 * reflect.Kind -> T: Struct,Int,UInt,Float
 * Struct  	   ->	--(indirect)--> T
 * SliceStruct ->	--(indirect)--> []T, []*...*T
 * ArrayStruct ->	--(indirect)--> [N]T, [N]*...*T

2. Map
 * Map -> map[string]any, *map[string]any
 * SliceMap -> []map[string]any, *map[string]any
*/
