package gormx_testv2

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/Juminiy/gormx/schemas/types"
	"github.com/Juminiy/gormx/tenants"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"testing"
	"time"
)

type BreadSale struct {
	gorm.Model
	BreadID        uint       `gorm:"<-:create"`
	MerchantID     tenants.ID `gorm:"<-:create" x:"merchant"`
	FreshPrice     types.RMBCent
	DayOffPrice    types.RMBCent
	DiscountPrice  types.RMBCent
	DiscountTime   types.Timestamp
	DiscountType   int
	ReleaseCount   int64
	ExpirationTime types.Timestamp `gorm:"<-:create"`
	FirstSaleTime  types.DateTime  `gorm:"<-:update"`
	LastSaleTime   types.DateTime  `gorm:"<-:update"`
	SaleCount      int64
}

func RandomBreadSale(breadInfo *BreadProduct) *BreadSale {
	return &BreadSale{
		BreadID:        breadInfo.ID,
		FreshPrice:     types.RMBCent(gofakeit.Price(815, 3500)),
		DayOffPrice:    types.RMBCent(gofakeit.Price(405, 3500)),
		DiscountPrice:  types.RMBCent(gofakeit.Price(50, 500)),
		ReleaseCount:   int64(gofakeit.IntRange(100, 1200)),
		ExpirationTime: types.Timestamp{Int64: time.Now().Add(breadInfo.ShelfLife).Unix(), Valid: true},
	}
}

type BreadProduct struct {
	gorm.Model
	Name          string
	Type          int             `x:"enum;1:Yeast,2:Unleavened,3:Quick,4:Sourdough,5:Steamed,6:Fried,7:Flatbread,8:Sweet"`
	Category      int             `x:"enum;1:Fresh,2.Packaged,3:Frozen,4:Dessert,5:Healthy,6:Specialty,7:Light,8:Baking"`
	Ingredient    BreadIngredient `gorm:"type:text"`
	WaterContent  int             `x:"range:0~100"`
	CookingMethod int             `x:"enum;1:Baked,2:Steamed,3:Fried,4:Flatbread,5:Boiled,6:Microwaved,7:Grilled,8:Stone-Baked"`
	ShelfLife     time.Duration
	Weight        int
	Size          int
	Status        int `x:"enum;1:online,2:offline"`
}

func RandomBread() *BreadProduct {
	return &BreadProduct{
		Name:          gofakeit.Breakfast(),
		Type:          gofakeit.IntRange(1, 8),
		Category:      gofakeit.IntRange(1, 8),
		Ingredient:    RandomIngredient(),
		WaterContent:  gofakeit.IntRange(30, 80),
		CookingMethod: gofakeit.IntRange(1, 8),
		ShelfLife:     util.DurationDay * time.Duration(gofakeit.IntRange(5, 13)),
		Weight:        gofakeit.IntN(2280),
		Size:          gofakeit.IntN(3330),
		Status:        1,
	}
}

type BreadIngredient map[string]int

func RandomIngredient() BreadIngredient {
	ingreds := make(BreadIngredient, 8)
	for range gofakeit.IntRange(3, 8) {
		ingreds[gofakeit.MinecraftFood()] = gofakeit.IntRange(1, 120)
	}
	return ingreds
}

func (b *BreadIngredient) Scan(src any) error {
	if srcBytes, ok := src.([]byte); ok {
		return json.Unmarshal(srcBytes, b)
	}
	return types.ValueFromDBError("BreadIngredient")
}

func (b BreadIngredient) Value() (driver.Value, error) {
	return json.Marshal(b)
}

type BreadMerchant struct {
	gorm.Model
	Name         string
	Address      string
	ConcatMobile string
	ConcatPerson string
	ConcatPhone  string
}

func RandomMerchant() *BreadMerchant {
	return &BreadMerchant{
		Name:         gofakeit.Name(),
		Address:      gofakeit.Address().Address,
		ConcatMobile: gofakeit.Phone(),
		ConcatPerson: gofakeit.Person().LastName,
		ConcatPhone:  gofakeit.Person().Contact.Phone,
	}
}

func TestTypes(t *testing.T) {
	var (
		iSMrcht   func() *gorm.DB
		mrcht     *BreadMerchant
		breadList []*BreadProduct
		saleList  []*BreadSale
	)
	t.Run("Create BreadList", func(tt *testing.T) {
		breadList = []*BreadProduct{RandomBread(), RandomBread(), RandomBread()}
		Err(tt, iSqlite().Create(&breadList))
		tt.Log(Enc(map[string]any{"bread_id_list": lo.Map(breadList, func(brd *BreadProduct, _ int) uint {
			return brd.ID
		})}))
	})
	t.Run("Create Merchant", func(tt *testing.T) {
		mrcht = RandomMerchant()
		Err(tt, iSqlite().Create(&mrcht))
		tt.Log(Enc(map[string]any{"merchant_id": mrcht.ID}))
		iSMrcht = func() *gorm.DB {
			return iSqlite().Set("merchant_id", mrcht.ID)
		}
	})
	t.Run("Create SaleInfo", func(tt *testing.T) {
		saleList = lo.Map(breadList, func(brd *BreadProduct, _ int) *BreadSale {
			return RandomBreadSale(brd)
		})
		Err(tt, iSMrcht().Create(&saleList))
		tt.Log(Enc(saleList))
	})
	t.Run("Update SaleInfo", func(tt *testing.T) {
		err := iSMrcht().Transaction(func(tx *gorm.DB) error {
			for _, sale := range saleList {
				err := tx.Model(sale).Updates(map[string]any{
					"first_sale_time": time.Now(),
					"last_sale_time":  time.Now(),
				}).Error
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestParseTypes(t *testing.T) {
	for _, tCase := range []string{
		`{
	"TimeLine": "2025-03-22T22:21:50+08:00",
	"TimeOff": "2025-03-22T22:21:50+08:00",
	"TimeIn": "2025-03-22T22:21:50+08:00",
	"CostCent": "3.33",
	"BinSize": "10Mi"
}`,
		`{
	"TimeLine": 1742653292,
	"TimeOff": 1742653292,
	"TimeIn": 1742653292,
	"CostCent": "3.33",
	"BinSize": "10Mi"
}`,
	} {
		var cur struct {
			TimeLine types.Time
			TimeOff  types.Timestamp
			TimeIn   types.DateTime
			CostCent types.RMBCent
			BinSize  types.BinarySize
		}
		if err := json.Unmarshal([]byte(tCase), &cur); err != nil {
			t.Error(err)
		} else {
			t.Log(Enc(cur))
		}
	}

}
