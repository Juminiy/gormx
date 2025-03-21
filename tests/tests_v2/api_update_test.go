package gormx_testv2

import (
	"github.com/Juminiy/gormx"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"testing"
)

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

func TestUpdateByImplicitPk(t *testing.T) {
	testUpdateByImplicitPk(t, "NoPlugin", iSqlite0)
	testUpdateByImplicitPk(t, "IsPlugin", func() *gorm.DB {
		return iSqlite().Set(gormx.OptionKey, gormx.Option{UpdateMapSetPkToClause: true})
	})
}

func testUpdateByImplicitPk(t *testing.T, called string, ntxFn func() *gorm.DB) {
	t.Run(called+" DestPk Updates(Struct)", func(tt *testing.T) {
		chkBlk := RandomChuckBlock()
		Err(tt, ntxFn().Create(&chkBlk))
		Err(tt, ntxFn().Updates(&ChuckBlock{
			ID:      chkBlk.ID,
			MaxSize: gofakeit.Int64(),
		}))
	})
	t.Run(called+" ModelPk ModelPtr Updates(Struct)", func(tt *testing.T) {
		chkBlk := RandomChuckBlock()
		Err(tt, ntxFn().Create(&chkBlk))
		Err(tt, ntxFn().Model(&chkBlk).Updates(&ChuckBlock{
			MaxSize: gofakeit.Int64(),
		}))
	})
	t.Run(called+" ModelPk ModelValue Updates(Struct)", func(tt *testing.T) {
		chkBlk := RandomChuckBlock()
		Err(tt, ntxFn().Create(&chkBlk))
		Err(tt, ntxFn().Model(chkBlk).Updates(&ChuckBlock{
			MaxSize: gofakeit.Int64(),
		}))
	})
	t.Run(called+" ModelPkAndDestPk ModelPtr Updates(Struct)", func(tt *testing.T) {
		chkBlk := RandomChuckBlock()
		Err(tt, ntxFn().Create(&chkBlk))
		Err(tt, ntxFn().Model(&chkBlk).Updates(&ChuckBlock{
			ID:      chkBlk.ID,
			MaxSize: gofakeit.Int64(),
		}))
	})
	t.Run(called+" ModelPkAndDestPk ModelValue Updates(Struct)", func(tt *testing.T) {
		chkBlk := RandomChuckBlock()
		Err(tt, ntxFn().Create(&chkBlk))
		Err(tt, ntxFn().Model(chkBlk).Updates(&ChuckBlock{
			ID:      chkBlk.ID,
			MaxSize: gofakeit.Int64(),
		}))
	})
	t.Run(called+" ModelPk ModelPtr Updates(Map)", func(tt *testing.T) {
		chkBlk := RandomChuckBlock()
		Err(tt, ntxFn().Create(&chkBlk))
		Err(tt, ntxFn().Model(&chkBlk).Updates(map[string]any{
			"max_size": gofakeit.Int64(),
		}))
	})
	t.Run(called+" ModelPk ModelValue Updates(Map)", func(tt *testing.T) {
		chkBlk := RandomChuckBlock()
		Err(tt, ntxFn().Create(&chkBlk))
		Err(tt, ntxFn().Model(chkBlk).Updates(map[string]any{
			"max_size": gofakeit.Int64(),
		}))
	})
	t.Run(called+" DestPk Updates(Map)", func(tt *testing.T) {
		chkBlk := RandomChuckBlock()
		Err(tt, ntxFn().Create(&chkBlk))
		if err := ntxFn().Table(`tbl_chuck_block`).Updates(map[string]any{
			"id":       chkBlk.ID,
			"max_size": gofakeit.Int64(),
		}).Error; err != nil {
			t.Logf("%s not support UpdateMap Pk in Map", called)
		}
	})
	t.Run(called+" ModelPkAndDestPk ModelPtr Updates(Map)", func(tt *testing.T) {
		chkBlk := RandomChuckBlock()
		Err(tt, ntxFn().Create(&chkBlk))
		Err(tt, ntxFn().Model(&chkBlk).Updates(map[string]any{
			"id":       chkBlk.ID,
			"max_size": gofakeit.Int64(),
		}))
	})
	t.Run(called+" ModelPkAndDestPk ModelValue Updates(Map)", func(tt *testing.T) {
		chkBlk := RandomChuckBlock()
		Err(tt, ntxFn().Create(&chkBlk))
		Err(tt, ntxFn().Model(chkBlk).Updates(map[string]any{
			"id":       chkBlk.ID,
			"max_size": gofakeit.Int64(),
		}))
	})
}
