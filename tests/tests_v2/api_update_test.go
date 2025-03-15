package gormx_testv2

import "testing"

func TestUpdateStruct(t *testing.T) {
	t.Run("Update Struct NoPlugin", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSqlite0().Create(&order))
		Err(tt, iSqlite0().Updates(order.SetPayInfo()))
		Err(tt, iSqlite0().Updates(order.SetShipInfo()))
	})
	t.Run("Update Struct IsPlugin: NoScopes", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSqlite().Create(&order))
		Err(tt, iSqlite().Updates(order.SetPayInfo()))
		Err(tt, iSqlite().Updates(order.SetShipInfo()))
	})
	t.Run("Update Struct IsPlugin: Scopes(user_id)", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSUser().Create(&order))
		Err(tt, iSUser().Updates(order.SetPayInfo()))
		Err(tt, iSUser().Updates(order.SetShipInfo()))
	})
	t.Run("Update Struct IsPlugin: Scopes(user_id, tenant_id)", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSUserTenant().Create(&order))
		Err(tt, iSUserTenant().Updates(order.SetPayInfo()))
		Err(tt, iSUserTenant().Updates(order.SetShipInfo()))
	})
	t.Run("Update Struct IsPlugin: Scopes(user_id, tenant_id), with another clause", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSUserTenant().Create(&order))
		Err(tt, iSUserTenant().Where("id = ?", order.ID).Updates(order.SetPayInfo()))
		Err(tt, iSUserTenant().Where("id = ?", order.ID).Updates(order.SetShipInfo()))
	})
}

func TestUpdateMap(t *testing.T) {
	t.Run("Update Map NoPlugin", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSqlite0().Create(&order))
		Err(tt, iSqlite0().Model(order).Updates(RandomShipMap(order.ShippingFee)))
	})
	t.Run("Update Map IsPlugin: NoScopes", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSqlite().Create(&order))
		Err(tt, iSqlite().Model(order).Updates(RandomShipMap(order.ShippingFee)))
	})
	t.Run("Update Map IsPlugin: Scopes(user_id)", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSUser().Create(&order))
		Err(tt, iSUser().Model(order).Updates(RandomShipMap(order.ShippingFee)))
	})
	t.Run("Update Map IsPlugin: Scopes(user_id, tenant_id)", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSUserTenant().Create(&order))
		Err(tt, iSUserTenant().Model(order).Updates(RandomShipMap(order.ShippingFee)))
	})
	t.Run("Update Map IsPlugin: Scopes(user_id, tenant_id), with another clause", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, iSUserTenant().Create(&order))
		Err(tt, iSUserTenant().Model(order).Where("id = ?", order.ID).Updates(RandomShipMap(order.ShippingFee)))
	})
}
