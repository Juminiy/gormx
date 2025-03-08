package gormx_tests

import (
	"encoding/csv"
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gorm.io/plugin/soft_delete"
	"math/rand/v2"
	"os"
	"testing"
	"time"
)

type BabyTrade struct {
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	soft_delete.DeletedAt
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	SimUUID   string `gorm:"primaryKey"`
	UserID    uint   `mt:"user"`
	TenantID  uint   `mt:"tenant"`
	AuctionID uint   `mt:"unique"`
	CatID     uint   `mt:"unique"`
	Cat       int
	BuyMount  int
	Day       string
}

func TestInit2(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(BabyTrade{}))
}

func Test2(t *testing.T) {
	csvFile, err := os.Open("testdata/baby_trade_copy_202503082022.csv")
	if err != nil {
		t.Error(err)
	}
	defer util.SilentCloseIO("csv file", csvFile)
	arr2d, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		t.Error(err)
	}
	Err(t, txFull().Create(util.New(lo.Map(arr2d[1:500], func(item []string, _ int) BabyTrade {
		return BabyTrade{
			SimUUID:   uuid.NewString(),
			UserID:    cast.ToUint(item[0]),
			AuctionID: cast.ToUint(item[1]),
			CatID:     cast.ToUint(item[2]),
			Cat:       cast.ToInt(item[3]),
			BuyMount:  cast.ToInt(item[4]),
			Day:       item[5],
		}
	}))).Error)
}

var sessionOpt = gormx.Option{
	DisableFieldDup:          false, // ✅
	EnableComplexFieldDup:    false, // ✅
	AllowTenantGlobalDelete:  false, // todo: simple-test
	BeforeDeleteDoQuery:      false, // todo: bugfix
	AllowTenantGlobalUpdate:  false, // todo: simple-test
	UpdateMapOmitZeroElemKey: false, // ✅
	UpdateMapOmitUnknownKey:  false, // ✅
	UpdateMapSetPkToClause:   false, // ✅
	AfterCreateShowTenant:    false,
	BeforeQueryOmitField:     false,
	AfterQueryShowTenant:     false,
	BeforeCreateMapCallHooks: false,
	AfterCreateMapCallHooks:  false,
	UpdateMapCallHooks:       false,
	AfterFindMapCallHooks:    false,
}

func TestNoTenant(t *testing.T) {
	Err(t, txPure().
		Set("user_id", 114514).
		Set("tenant_id", 114514).
		Set(gormx.OptionKey, gormx.Option{DisableFieldDup: true}).
		Create(&BabyTrade{
			SimUUID:   uuid.NewString(),
			AuctionID: rand.UintN(1024),
			CatID:     rand.UintN(4096),
			Cat:       rand.IntN(1),
			BuyMount:  rand.IntN(11030),
			Day:       "20250102",
		}).Error)
}

func TestCreateListDupOrNot(t *testing.T) {
	Err(t, txPure().
		Set(gormx.OptionKey, gormx.Option{EnableComplexFieldDup: true}).
		Create(&[]BabyTrade{
			{Day: "20140209", AuctionID: 1},
			{Day: "20240509", AuctionID: 1}},
		).Error)
}

func TestUpdateMapMixedMap(t *testing.T) {
	Err(t, txPure().
		Set(gormx.OptionKey, gormx.Option{
			UpdateMapOmitZeroElemKey: true,
			UpdateMapSetPkToClause:   true,
			UpdateMapOmitUnknownKey:  true,
		}).
		Table(`tbl_baby_trade`).
		Updates(map[string]any{
			"id":          1,                                      // Pk
			"auction_id":  10,                                     // set
			"cat_id":      11,                                     // set
			"missing_id":  24,                                     // unknown
			"unknown_id":  46,                                     // unknown
			"sim_uuid":    "bc06526c-57b1-4579-831a-8354b5672d87", // Pk, (zero, omit and do nothing)/(not-zero, as clause)
			"buy_mount":   0,                                      // zero
			"zero_id":     0,                                      // unknown,zero
			"zero_id_str": "",                                     // unknown,zero
		}).Error)
}

func TestUpdateMapMixedStruct(t *testing.T) {
	Err(t, txPure().
		Set(gormx.OptionKey, gormx.Option{
			UpdateMapOmitZeroElemKey: true,
			UpdateMapSetPkToClause:   true,
			UpdateMapOmitUnknownKey:  true,
		}).
		Table(`tbl_baby_trade`).
		Updates(&BabyTrade{
			ID:        1,  // Pk
			AuctionID: 10, // set
			CatID:     11, // set
			//"missing_id":  24,                                     // unknown
			//"unknown_id":  46,                                     // unknown
			SimUUID:  "bc06526c-57b1-4579-831a-8354b5672d87", // Pk, (zero, omit and do nothing)/(not-zero, as clause)
			BuyMount: 0,                                      // zero
			//"zero_id":     0,                                      // unknown,zero
			//"zero_id_str": "", // unknown,zero
		}).Error)
}
