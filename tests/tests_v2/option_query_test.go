package gormx_testv2

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/gormx/schemas/types"
	"github.com/Juminiy/gormx/tenants"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
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

type BreadHacker struct {
	VID        tenants.HideID `x:"merchant"`
	ID         uint
	CreateTime int64 `gorm:"autoCreateTime:milli"`
	UpdateTime int64 `gorm:"autoUpdateTime:milli"`
	soft_delete.DeletedAt
	SoftMin string `gorm:"default:$?"` // default in tag
	SoftAvg string // default in hooks
	SoftMax string
}

func (h *BreadHacker) BeforeCreate(tx *gorm.DB) error {
	if len(h.SoftAvg) == 0 {
		h.SoftAvg = "%^"
	}
	return nil
}

func TestPluckJSON(t *testing.T) {
	// stdlib not support???
	// sql: Scan called without calling Next
	pluckJSON := func(tt *testing.T, txF func() *gorm.DB) {
		var bread = RandomBread()
		Err(tt, txF().Create(&bread))
		var ingdt = map[string]int{}
		if err := txF().Model(&BreadProduct{Model: gorm.Model{ID: bread.ID}}).
			Pluck("ingredient", &ingdt).Error; err != nil {
			t.Logf("stdlib not support type: %s", err.Error())
		}
		t.Log(ingdt)
		var newBread = BreadProduct{Model: gorm.Model{ID: bread.ID}}
		Err(tt, txF().Select("ingredient").First(&newBread))
		t.Log(newBread.Ingredient)
	}
	pluckJSON(t, iSqlite0)
	pluckJSON(t, func() *gorm.DB {
		return iSqlite().Set(gormx.OptionKey, gormx.Option{PluckQueryByPkClause: true})
	})
}
