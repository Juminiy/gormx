package gormx_testv2

import (
	"testing"
)

// not a bug
func TestBugFixPluginCreateUniques(t *testing.T) {
	order := RandomOrder()
	Err(t, iSqlite().Create(&order))
	// Because CreateHooks After Plugin FieldDupCount, Order.Serial = ""
	t.Log(order.JSONString())
}
