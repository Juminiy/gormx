package gormx_tests

import (
	"github.com/Juminiy/gormx"
	"gorm.io/gorm"
	"testing"
)

func _txListDup() *gorm.DB {
	return _txTenant().Set(gormx.OptionKey, gormx.Option{
		EnableComplexFieldDup: true,
	})
}

func TestCreateStructListDup(t *testing.T) {
	Err(t,
		_txListDup().Create(&[]Product{
			{Name: "Milk", Desc: "Fresh milk", NetContent: "1L", Code: 100001, Price: 800},
			{Name: "Bread", Desc: "Whole wheat bread", NetContent: "500g", Code: 100002, Price: 1200},
			{Name: "Rice", Desc: "Long grain rice", NetContent: "5kg", Code: 100003, Price: 4500},
			{Name: "Eggs", Desc: "Free-range eggs", NetContent: "12 pcs", Code: 100004, Price: 1500},
			{Name: "Chicken", Desc: "Fresh chicken breast", NetContent: "1kg", Code: 100006, Price: 3500},
		}).Error)
}

func TestCreateMapListDup(t *testing.T) {
	Err(t,
		_txListDup().Model(&Product{}).Create(&[]map[string]any{
			{"Name": "Beer", "Desc": "Local lager beer", "NetContent": "500ml", "Code": 100007, "Price": 500},
			{"Name": "Noodles", "Desc": "Instant noodles", "NetContent": "5 packs", "Code": 100008, "Price": 1000},
			{"Name": "Shampoo", "Desc": "Herbal shampoo", "NetContent": "400ml", "Code": 100009, "Price": 2500},
			{"Name": "Toothpaste", "Desc": "Mint toothpaste", "NetContent": "120g", "Code": 100010, "Price": 800},
		}).Error)
}

func TestCreateStructDupInExpr(t *testing.T) {
	Err(t, txScope().Create((&BabyTrade{}).RandomSet()).Error)
}
