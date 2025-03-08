package gormx_tests

import (
	"github.com/Juminiy/gormx"
	"testing"
)

func TestFindMapHooks(t *testing.T) {
	var consumerMap map[string]any
	Err(t, txHooks().Table(`tbl_consumer`).First(&consumerMap, 156).Error)
	t.Log(Enc(consumerMap))

	var consumerMapList []map[string]any
	Err(t, txHooks().Table(`tbl_consumer`).Find(&consumerMapList, 157, 158, 159).Error)
	t.Log(Enc(consumerMapList))
}

func TestBeforeQueryOmitField(t *testing.T) {
	var list []AppUser
	Err(t, _txTenant().Set(gormx.OptionKey, gormx.Option{
		BeforeQueryOmitField: true,
	}).First(&list).Error)
	t.Log(Enc(list))
}
