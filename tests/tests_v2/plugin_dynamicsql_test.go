package gormx_testv2

import (
	"errors"
	"github.com/Juminiy/gormx"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"testing"
)

var ErrDynamicSQL = errors.New("dynamic sql func is error")

func TestRowRaw(t *testing.T) {
	addClause := func(tx *gorm.DB) *gorm.DB {
		return tx.
			Raw(`SELECT * FROM tbl_order`).
			Where("id BETWEEN ? AND ?", gofakeit.IntN(16), gofakeit.IntN(256)).
			Where("amount_total IN ?"). // not enough args
			Where("pay_time IS NULL").
			Where("logistics_id <> ?", gofakeit.IntN(20), gofakeit.IntN(40)). // overflow args
			Order("id asc").
			Order(""). // null order
			Limit(10).
			Offset(3)
	}

	t.Run("NoPlugin", func(tt *testing.T) {
		var orders []Order
		Err(tt, addClause(iSqlite0()).Scan(&orders))
	})
	t.Run("IsPlugin NoOption", func(tt *testing.T) {
		var orders []Order
		Err(tt, addClause(iSqlite()).Scan(&orders))
	})
	t.Run("IsPlugin Option", func(tt *testing.T) {
		var orders []Order
		Err(tt, addClause(iSqlite().
			Set(gormx.OptionKey, gormx.Option{WriteClauseToRowOrRaw: true})).
			Scan(&orders))
	})
}

func TestDynamicSQL(t *testing.T) {
	addClause := func(tx *gorm.DB) *gorm.DB {
		return tx.
			Where("id >= ? AND id <= ?"). // lack 2args
			Where("id IN ?").             // lack 1args
			Where("pay_method IS NOT NULL").
			Where("id = ?", gofakeit.IntN(18), gofakeit.IntN(22)).                     // more 1args
			Where("serial = ?", uuid.NewString(), uuid.NewString(), uuid.NewString()). // more 2args
			Where("shipped_time IS NULL").
			Where("order_type = ?", gofakeit.IntRange(1, 4)).
			Order("created_at DESC").
			Order("amount_total ASC").
			Order("").
			Limit(10).
			Offset(3)
	}
	t.Run("NoPlugin", func(tt *testing.T) {
		var orders []Order
		if err := addClause(iSqlite0()).Find(&orders).Error; err == nil {
			tt.Log("gorm support omit illegal clause")
		} else {
			tt.Log("gorm not support omit illegal clause")
		}
	})
	t.Run("IsPlugin, NoOption", func(tt *testing.T) {
		var orders []Order
		if err := addClause(iSqlite()).Find(&orders).Error; err == nil {
			tt.Log(ErrDynamicSQL)
		}
	})
	t.Run("IsPlugin, Option", func(tt *testing.T) {
		var orders []Order
		Err(tt, addClause(iSqlite().
			Set(gormx.OptionKey, gormx.Option{QueryDynamicSQL: true})).
			Find(&orders))
	})
}
