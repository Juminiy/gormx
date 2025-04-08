package gormx_testv2

import (
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"testing"
)

func TestOptimizeSQLDelayJoin(t *testing.T) {
	var orders []Order

	var delayJoinCfg = gormx.Option{Optimize: &gormx.OptimizeOption{QueryOffsetDelayJoin: util.New[int64](50000)}}

	queryFn := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("amount_total BETWEEN ? AND ?", 10, 5000000).
			Order("shipping_fee desc").
			Limit(10).Offset(500000).Find(&orders)
	}

	t.Run("NoPlugin", func(tt *testing.T) {
		Err(tt, queryFn(iSqlite()))
	})

	t.Run("IsPlugin", func(tt *testing.T) {
		Err(tt, queryFn(iSqlite().Set(gormx.OptionKey, delayJoinCfg)))
	})
}
