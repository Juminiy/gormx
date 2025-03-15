package gormx_testv2

import (
	"errors"
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm/utils/tests"
	"slices"
	"testing"
)

var QHideScope = gormx.Option{AfterQueryShowTenant: false}
var ErrQHideScope = errors.New("plugin after query hide scope error")
var QShowScope = gormx.Option{AfterQueryShowTenant: true}
var ErrQShowScope = errors.New("plugin after query show scope error")

func TestQueryStruct(t *testing.T) {
	t.Run("NoPlugin", func(tt *testing.T) {
		var order Order
		Err(tt, iSqlite0().First(&order))
	})
	t.Run("IsPlugin, NoScopes", func(tt *testing.T) {
		var order Order
		Err(tt, iSqlite().Set(gormx.OptionKey, QShowScope).First(&order))
	})
	t.Run("IsPlugin, Scopes(user_id)", func(tt *testing.T) {
		var order Order
		Err(tt, iSUser().Set(gormx.OptionKey, QShowScope).First(&order))
		if order.UserID == 0 {
			tt.Error(ErrQShowScope)
		}
	})
	t.Run("IsPlugin, Scopes(user_id), HideScope", func(tt *testing.T) {
		var order Order
		Err(tt, iSUser().Set(gormx.OptionKey, QHideScope).First(&order))
		tests.AssertEqual(tt, order.UserID, 0)
	})
	t.Run("IsPlugin, Scopes(user_id, tenant_id)", func(tt *testing.T) {
		var order Order
		Err(tt, iSUserTenant().Set(gormx.OptionKey, QShowScope).First(&order))
		if order.UserID == 0 || order.TenantID == 0 {
			tt.Error(ErrQShowScope)
		}
	})
	t.Run("IsPlugin, Scopes(user_id, tenant_id), HideScope", func(tt *testing.T) {
		var order Order
		Err(tt, iSUserTenant().Set(gormx.OptionKey, QHideScope).First(&order))
		tests.AssertEqual(tt, order.UserID, 0)
		tests.AssertEqual(tt, order.TenantID, 0)
	})
}

func TestQueryStructList(t *testing.T) {
	t.Run("NoPlugin", func(tt *testing.T) {
		var orders []Order
		Err(tt, iSqlite0().Limit(3).Find(&orders))
	})
	t.Run("IsPlugin, NoScopes", func(tt *testing.T) {
		var orders []Order
		Err(tt, iSqlite().Set(gormx.OptionKey, QShowScope).
			Limit(3).Find(&orders))
	})
	t.Run("IsPlugin, Scopes(user_id)", func(tt *testing.T) {
		var orders []Order
		Err(tt, iSUser().Set(gormx.OptionKey, QShowScope).
			Limit(3).Find(&orders))
		slices.Values(orders)(func(order Order) bool {
			if order.UserID == 0 {
				tt.Error(ErrQShowScope)
			}
			return true
		})
	})
	t.Run("IsPlugin, Scopes(user_id), HideScope", func(tt *testing.T) {
		var orders []Order
		Err(tt, iSUser().Set(gormx.OptionKey, QHideScope).
			Limit(3).Find(&orders))
		slices.Values(orders)(func(order Order) bool {
			tests.AssertEqual(tt, order.UserID, 0)
			return true
		})
	})
	t.Run("IsPlugin, Scopes(user_id, tenant_id)", func(tt *testing.T) {
		var orders []Order
		Err(tt, iSUserTenant().Set(gormx.OptionKey, QShowScope).
			Limit(3).Find(&orders))
		slices.Values(orders)(func(order Order) bool {
			if order.UserID == 0 || order.TenantID == 0 {
				tt.Error(ErrQShowScope)
			}
			return true
		})
	})
	t.Run("IsPlugin, Scopes(user_id, tenant_id), HideScope", func(tt *testing.T) {
		var orders []Order
		Err(tt, iSUserTenant().Set(gormx.OptionKey, QHideScope).
			Limit(3).Find(&orders))
		slices.Values(orders)(func(order Order) bool {
			tests.AssertEqual(tt, order.UserID, 0)
			tests.AssertEqual(tt, order.TenantID, 0)
			return true
		})
	})
}

func TestQueryMapType(t *testing.T) {
	// panic: assignment to entry in nil map
	/*t.Run("NoPlugin Map", func(tt *testing.T) {
		var order map[string]any
		Err(tt, iSqlite0().Model(&Order{}).First(order))
		t.Log(Enc(order))
	})*/
	t.Run("NoPlugin *Map", func(tt *testing.T) {
		var order map[string]any
		Err(tt, iSqlite0().Model(&Order{}).First(&order))
		tt.Log(Enc(order))
	})
	// error: sql: Scan called without calling Next
	/*t.Run("NoPlugin **Map", func(tt *testing.T) {
		var order = util.New(make(map[string]any))
		Err(tt, iSqlite0().Model(&Order{}).First(&order))
		tt.Log(Enc(order))
	})*/
	// panic: reflect: reflect.Value.Set using unaddressable value
	/*t.Run("NoPlugin MapList", func(tt *testing.T) {
		var orders []map[string]any
		Err(tt, iSqlite0().Model(&Order{}).Limit(3).Find(orders))
		t.Log(Enc(orders))
	})*/
	t.Run("NoPlugin *MapList", func(tt *testing.T) {
		var orders []map[string]any
		Err(tt, iSqlite0().Model(&Order{}).Limit(3).Find(&orders))
		tt.Log(Enc(orders))
	})
	// nil sliceMap: reflect: call of reflect.Value.Field on map Value
	// not nil sliceMap sql: Scan called without calling Next
	/*t.Run("NoPlugin **MapList", func(tt *testing.T) {
		var orders = util.New(make([]map[string]any, 0, 8))
		Err(tt, iSqlite0().Model(&Order{}).Limit(3).Find(&orders))
		tt.Log(Enc(orders))
	})*/
}

func TestQueryMap(t *testing.T) {
	t.Run("NoPlugin", func(tt *testing.T) {
		var order map[string]any
		Err(tt, iSqlite0().Model(&Order{}).First(&order))
	})
	t.Run("IsPlugin NoScopes", func(tt *testing.T) {
		var order map[string]any
		Err(tt, iSqlite().Model(&Order{}).First(&order))
	})
	t.Run("IsPlugin Scopes(user_id), ShowScope", func(tt *testing.T) {
		var order map[string]any
		Err(tt, iSUser().Set(gormx.OptionKey, QShowScope).Model(&Order{}).First(&order))
		if userID, uOk := util.MapElemOk(order, "user_id"); uOk {
			if uID, uOk := userID.(uint); uOk && uID != 0 {
				return
			}
		}
		tt.Error(ErrQShowScope)
	})
	t.Run("IsPlugin Scopes(user_id), HideScope", func(tt *testing.T) {
		var order map[string]any
		Err(tt, iSUser().Set(gormx.OptionKey, QHideScope).Model(&Order{}).First(&order))
		if util.MapOk(order, "user_id") {
			tt.Error(ErrQHideScope)
		}
	})
	t.Run("IsPlugin Scopes(user_id, tenant_id), ShowScope", func(tt *testing.T) {
		var order map[string]any
		Err(tt, iSUserTenant().Set(gormx.OptionKey, QShowScope).Model(&Order{}).First(&order))
		if userID, uOk := util.MapElemOk(order, "user_id"); uOk {
			if uID, uOk := userID.(uint); uOk && uID != 0 {
				if tenantID, tOk := util.MapElemOk(order, "tenant_id"); tOk {
					if tID, tOk := tenantID.(uint); tOk && tID != 0 {
						return
					}
				}
			}
		}
		tt.Error(ErrQShowScope)
	})
	t.Run("IsPlugin Scopes(user_id, tenant_id), HideScope", func(tt *testing.T) {
		var order map[string]any
		Err(tt, iSUserTenant().Set(gormx.OptionKey, QHideScope).Model(&Order{}).First(&order))
		if util.MapOk(order, "user_id") || util.MapOk(order, "tenant_id") {
			tt.Error(ErrQHideScope)
		}
	})
}

func TestQueryMapList(t *testing.T) {
	t.Run("NoPlugin", func(tt *testing.T) {
		var orders []map[string]any
		Err(tt, iSqlite0().Model(&Order{}).Limit(3).Find(&orders))

	})
	t.Run("IsPlugin NoScopes", func(tt *testing.T) {
		var orders []map[string]any
		Err(tt, iSqlite().Model(&Order{}).Limit(3).Find(&orders))

	})
	t.Run("IsPlugin Scopes(user_id), ShowScope", func(tt *testing.T) {
		var orders []map[string]any
		Err(tt, iSUser().Set(gormx.OptionKey, QShowScope).Model(&Order{}).First(&orders))
		slices.Values(orders)(func(order map[string]any) bool {
			if userID, uOk := util.MapElemOk(order, "user_id"); uOk {
				if uID, uOk := userID.(uint); uOk && uID != 0 {
					return true
				}
			}
			tt.Error(ErrQShowScope)
			return false
		})

	})
	t.Run("IsPlugin Scopes(user_id), HideScope", func(tt *testing.T) {
		var orders []map[string]any
		Err(tt, iSUser().Set(gormx.OptionKey, QHideScope).Model(&Order{}).First(&orders))
		slices.Values(orders)(func(order map[string]any) bool {
			if util.MapOk(order, "user_id") {
				tt.Error(ErrQHideScope)
			}
			return true
		})
	})
	t.Run("IsPlugin Scopes(user_id, tenant_id), ShowScope", func(tt *testing.T) {
		var orders []map[string]any
		Err(tt, iSUserTenant().Set(gormx.OptionKey, QShowScope).Model(&Order{}).First(&orders))
		slices.Values(orders)(func(order map[string]any) bool {
			if userID, uOk := util.MapElemOk(order, "user_id"); uOk {
				if uID, uOk := userID.(uint); uOk && uID != 0 {
					if tenantID, tOk := util.MapElemOk(order, "tenant_id"); tOk {
						if tID, tOk := tenantID.(uint); tOk && tID != 0 {
							return true
						}
					}
				}
			}
			tt.Error(ErrQShowScope)
			return false
		})
	})
	t.Run("IsPlugin Scopes(user_id, tenant_id), HideScope", func(tt *testing.T) {
		var orders []map[string]any
		Err(tt, iSUserTenant().Set(gormx.OptionKey, QHideScope).Model(&Order{}).First(&orders))
		slices.Values(orders)(func(order map[string]any) bool {
			if util.MapOk(order, "user_id") || util.MapOk(order, "tenant_id") {
				tt.Error(ErrQHideScope)
			}
			return true
		})
	})
}
