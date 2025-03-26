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

func TestExplainQueryPg(t *testing.T) {
	/*var breadProduct BreadProduct
	var breadSale BreadSale
	Err(t, iPG().Create(&breadProduct))
	Err(t, iPG().Create(RandomBreadSale(&breadProduct)))
	Err(t, iPG().
		Set(gormx.OptionKey, gormx.Option{ExplainQueryOrRow: true}).
		First(&breadSale, 1))*/
}

func TestExplainQueryInnoDB(t *testing.T) {
	/*var breadProduct BreadProduct
	var breadSale BreadSale
	Err(t, iInnoDB().Create(&breadProduct))
	Err(t, iInnoDB().Create(RandomBreadSale(&breadProduct)))
	Err(t, iInnoDB().
		Set(gormx.OptionKey, gormx.Option{ExplainQueryOrRow: true}).
		First(&breadSale, 1))*/
}
