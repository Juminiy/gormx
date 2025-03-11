package gormx_tests

import (
	"github.com/Juminiy/gormx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
)

func TestSqliteReturning(t *testing.T) {
	var bt BabyTrade
	Err(t, txPure().Clauses(clause.Returning{}).Delete(&bt, 39).Error)
	t.Log(Enc(bt))
}

func txMysql() *gorm.DB {
	return iMySQL().
		Set("tenant_id", 114514).
		Set("user_id", 114514).
		Set(gormx.OptionKey, gormx.Option{
			UpdateMapSetPkToClause:  true,
			UpdateMapCallHooks:      true,
			UpdateMapOmitUnknownKey: true,
			AfterUpdateReturning:    true,
			BeforeDeleteReturning:   true,
		})
}

func TestMySQLDeleteReturning(t *testing.T) {
	var bt BabyTrade

	Err(t, txMysql().Clauses(clause.Returning{}).
		Delete(&bt, 32021).Error)
	t.Log(Enc(bt))
}

func TestMySQLUpdateReturning(t *testing.T) {
	mapValue := map[string]any{
		"id":         32024, // Pk
		"auction_id": 10,    // set
		"cat_id":     11,    // set
		"missing_id": 24,    // unknown
		"unknown_id": 46,    // unknown
		//"sim_uuid":    "bc06526c-57b1-4579-831a-8354b5672d87", // Pk, (zero, omit and do nothing)/(not-zero, as clause)
		"buy_mount":   0,  // zero
		"zero_id":     0,  // unknown,zero
		"zero_id_str": "", // unknown,zero
		"tenant_id":   10,
		"user_id":     10,
	}
	Err(t, txMysql().Table(`tbl_baby_trade`).
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "auction_id"}, {Name: "buy_mount"}, {Name: "cat_id"}}}).
		Updates(mapValue).Error)
	t.Log(Enc(mapValue))
}
