package gormx_testv2

import (
	"encoding/json"
	"github.com/Juminiy/gormx/schemas/types"
	"github.com/Juminiy/gormx/tenants"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"testing"
	"time"
)

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
					"sale_count":      gofakeit.IntRange(100, 2000),
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
	t.Run("Count SaleInfo", func(tt *testing.T) {
		var cntSale int64
		Err(tt, iSMrcht().Table(`tbl_bread_sale`).
			Count(&cntSale))
		t.Log(cntSale)
	})
	t.Run("Pluck SaleReleaseCount", func(tt *testing.T) {
		var releaseCounts []int
		Err(tt, iSMrcht().Table(`tbl_bread_sale`).
			Pluck("release_count", &releaseCounts))
	})
	t.Run("Find saleSales", func(tt *testing.T) {
		var saleSales []struct {
			ID            uint
			FirstSaleTime time.Time
			LastSaleTime  time.Time
			SaleCount     int64
		}
		Err(tt, iSMrcht().Table(`tbl_bread_sale`).
			Find(&saleSales))
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
}`, `{
	"TimeLine": 1742653292,
	"TimeOff": 1742653292,
	"TimeIn": 1742653292,
	"CostCent": "3.33",
	"BinSize": "10Mi"
}`, `{
	"CostCent": "1",
	"BinSize": "920"
}`, `{
	"CostCent": "2.2",
	"BinSize": "4Ki"
}`, `{
	"CostCent": "2.03",
	"BinSize": "8Mi"
}`, `{
	"CostCent": "3.00",
	"BinSize": "10Gi"
}`, `{
	"CostCent": "5.30",
	"BinSize": "22Ti"
}`, `{
	"CostCent": "89.03",
	"BinSize": "22Pi"
}`, `{
	"Example": {
		"IntTyp": 114514,
		"RMap": {
			"Key":"Value",
			"V8": null
		}
	}
}`, `{
	"Example": null
}
`,
	} {
		var cur struct {
			TimeLine types.Time
			TimeOff  types.Timestamp
			TimeIn   types.DateTime
			CostCent types.RMBCent
			BinSize  types.BinarySize
			Example  types.ExampleTyp
		}
		if err := json.Unmarshal([]byte(tCase), &cur); err != nil {
			t.Error(err)
		} else {
			t.Log(Enc(cur))
		}
	}

}

func TestTenantsTypes(t *testing.T) {
	for _, tCase := range []string{
		`{
	"HideUID": 10,
	"UID": 20,
	"StrID": "30"
}`,
		`{}`,
		`{"HideUID": "20"}`, `{"HideUID": false}`, `{"HideUID": 3.33}`,
		`{"UID": "20"}`, `{"UID": false}`, `{"UID": 3.33}`,
		`{"StrID": 20}`, `{"StrID": false}`, `{"StrID": 3.33}`,
	} {
		var cur struct {
			HideUID tenants.HideID
			UID     tenants.ID
			StrID   tenants.SID
		}
		if err := json.Unmarshal([]byte(tCase), &cur); err != nil {
			t.Log(err.Error())
		} else {
			t.Log(Enc(cur))
		}
	}
}
