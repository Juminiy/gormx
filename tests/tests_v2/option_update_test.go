package gormx_testv2

import (
	"database/sql"
	"github.com/Juminiy/gormx"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
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

var UOptLock = gormx.Option{UpdateOptimisticLock: true, UpdateMapSetPkToClause: true}

func updateSkipTxn() *gorm.DB {
	return iSqlite().Session(&gorm.Session{NewDB: true, SkipDefaultTransaction: true})
}

func TestUpdateMapOptLock(t *testing.T) {
	t.Run("IsPlugin, No Hide Pk Clause", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, updateSkipTxn().Create(&order))
		Err(tt, updateSkipTxn().Set(gormx.OptionKey, UOptLock).
			Where("id = ?", order.ID).Updates(&Order{PayTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}}))
		Err(tt, updateSkipTxn().Set(gormx.OptionKey, UOptLock).
			Table(`tbl_order`).
			Where("id = ?", order.ID).Updates(map[string]any{
			"pay_time": time.Now(),
		}))
	})
	t.Run("IsPlugin, No Field Can Version", func(tt *testing.T) {

	})
	t.Run("IsPlugin, Updates(Struct) DestPk", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, updateSkipTxn().Create(&order))
		Err(tt, updateSkipTxn().Set(gormx.OptionKey, UOptLock).
			Updates(&Order{
				Model: gorm.Model{ID: order.ID},
				PayTime: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				}}))
	})
	t.Run("IsPlugin, Updates(Map) DestPk", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, updateSkipTxn().Create(&order))
		Err(tt, updateSkipTxn().Set(gormx.OptionKey, UOptLock).
			Table(`tbl_order`).
			Updates(map[string]any{
				"id":       order.ID,
				"pay_time": time.Now(),
			}))
	})
	t.Run("IsPlugin, Model(Struct).Updates(Struct) ModelPk", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, updateSkipTxn().Create(&order))
		Err(tt, updateSkipTxn().Set(gormx.OptionKey, UOptLock).
			Model(&order).
			Updates(&Order{
				PayTime: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				}}))
	})
	t.Run("IsPlugin, Model(Struct).Updates(Map) ModelPk", func(tt *testing.T) {
		order := RandomOrder()
		Err(tt, updateSkipTxn().Create(&order))
		Err(tt, updateSkipTxn().Set(gormx.OptionKey, UOptLock).
			Model(&order).
			Updates(map[string]any{
				"pay_time": time.Now(),
			}))
	})
}

type ChuckBlock struct {
	soft_delete.DeletedAt
	ID        uint  `gorm:"primaryKey;autoIncrement"`
	CreatedAt int64 `gorm:"autoCreateTime:milli"`
	Version   int64 `x:"version"`
	IdleSize  int64
	MaxSize   int64
	MinSize   int64
	WiseDesc  string
}

func RandomChuckBlock() *ChuckBlock {
	return &ChuckBlock{
		IdleSize: gofakeit.Int64(),
		MaxSize:  gofakeit.Int64(),
		MinSize:  gofakeit.Int64(),
		WiseDesc: gofakeit.EmojiDescription(),
	}
}

func TestUpdateMapOptLockByTag(t *testing.T) {
	t.Run("Plugin Version Int", func(tt *testing.T) {
		chuckBlk := RandomChuckBlock()
		Err(tt, updateSkipTxn().Create(&chuckBlk))
		Err(tt, updateSkipTxn().
			Set(gormx.OptionKey, UOptLock).
			Table(`tbl_chuck_block`).
			Updates(map[string]any{
				"id":       chuckBlk.ID,
				"min_size": gofakeit.Int64(),
			}))
		Err(tt, updateSkipTxn().
			Set(gormx.OptionKey, UOptLock).
			Updates(&ChuckBlock{
				ID:      chuckBlk.ID,
				MaxSize: gofakeit.Int64(),
			}))
	})
}

func TestUpdateMapOptLockByTagFull(t *testing.T) {
	t.Run("IsPlugin, Updates(Struct) DestPk", func(tt *testing.T) {
		chuckBlk := RandomChuckBlock()
		Err(tt, updateSkipTxn().Create(&chuckBlk))
		Err(tt, updateSkipTxn().Set(gormx.OptionKey, UOptLock).
			Updates(&ChuckBlock{
				ID:      chuckBlk.ID,
				MinSize: gofakeit.Int64(),
			}))
	})
	t.Run("IsPlugin, Updates(Map) DestPk", func(tt *testing.T) {
		chuckBlk := RandomChuckBlock()
		Err(tt, updateSkipTxn().Create(&chuckBlk))
		Err(tt, updateSkipTxn().Set(gormx.OptionKey, UOptLock).
			Table(`tbl_chuck_block`).
			Updates(map[string]any{
				"id":       chuckBlk.ID,
				"max_size": gofakeit.Int64(),
			}))
	})
	t.Run("IsPlugin, Model(Struct).Updates(Struct) ModelPk", func(tt *testing.T) {
		chuckBlk := RandomChuckBlock()
		Err(tt, updateSkipTxn().Create(&chuckBlk))
		Err(tt, updateSkipTxn().Set(gormx.OptionKey, UOptLock).
			Model(&chuckBlk).
			Updates(&ChuckBlock{
				IdleSize: gofakeit.Int64(),
			}))
	})
	t.Run("IsPlugin, Model(Struct).Updates(Map) ModelPk", func(tt *testing.T) {
		chuckBlk := RandomChuckBlock()
		Err(tt, updateSkipTxn().Create(&chuckBlk))
		Err(tt, updateSkipTxn().Set(gormx.OptionKey, UOptLock).
			Model(&chuckBlk).
			Updates(map[string]any{
				"wise_desc": gofakeit.EmojiDescription(),
			}))
	})
}
