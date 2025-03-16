package gormx_testv2

import (
	"github.com/Juminiy/gormx"
	"github.com/google/uuid"
	"maps"
	"testing"
	"time"
)

func TestUpdateOptionOmitUnknownKey(t *testing.T) {
	omitUnknown := gormx.Option{UpdateMapOmitUnknownKey: true}
	t.Run("NoPlugin", func(tt *testing.T) {
		if err := iSqlite0().
			Table(`tbl_order`).Where("id = 1").
			Update("no_column", 1).Error; err != nil {
			tt.Log(err.Error())
		} else {
			tt.Log("gorm support update map omit unknown")
		}
	})
	t.Run("IsPlugin, NoOption", func(tt *testing.T) {
		if err := iSqlite().
			Table(`tbl_order`).Where("id = 1").
			Update("no_column", 1).Error; err != nil {
			tt.Log(err.Error())
		} else {
			tt.Error("plugin update map omit unknown option misOpen")
		}
	})
	t.Run("IsPlugin, IsOption", func(tt *testing.T) {
		Err(tt, iSqlite().
			Set(gormx.OptionKey, omitUnknown).
			Table(`tbl_order`).Where("id = 1").
			Update("no_column", 1))
	})
}

func TestUpdateOptionOmitZeroElem(t *testing.T) {
	omitZero := gormx.Option{
		UpdateMapOmitUnknownKey: true, UpdateMapOmitZeroElemKey: true}
	updateMap := map[string]any{
		"logistics":      "", // no such column
		"shipping_fee":   0,  // zero int
		"logistics_name": "", // zero string
	}
	t.Run("NoPlugin", func(tt *testing.T) {
		if err := iSqlite0().
			Table(`tbl_order`).Where("id = 1").
			Updates(maps.Clone(updateMap)).Error; err != nil {
			tt.Log(err.Error())
		} else {
			tt.Log("gorm support update map omit zeroElem")
		}
	})
	t.Run("IsPlugin, NoOption", func(tt *testing.T) {
		if err := iSqlite().
			Table(`tbl_order`).Where("id = 1").
			Updates(maps.Clone(updateMap)).Error; err != nil {
			tt.Log(err.Error())
		} else {
			tt.Error("plugin update map omit zeroElem option misOpen")
		}
	})
	t.Run("IsPlugin, IsOption", func(tt *testing.T) {
		Err(tt, iSqlite().
			Set(gormx.OptionKey, omitZero).
			Table(`tbl_order`).Where("id = 1").
			Updates(maps.Clone(updateMap)))
	})
}

func TestUpdateOptionSetPkToClause(t *testing.T) {
	setPk := gormx.Option{
		UpdateMapOmitUnknownKey: true, UpdateMapOmitZeroElemKey: true,
		UpdateMapSetPkToClause: true}
	updateMap := map[string]any{
		"logistics":      "",               // no such column
		"shipping_fee":   0,                // zero int
		"logistics_name": "",               // zero string
		"id":             1,                // primaryKey id
		"serial":         uuid.NewString(), // primaryKey serial
	}
	t.Run("NoPlugin", func(tt *testing.T) {
		if err := iSqlite0().
			Table(`tbl_order`).
			Updates(maps.Clone(updateMap)).Error; err != nil {
			tt.Log(err.Error())
		} else {
			tt.Log("gorm support updateMap id as clause")
		}
	})
	t.Run("IsPlugin, NoOption", func(tt *testing.T) {
		if err := iSqlite().
			Table(`tbl_order`).
			Updates(maps.Clone(updateMap)).Error; err != nil {
			tt.Log(err.Error())
		} else {
			tt.Error("plugin update updateMap id as clause option misOpen")
		}
	})
	t.Run("IsPlugin, IsOption", func(tt *testing.T) {
		Err(tt, iSqlite().
			Set(gormx.OptionKey, setPk).
			Table(`tbl_order`).
			Updates(maps.Clone(updateMap)))
	})
}

func TestUpdateMapBeforeCallHooks(t *testing.T) {
	mapCallHooks := gormx.Option{UpdateMapCallHooks: true}
	t.Run("NoPlugin", func(tt *testing.T) {
		Err(tt, iSqlite0().
			Set("serial", uuid.NewString()).
			Table(`tbl_order`).Where("id = 1").
			Update("shipped_time", time.Now()))
	})
	t.Run("IsPlugin, NoOption", func(tt *testing.T) {
		Err(tt, iSqlite().
			Set("serial", uuid.NewString()).
			Table(`tbl_order`).Where("id = 1").
			Update("shipped_time", time.Now()))
	})
	t.Run("IsPlugin, IsOption", func(tt *testing.T) {
		Err(tt, iSqlite().
			Set("serial", uuid.NewString()).
			Set(gormx.OptionKey, mapCallHooks).
			Table(`tbl_order`).Where("id = 1").
			Update("shipped_time", time.Now()))
	})
}

func TestUpdateMapReturning(t *testing.T) {
	/*mapReturning := gormx.Option{AfterUpdateReturning: true}
	updateMap := map[string]any{
		"pay_method":      MethodWechatPay,
		"pay_time":        time.Now(),
		"order_status":    StatusPaid,
		"amount_discount": gorm.Expr("amount_discount + ?", gofakeit.IntRange(1, 20)),
	}
	t.Run("NoPlugin", func(tt *testing.T) {
		var updateDest = maps.Clone(updateMap)
		Err(tt, iInnoDB0().Model(&Order{Model: gorm.Model{ID: 1}}).
			Clauses(clause.Returning{}).Updates(updateDest))
		t.Log(Enc(updateDest))
	})
	t.Run("IsPlugin, NoOption", func(tt *testing.T) {
		var updateDest = maps.Clone(updateMap)
		Err(tt, iInnoDB().Model(&Order{Model: gorm.Model{ID: 1}}).
			Clauses(clause.Returning{}).Updates(updateDest))
		t.Log(Enc(updateDest))
	})
	t.Run("IsPlugin, Option", func(tt *testing.T) {
		var updateDest = maps.Clone(updateMap)
		Err(tt, iInnoDB().Set(gormx.OptionKey, mapReturning).
			Model(&Order{Model: gorm.Model{ID: 1}}).
			Clauses(clause.Returning{}).Updates(updateDest))
		t.Log(Enc(updateDest))
	})*/
}
