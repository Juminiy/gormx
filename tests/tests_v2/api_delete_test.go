package gormx_testv2

import "testing"

func TestDeleteStruct(t *testing.T) {
	t.Run("Delete Struct NoPlugin", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSqlite0().Create(&order))
		Err(tt, iSqlite0().Delete(&Order{}, order.ID))
	})
	t.Run("Delete Struct IsPlugin: NoScopes", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSqlite().Create(&order))
		Err(tt, iSqlite().Delete(&Order{}, order.ID))
	})
	t.Run("Delete Struct IsPlugin: 1Scopes(user_id)", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSUser().Create(&order))
		Err(tt, iSUser().Delete(&Order{}, order.ID))
	})
	t.Run("Delete Struct IsPlugin: 2Scopes(user_id, tenant_id)", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSUserTenant().Create(&order))
		Err(tt, iSUserTenant().Delete(&Order{}, order.ID))
	})
}

func TestDeleteStructList(t *testing.T) {
	t.Run("Delete StructList NoPlugin", func(tt *testing.T) {
		orders := []*Order{RandomOrder(), RandomOrder(), RandomOrder()}
		Err(tt, iSqlite0().Create(&orders))
		Err(tt, iSqlite0().Delete(orders))
	})
	t.Run("Delete StructList IsPlugin: NoScopes", func(tt *testing.T) {
		orders := []*Order{RandomOrder(), RandomOrder(), RandomOrder()}
		Err(tt, iSqlite().Create(&orders))
		Err(tt, iSqlite().Delete(&orders))
	})
	t.Run("Delete StructList IsPlugin: 1Scopes(user_id)", func(tt *testing.T) {
		orders := []*Order{RandomOrder(), RandomOrder(), RandomOrder()}
		Err(tt, iSUser().Create(&orders))
		Err(tt, iSUser().Delete(&orders))
	})
	t.Run("Delete StructList IsPlugin: 2Scopes(user_id, tenant_id)", func(tt *testing.T) {
		orders := []*Order{RandomOrder(), RandomOrder(), RandomOrder()}
		Err(tt, iSUserTenant().Create(&orders))
		Err(tt, iSUserTenant().Delete(&orders))
	})
}
