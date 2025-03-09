package callback

import "gorm.io/gorm"

type _Cfg struct {
	key string
}

func (c *_Cfg) Set(tx *gorm.DB) *gorm.DB {
	return tx.Set(c.key, struct{}{})
}

func (c *_Cfg) OK(tx *gorm.DB) bool {
	_, ok := tx.Get(c.key)
	return ok
}

var SkipCreate = _Cfg{
	key: "internal:skip_create_callback",
}
var SkipDelete = _Cfg{
	key: "internal:skip_delete_callback",
}
var SkipUpdate = _Cfg{
	key: "internal:skip_update_callback",
}
var SkipQuery = _Cfg{
	key: "internal:skip_query_callback",
}
var SkipRawRow = _Cfg{
	key: "internal:skip_raw_row_callback",
}
