package gormx_testv2

import (
	"errors"
	"github.com/Juminiy/gormx"
	"gorm.io/gorm/utils/tests"
	"slices"
	"testing"
)

var CHideScope = gormx.Option{AfterCreateShowTenant: false, EnableComplexFieldDup: true}
var ErrCHideScope = errors.New("plugin after create hide scope error")
var CShowScope = gormx.Option{AfterCreateShowTenant: true, EnableComplexFieldDup: true}
var ErrCShowScope = errors.New("plugin after create show scope error")

func TestCreateStruct(t *testing.T) {
	t.Run("CreateStruct NoPlugin", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSqlite0().Create(&order))
		tt.Log(order.JSONString())
	})
	t.Run("CreateStruct IsPlugin: Uniques, NoScopes", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSqlite().Create(&order))
		tt.Log(order.JSONString())
	})
	t.Run("CreateStruct IsPlugin: Uniques, 1Scopes(user_id)", func(tt *testing.T) {
		order0, order1 := RandomOrder(), RandomOrder()
		Err(tt, iSUser().Set(gormx.OptionKey, CHideScope).Create(&order0)) // hide one
		Err(tt, iSUser().Set(gormx.OptionKey, CShowScope).Create(&order1)) // show one
		if order0.UserID != 0 {
			tt.Error(ErrCHideScope)
		} /*else if order1.UserID == 0 {
			tt.Error(ErrCShowScope)
		}*/
	})
	t.Run("CreateStruct IsPlugin: Uniques, 2Scopes(user_id, tenant_id)", func(tt *testing.T) {
		order0, order1 := RandomOrder(), RandomOrder()
		Err(tt, iSUserTenant().Set(gormx.OptionKey, CHideScope).Create(&order0)) // hide two
		Err(tt, iSUserTenant().Set(gormx.OptionKey, CShowScope).Create(&order1)) // show two
		if order0.UserID != 0 || order0.TenantID != 0 {
			tt.Error(ErrCHideScope)
		} else if /*order1.UserID == 0 ||*/ order1.TenantID == 0 {
			tt.Error(ErrCShowScope)
		}
	})
}

func TestCreateStructList(t *testing.T) {
	t.Run("CreateStructList NoPlugin", func(tt *testing.T) {
		orders := []*Order{RandomOrder(), RandomOrder(), RandomOrder()}
		Err(tt, iSqlite0().Create(&orders))
	})
	t.Run("CreateStructList IsPlugin: NoUniques, NoScopes", func(tt *testing.T) {
		orders := []*Order{RandomOrder(), RandomOrder(), RandomOrder()}
		Err(tt, iSqlite().Create(&orders)) // no list field dup check
		slices.Values(orders)(func(order *Order) bool {
			if order.UserID != 0 || order.TenantID != 0 {
				tt.Error(errors.New("plugin not set scopes, but scopes effect"))
				return false
			}
			return true
		})
	})
	t.Run("CreateStructList IsPlugin: Uniques, 1Scopes(user_id)", func(tt *testing.T) {
		orders := []*Order{RandomOrder(), RandomOrder(), RandomOrder()}
		Err(tt, iSUser().Set(gormx.OptionKey, CHideScope).Create(&orders))
		slices.Values(orders)(func(order *Order) bool {
			if order.UserID != 0 || order.TenantID != 0 {
				tt.Error(ErrCHideScope)
				return false
			}
			return true
		})
	})
	t.Run("CreateStructList IsPlugin: Uniques, 2Scopes(user_id, tenant_id)", func(tt *testing.T) {
		orders := []*Order{RandomOrder(), RandomOrder(), RandomOrder()}
		Err(tt, iSUserTenant().Set(gormx.OptionKey, CShowScope).Create(&orders))
		slices.Values(orders)(func(order *Order) bool {
			if /*order.UserID == 0 ||*/ order.TenantID == 0 {
				tt.Error(ErrCShowScope)
				return false
			}
			return true
		})
	})

}

var CMapHideScope = gormx.Option{
	AfterCreateShowTenant: false, EnableComplexFieldDup: true,
}

var CMapShowScope = gormx.Option{
	AfterCreateShowTenant: true, EnableComplexFieldDup: true,
}

var CMapBeforeCallHook = gormx.Option{
	AfterCreateShowTenant: true, EnableComplexFieldDup: true,
	BeforeCreateMapCallHooks: true, AfterCreateMapCallHooks: false,
}

var CMapAfterCallHook = gormx.Option{
	AfterCreateShowTenant: true, EnableComplexFieldDup: true,
	BeforeCreateMapCallHooks: true, AfterCreateMapCallHooks: true,
}

func TestCreateMap(t *testing.T) {
	t.Run("CreateMap NoPlugin", func(tt *testing.T) {
		order := RandomOrderMap()
		Err(tt, iSqlite0().Model(&Order{}).Create(&order)) // raw gorm create map: 4-fields
		tt.Log(Enc(order))
		tests.AssertEqual(t, len(order), 5) // 4 + 1(Pk)
	})
	t.Run("CreateMap IsPlugin: Uniques, NoScopes", func(tt *testing.T) {
		order := RandomOrderMap()
		Err(tt, iSqlite().Table(`tbl_order`).Create(&order)) // plugins: before set default values, after set Pk
		tt.Log(Enc(order))
		tests.AssertEqual(t, len(order), 7) // 4 + 1(Pk) + 2(AutoTime)
	})
	t.Run("CreateMap IsPlugin: Uniques, 1Scopes(user_id)", func(tt *testing.T) {
		order := RandomOrderMap()
		Err(tt, iSUser().Set(gormx.OptionKey, CMapHideScope).
			Table(`tbl_order`).Create(&order))
		tt.Log(Enc(order))
		tests.AssertEqual(t, len(order), 7) // 4 + 1(Pk) + 2(AutoTime) + 0(1HideScope)
	})
	t.Run("CreateMap IsPlugin: Uniques, 2Scopes(user_id, tenant_id)", func(tt *testing.T) {
		order := RandomOrderMap()
		Err(tt, iSUserTenant().Set(gormx.OptionKey, CMapShowScope).
			Table(`tbl_order`).Create(&order))
		tt.Log(Enc(order))
		tests.AssertEqual(t, len(order), 9) // 4 + 1(Pk) + 2(AutoTime) + 2(ShowScope)
	})
	t.Run("CreateMap IsPlugin: Uniques, 2Scopes(user_id, tenant_id); Reinforced: before call_hooks", func(tt *testing.T) {
		order := RandomOrderMap()
		Err(tt, iSUserTenant().Set(gormx.OptionKey, CMapBeforeCallHook).
			Table(`tbl_order`).Create(&order)) // before call_hooks
		tt.Log(Enc(order))
		tests.AssertEqual(t, len(order), 12) // 4 + 1(Pk) + 2(AutoTime) + 2(ShowScope) + 3(BeforeCallHooks)
	})
	t.Run("CreateMap IsPlugin: Uniques, 2Scopes(user_id, tenant_id); Reinforced: before and after call_hooks", func(tt *testing.T) {
		order := RandomOrderMap()
		Err(tt, iSUserTenant().Set(gormx.OptionKey, CMapAfterCallHook).
			Table(`tbl_order`).Create(&order)) // after call_hooks
		tt.Log(Enc(order))
		tests.AssertEqual(t, len(order), 11) // 4 + 1(Pk) + 2(AutoTime) + 2(ShowScope) + 3(BeforeCallHooks) - 1(AfterCallHooks)
	})
}

func TestCreateMapList(t *testing.T) {
	t.Run("CreateMap NoPlugin", func(tt *testing.T) {
		orders := []map[string]any{RandomOrderMap(), RandomOrderMap(), RandomOrderMap()}
		Err(tt, iSqlite0().Model(&Order{}).Create(&orders)) // raw gorm create map: 4-fields
		// tests.AssertEqual(t, len(orders), 3) // len is 6, hahaha, gorm framework not fix it
		tt.Log(Enc(orders))
	})
	t.Run("CreateMap IsPlugin: Uniques, NoScopes", func(tt *testing.T) {
		orders := []map[string]any{RandomOrderMap(), RandomOrderMap(), RandomOrderMap()}
		Err(tt, iSqlite().Table(`tbl_order`).Create(&orders)) // plugins: before set default values, after set Pk
		tt.Log(Enc(orders))
		tests.AssertEqual(t, len(orders), 3) // Must 3
		slices.Values(orders)(func(order map[string]any) bool {
			tests.AssertEqual(t, len(order), 7) // 4 + 1(Pk) + 2(AutoTime)
			return true
		})
	})
	t.Run("CreateMap IsPlugin: Uniques, 1Scopes(user_id)", func(tt *testing.T) {
		orders := []map[string]any{RandomOrderMap(), RandomOrderMap(), RandomOrderMap()}
		Err(tt, iSUser().Set(gormx.OptionKey, CMapHideScope).
			Table(`tbl_order`).Create(&orders))
		tt.Log(Enc(orders))
		tests.AssertEqual(t, len(orders), 3) // Must 3
		slices.Values(orders)(func(order map[string]any) bool {
			tests.AssertEqual(t, len(order), 7) // 4 + 1(Pk) + 2(AutoTime) + 0(1HideScope)
			return true
		})
	})
	t.Run("CreateMap IsPlugin: Uniques, 2Scopes(user_id, tenant_id)", func(tt *testing.T) {
		orders := []map[string]any{RandomOrderMap(), RandomOrderMap(), RandomOrderMap()}
		Err(tt, iSUserTenant().Set(gormx.OptionKey, CMapShowScope).
			Table(`tbl_order`).Create(&orders))
		tt.Log(Enc(orders))
		tests.AssertEqual(t, len(orders), 3) // Must 3
		slices.Values(orders)(func(order map[string]any) bool {
			tests.AssertEqual(t, len(order), 9) // 4 + 1(Pk) + 2(AutoTime) + 2(ShowScope)
			return true
		})
	})
	t.Run("CreateMap IsPlugin: Uniques, 2Scopes(user_id, tenant_id); Reinforced: before call_hooks", func(tt *testing.T) {
		orders := []map[string]any{RandomOrderMap(), RandomOrderMap(), RandomOrderMap()}
		Err(tt, iSUserTenant().Set(gormx.OptionKey, CMapBeforeCallHook).
			Table(`tbl_order`).Create(&orders)) // before call_hooks
		tt.Log(Enc(orders))
		tests.AssertEqual(t, len(orders), 3) // Must 3
		slices.Values(orders)(func(order map[string]any) bool {
			tests.AssertEqual(t, len(order), 12) // 4 + 1(Pk) + 2(AutoTime) + 2(ShowScope) + 3(BeforeCallHooks)
			return true
		})
	})
	t.Run("CreateMap IsPlugin: Uniques, 2Scopes(user_id, tenant_id); Reinforced: before and after call_hooks", func(tt *testing.T) {
		orders := []map[string]any{RandomOrderMap(), RandomOrderMap(), RandomOrderMap()}
		Err(tt, iSUserTenant().Set(gormx.OptionKey, CMapAfterCallHook).
			Table(`tbl_order`).Create(&orders)) // after call_hooks
		tt.Log(Enc(orders))
		tests.AssertEqual(t, len(orders), 3) // Must 3
		slices.Values(orders)(func(order map[string]any) bool {
			tests.AssertEqual(t, len(order), 11) // 4 + 1(Pk) + 2(AutoTime) + 2(ShowScope) + 3(BeforeCallHooks) - 1(AfterCallHooks)
			return true
		})
	})
}
