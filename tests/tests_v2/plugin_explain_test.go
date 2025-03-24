package gormx_testv2

import (
	"github.com/Juminiy/gormx"
	"testing"
)

func TestExplainRow(t *testing.T) {
	var breadSale BreadSale
	Err(t, iSqlite().
		Set(gormx.OptionKey, gormx.Option{ExplainQueryOrRow: true}).
		Raw("SELECT * FROM tbl_bread_sale WHERE id = 1 and deleted_at is NULL").
		Scan(&breadSale))
}

func TestExplainQuery(t *testing.T) {
	var breadSale BreadSale
	Err(t, iSqlite().
		Set(gormx.OptionKey, gormx.Option{ExplainQueryOrRow: true}).
		First(&breadSale, 1))
}
