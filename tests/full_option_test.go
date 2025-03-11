package gormx_tests

import (
	"encoding/csv"
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/gormx/plugins"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
	"math/rand/v2"
	"os"
	"slices"
	"testing"
	"time"
)

/*func init() {
	var modelList = []any{&BabyTrade{}, &Consumer{}, &Product{}, &CalicoWeave{}, &AppUser{}}

	iSqlite().Plugins["gormx"].(*gormx.Config).
		SchemasCfg().
		GraspSchema(iSqlite(), modelList...)

	iMySQL().Plugins["gormx"].(*gormx.Config).
		SchemasCfg().
		GraspSchema(iMySQL(), modelList...)

	iPg().Plugins["gormx"].(*gormx.Config).
		SchemasCfg().
		GraspSchema(iPg(), modelList...)

}*/

func init() {
	var modelList = []any{&BabyTrade{}, &Consumer{}, &Product{}, &CalicoWeave{}, &AppUser{}}

	slices.Values([]*gorm.DB{isqlite, imysql, ipg})(func(db *gorm.DB) bool {
		util.Must(db.Use(&gormx.Config{
			PluginName:  "gormx",
			TagKey:      "mt",
			KnownModels: modelList,
			KnownScopes: map[string]string{
				"tenant":  "tenant_id",
				"user":    "user_id",
				"project": "project_id",
			},
		}))
		return true
	})
}

func TestInit(t *testing.T) {
	var modelList = []any{&BabyTrade{}, &Consumer{}, &Product{}, &CalicoWeave{}, &AppUser{}}

	util.Must(plugins.OneError(
		txMigrate(iSqlite()).AutoMigrate(modelList...),
		txMigrate(iMySQL()).AutoMigrate(modelList...),
		txMigrate(iPg()).AutoMigrate(modelList...),
	))
}

type BabyTrade struct {
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	soft_delete.DeletedAt
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	SimUUID   string `gorm:"primaryKey" mt:"unique"`
	UserID    uint   `mt:"user"`
	TenantID  uint   `mt:"tenant"`
	AuctionID uint   `mt:"unique:auc_cat"`
	CatID     uint   `mt:"unique:auc_cat,cat"`
	Cat       int    `mt:"unique:cat"`
	BuyMount  int    `gorm:"->:false;<-"` // query omit
	Day       string `gorm:"default:2024-02-05"`
}

func (t *BabyTrade) BeforeCreate(tx *gorm.DB) error {
	t.BuyMount = 666666
	tx.Logger.Info(tx.Statement.Context, "you are in before create hooks")
	return nil
}

func (t *BabyTrade) AfterCreate(tx *gorm.DB) error {
	t.BuyMount = 88888888
	tx.Logger.Info(tx.Statement.Context, "you are in after create hooks")
	return nil
}

func (t *BabyTrade) BeforeUpdate(tx *gorm.DB) error {
	t.BuyMount = 666666
	tx.Logger.Info(tx.Statement.Context, "you are in before update hooks")
	return nil
}

func (t *BabyTrade) AfterUpdate(tx *gorm.DB) error {
	t.BuyMount = 88888888
	tx.Logger.Info(tx.Statement.Context, "you are in after update hooks")
	return nil
}

func (t *BabyTrade) BeforeDelete(tx *gorm.DB) error {
	t.BuyMount = 666666
	tx.Logger.Info(tx.Statement.Context, "you are in before delete hooks")
	return nil
}

func (t *BabyTrade) AfterDelete(tx *gorm.DB) error {
	t.BuyMount = 88888888
	tx.Logger.Info(tx.Statement.Context, "you are in after delete hooks")
	return nil
}

func (t *BabyTrade) AfterFind(tx *gorm.DB) error {
	t.BuyMount = 88888888
	tx.Logger.Info(tx.Statement.Context, "you are in after find hooks")
	return nil
}

func (t *BabyTrade) RandomSet() *BabyTrade {
	t.SimUUID = uuid.NewString()
	t.AuctionID = rand.UintN(1 << 10)
	t.CatID = rand.UintN(1 << 4)
	t.Cat = rand.IntN(2)
	t.BuyMount = rand.IntN(1 << 16)
	t.Day = gofakeit.Date().String()
	return t
}

func (t *BabyTrade) HardCodeSet() *BabyTrade {
	t.SimUUID = _HardCodeSim
	t.AuctionID = 888
	t.CatID = 888888
	t.Cat = 6
	t.BuyMount = 666
	t.Day = gofakeit.Date().String()
	return t
}

func BabyTradeMapRandom() map[string]any {
	return map[string]any{
		"SimUUID":   uuid.NewString(),
		"AuctionID": rand.UintN(1 << 10),
		"CatID":     rand.UintN(1 << 9),
		"Cat":       rand.IntN(1 << 8),
		"BuyMount":  rand.IntN(1 << 12),
		"Day":       gofakeit.Date().String(),
	}
}

func BabyTradeMapHardCode() map[string]any {
	return map[string]any{
		"SimUUID":   _HardCodeSim,
		"AuctionID": 888,
		"CatID":     888888,
		"Cat":       6,
		"BuyMount":  666,
		"Day":       gofakeit.Date().String(),
	}
}

var _HardCodeSim = uuid.NewString()

func TestInitBabyTrade(t *testing.T) {
	//Err(t, txMigrate().AutoMigrate(&BabyTrade{}))
}

func TestBatchCreate(t *testing.T) {
	csvFile, err := os.Open("testdata/baby_trade.csv")
	if err != nil {
		t.Error(err)
	}
	defer util.SilentCloseIO("csv file", csvFile)
	arr2d, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		t.Error(err)
	}
	startT := time.Now()
	Err(t, txMysql().CreateInBatches(util.New(lo.Map(arr2d, func(item []string, _ int) BabyTrade {
		return BabyTrade{
			SimUUID:   uuid.NewString(),
			UserID:    cast.ToUint(item[0]),
			AuctionID: cast.ToUint(item[1]),
			CatID:     cast.ToUint(item[2]),
			Cat:       cast.ToInt(item[3]),
			BuyMount:  cast.ToInt(item[4]),
			Day:       item[5],
		}
	})), 1024).Error)
	t.Logf("time dur: %f", time.Now().Sub(startT).Seconds())
}

var sessionOpt = gormx.Option{
	DisableFieldDup:          false,
	EnableComplexFieldDup:    false,
	AfterCreateShowTenant:    false,
	BeforeCreateMapCallHooks: false,
	AfterCreateMapCallHooks:  false,
	AllowTenantGlobalDelete:  false,
	BeforeDeleteReturning:    false,
	AllowTenantGlobalUpdate:  false,
	UpdateMapOmitZeroElemKey: false,
	UpdateMapOmitUnknownKey:  false,
	UpdateMapSetPkToClause:   false,
	UpdateMapCallHooks:       false,
	AfterUpdateReturning:     false,
	BeforeQueryOmitField:     false,
	AfterQueryShowTenant:     false,
	AfterFindMapCallHooks:    false,
	QueryDynamicSQL:          false,
	WriteClauseToRowOrRaw:    false,
}

func TestNoTenant(t *testing.T) {
	Err(t, txMysql().
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

func txScope() *gorm.DB {
	return txPure().
		Set("user_id", 114514).
		Set("tenant_id", 114514).
		Set(gormx.OptionKey, gormx.Option{
			DisableFieldDup:          false,
			EnableComplexFieldDup:    true,
			AfterCreateShowTenant:    true,
			BeforeCreateMapCallHooks: true,
			AfterCreateMapCallHooks:  true,
			UpdateMapSetPkToClause:   true,
			UpdateMapCallHooks:       true,
		})
}

func TestCreateOrFindShowTenantOption(t *testing.T) {
	createBt := BabyTrade{
		SimUUID:   uuid.NewString(),
		AuctionID: rand.UintN(1 << 10),
		CatID:     rand.UintN(1 << 4),
		Cat:       rand.IntN(2),
		BuyMount:  rand.IntN(1 << 16),
		Day:       gofakeit.Date().String(),
	}
	Err(t, txScope().Create(&createBt).Error)
	t.Log(Enc(createBt))

	findBt := BabyTrade{ID: createBt.ID}
	Err(t, txScope().First(&findBt).Error)
	t.Log(Enc(findBt))
}

func TestUpdateDeleteScope(t *testing.T) {
	Err(t, txScope().
		Table(`tbl_baby_trade`).
		Update("cat_id", rand.UintN(1<<8)).Error)

	Err(t, txScope().Delete(&BabyTrade{}).Error)
}

func TestFieldDupOneStruct(t *testing.T) {
	// no scope
	Err(t, txPure().Create((&BabyTrade{}).RandomSet()).Error)

	// one scope
	Err(t, txPure().Set("user_id", 114514).
		Create((&BabyTrade{}).RandomSet()).Error)

	// two scope
	Err(t, txPure().Set("user_id", 114514).
		Set("tenant_id", 114514).
		Create((&BabyTrade{}).RandomSet()).Error)
}

func TestFieldDupOneMap(t *testing.T) {
	// no scope
	Err(t, txPure().Table(`tbl_baby_trade`).Create(BabyTradeMapRandom()).Error)

	// one scope
	Err(t, txPure().Set("user_id", 114514).
		Table(`tbl_baby_trade`).Create(BabyTradeMapRandom()).Error)

	// two scope
	Err(t, txPure().Set("user_id", 114514).
		Set("tenant_id", 114514).
		Table(`tbl_baby_trade`).Create(BabyTradeMapRandom()).Error)
}

func complexFieldDupOpt() *gorm.DB {
	return txPure().Set(gormx.OptionKey, gormx.Option{
		DisableFieldDup:          false,
		EnableComplexFieldDup:    true,
		BeforeCreateMapCallHooks: true,
		AfterCreateMapCallHooks:  true,
	})
}

func TestGormNotSupportType(t *testing.T) {
	// panic: sql: Scan error on column index 0,
	//name "id": unsupported Scan,
	//storing driver.Value type int64 into type *gormx_tests.BabyTrade;
	//sql: Scan error on column index 0, name "id": unsupported Scan,
	//storing driver.Value type int64 into type *gormx_tests.BabyTrade [recovered]
	/*var structList = &[]**BabyTrade{
		util.New((&BabyTrade{}).Set()),
		util.New((&BabyTrade{}).Set()),
	}
	Err(t, complexFieldDupOpt().Create(&structList).Error)
	t.Log(Enc(structList))*/
}

func TestFieldDupStructList(t *testing.T) {
	var structList = &[]*BabyTrade{
		(&BabyTrade{}).HardCodeSet(),
		(&BabyTrade{}).HardCodeSet(),
	}
	var slPtr = util.New2(structList)
	Err(t, complexFieldDupOpt().Create(&slPtr).Error)
	t.Log(Enc(structList))
}

func TestFieldDupMapList(t *testing.T) {
	var mapList = []map[string]any{
		BabyTradeMapHardCode(),
		BabyTradeMapHardCode(),
	}
	Err(t, complexFieldDupOpt().Table(`tbl_baby_trade`).Create(&mapList).Error)
	t.Log(Enc(mapList))
}

func TestFieldDupWithTenantsCreateMapCallHooks(t *testing.T) {
	var mapList = []map[string]any{
		BabyTradeMapRandom(),
		BabyTradeMapRandom(),
	}
	Err(t, txScope().Table(`tbl_baby_trade`).Create(&mapList).Error)
	t.Log(Enc(mapList))
}

func TestUpdateMapCallHooks(t *testing.T) {
	Err(t, txScope().Table(`tbl_baby_trade`).Updates(map[string]any{
		"id":         2,
		"auction_id": 10,
	}).Error)
}

func TestDeleteMapCallHooks(t *testing.T) {
	var deleteBt BabyTrade
	Err(t, txScope().Clauses(clause.Returning{}).Delete(&deleteBt, 21).Error)
	t.Log(Enc(deleteBt))
}
