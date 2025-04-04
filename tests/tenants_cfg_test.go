package gormx_tests

import (
	"github.com/Juminiy/gormx"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
)

func _txTenant() *gorm.DB {
	return txPure().Set("tenant_id", uint(114514))
}

func TestQueryBeforeDeleteOne(t *testing.T) {
	prod := Product{}
	err := _txTenant().Set(gormx.OptionKey, gormx.Option{
		BeforeDeleteReturning: true,
	}).Clauses(clause.Returning{Columns: []clause.Column{
		{Name: "name"}, {Name: "desc"}, {Name: "code"},
	}}).Delete(&prod, 40).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(Enc(prod))
}

func TestQueryBeforeDeleteList(t *testing.T) {
	prod := []Product{}
	err := _txTenant().Set(gormx.OptionKey, gormx.Option{
		BeforeDeleteReturning: true,
	}).Delete(&prod, []int{2, 3, 4}).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(Enc(prod))
}

func TestDeleteOne(t *testing.T) {
	prod := Product{}
	err := _txTenant().Delete(&prod, 5).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(Enc(prod))
}

func TestDeleteList(t *testing.T) {
	prod := Product{}
	err := _txTenant().Delete(&prod, []int{6, 7, 8}).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(Enc(prod))
}

func mtGrp() *gorm.DB {
	return txPure().Set("tenant_id_list", []int{1, 2, 3, 4, 5, 114514})
}

func TestQueryListInTenants(t *testing.T) {
	var list []Product
	Err(t, mtGrp().Find(&list).Error)
	t.Log(Enc(list))
}

func mtSglAndGrp() *gorm.DB {
	return txPure().Set("tenant_id", "114514").
		Set("tenant_id_list", []int{1, 2, 3, 4, 5, 114514})
}

func TestQueryListTenantChooseBest(t *testing.T) {
	var list []Product
	Err(t, mtSglAndGrp().Find(&list).Error)
	t.Log(Enc(list))
}

func TestSkipTenant(t *testing.T) {
	err := _txTenant().Set("skip_tenant", struct{}{}).
		Find(&Product{}).Error
	Err(t, err)
}

func TestSkipFieldDup(t *testing.T) {
	err := _txTenant().Set(gormx.OptionKey, gormx.Option{
		DisableFieldDup: true,
	}).Create(&Product{
		Name:       "Coca-Cola",
		Desc:       "Most Popular Drink in the World",
		NetContent: "500ml",
		Code:       8,
		Price:      299,
	}).Error
	Err(t, err)
}

func TestDeleteTenantAll(t *testing.T) {
	Err(t, _txTenant().Set(gormx.OptionKey, gormx.Option{
		AllowTenantGlobalDelete: true,
	}).Delete(&Product{}).Error)
}

func TestUpdateTenantAll(t *testing.T) {
	Err(t, _txTenant().Set(gormx.OptionKey, gormx.Option{
		AllowTenantGlobalUpdate: true,
	}).Model(&Product{}).Update("code", "114514").Error)
}

func showT() *gorm.DB {
	return _txTenant().Set(gormx.OptionKey, gormx.Option{
		DisableFieldDup:       true,
		AfterCreateShowTenant: true,
		AfterQueryShowTenant:  true,
	})
}

func TestCreateShowTenant(t *testing.T) {
	prod := Product{
		Name:       "Orange",
		Desc:       "orange is a kind of fruit",
		NetContent: "1kg",
		Code:       3301119,
		Price:      80,
		TenantID:   25, // no effect
	}
	Err(t, showT().Create(&prod).Error)
	t.Log(prod.TenantID)
}

func TestQueryShowTenant(t *testing.T) {
	prods := []Product{}
	Err(t, showT().Find(&prods, []uint{1, 2, 3}).Error)
	t.Log(lo.Map(prods, func(item Product, _ int) uint {
		return item.TenantID
	}))
}
