package gormx_testv2

import (
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/gormx/tenants"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"testing"
)

type Product struct {
	gorm.Model
	Name       string `x:"unique:name"`
	Desc       string
	Price      int64
	Total      int64
	Status     int
	PCategory  int         `x:"unique:p"`
	PBrand     string      `x:"unique:p"`
	PModel     string      `x:"unique:p"`
	EAN        string      `x:"unique:ean,name"`
	UPC        string      `gorm:"<-;->:false" x:"unique:upc,name"`
	MerchantID tenants.SID `gorm:"index" x:"tenant"`
}

func RandomProduct() *Product {
	fakeProduct := gofakeit.Product()
	return &Product{
		Name:      fakeProduct.Name,
		Desc:      fakeProduct.Description,
		Price:     int64(math.Floor(fakeProduct.Price * 100)),
		Total:     int64(gofakeit.IntRange(100, 10000)),
		Status:    gofakeit.IntN(2),
		PCategory: gofakeit.IntN(18),
		PBrand:    gofakeit.BookGenre(),
		PModel:    fakeProduct.Material,
		EAN:       gofakeit.BeerYeast(),
		UPC:       fakeProduct.UPC,
	}
}

func iSTenantX() *gorm.DB {
	return iSUserTenant().
		Set(gormx.OptionKey, gormx.Option{
			DisableFieldDup:          false,
			EnableComplexFieldDup:    true,
			AfterCreateShowTenant:    true,
			BeforeCreateMapCallHooks: true,
			AfterCreateMapCallHooks:  true,
			AllowTenantGlobalDelete:  false,
			BeforeDeleteReturning:    true,
			AllowTenantGlobalUpdate:  false,
			UpdateMapOmitZeroElemKey: false,
			UpdateMapOmitUnknownKey:  true,
			UpdateMapSetPkToClause:   true,
			UpdateMapCallHooks:       true,
			AfterUpdateReturning:     true,
			BeforeQueryOmitField:     true,
			AfterQueryShowTenant:     true,
			QueryDynamicSQL:          true,
			WriteClauseToRowOrRaw:    true,
		})
}

func TestTenantsString(t *testing.T) {
	t.Run("crud product", func(tt *testing.T) {
		var cProduct = RandomProduct()
		Err(tt, iSTenantX().Create(&cProduct))
		tt.Log(Enc(cProduct))

		var rProduct = &Product{Model: gorm.Model{ID: cProduct.ID}}
		Err(tt, iSTenantX().First(&rProduct))
		tt.Log(Enc(rProduct))

		var uProduct = map[string]any{
			"id":          cProduct.ID, // Pk to clause
			"p_category":  gofakeit.IntN(18),
			"p_brand":     gofakeit.BookGenre(),
			"p_model":     "", // update to ""
			"ean":         gofakeit.BeerYeast(),
			"some_column": "some_value", // omit unknown
		}
		Err(tt, iSTenantX().Table(`tbl_product`).Updates(uProduct))
		tt.Log(Enc(uProduct))

		var dProduct = &Product{Model: gorm.Model{ID: cProduct.ID}}
		Err(tt, iSTenantX().Clauses(clause.Returning{}).Delete(&dProduct))
		tt.Log(Enc(dProduct))
	})
}
