package gormx_testv2

import (
	"errors"
	"github.com/Juminiy/gormx"
	"gorm.io/gorm"
	"testing"
)

var CComplexCheck = gormx.Option{EnableComplexFieldDup: true}
var ErrUniquesNotPreCheckList = errors.New("uniques plugin create not pre check list")

func TestCreateListPreCheck(t *testing.T) {
	t.Run("NoPlugin", func(tt *testing.T) {
		order := RandomOrder()
		orders := []*Order{order, order, order}
		Err(tt, iSqlite0().Create(&orders))
	})
	t.Run("IsPlugin, NoCheck, NoScopes", func(tt *testing.T) {
		order := RandomOrder()
		orders := []*Order{order, order, order}
		Err(tt, iSqlite().Create(&orders))
	})
	t.Run("IsPlugin, IsCheck, NoScopes", func(tt *testing.T) {
		order := RandomOrder()
		orders := []*Order{order, order, order}
		if err := iSqlite().Set(gormx.OptionKey, CComplexCheck).Create(&orders).Error; err == nil {
			tt.Error(ErrUniquesNotPreCheckList)
		} else {
			tt.Log(err.Error())
		}
	})
	t.Run("IsPlugin, IsCheck, Scopes(user_id)", func(tt *testing.T) {
		order := RandomOrder()
		orders := []*Order{order, order, order}
		if err := iSUser().Set(gormx.OptionKey, CComplexCheck).Create(&orders).Error; err == nil {
			tt.Error(ErrUniquesNotPreCheckList)
		} else {
			tt.Log(err.Error())
		}
	})
	t.Run("IsPlugin, IsCheck, Scopes(user_id, tenant_id)", func(tt *testing.T) {
		order := RandomOrder()
		orders := []*Order{order, order, order}
		if err := iSUserTenant().Set(gormx.OptionKey, CComplexCheck).Create(&orders).Error; err == nil {
			tt.Error(ErrUniquesNotPreCheckList)
		} else {
			tt.Log(err.Error())
		}
	})
}

// ExtremeScenarioTestingCast
func TestCreateScopeHideID(t *testing.T) {
	var iSMrcht func() *gorm.DB
	t.Run("set merchant", func(tt *testing.T) {
		mrcht := RandomMerchant()
		Err(tt, iSqlite().Create(&mrcht))
		iSMrcht = func() *gorm.DB {
			return iSqlite().Set("merchant_id", mrcht.ID).
				Set(gormx.OptionKey, gormx.Option{
					DisableFieldDup:          false,
					EnableComplexFieldDup:    true,
					AfterCreateShowTenant:    true,
					BeforeCreateMapCallHooks: true,
					AfterCreateMapCallHooks:  true,
				})
		}
	})
	t.Run("Create Struct Nil", func(tt *testing.T) {
		var hkr *BreadHacker
		Err(tt, iSMrcht().Create(&hkr))
		tt.Log(Enc(hkr))
	})
	t.Run("Create Struct Zero", func(tt *testing.T) {
		var hkr BreadHacker
		Err(tt, iSMrcht().Create(&hkr))
		tt.Log(Enc(hkr))
	})
	t.Run("Create StructList Nil", func(tt *testing.T) {
		var hkrs []BreadHacker
		if err := iSMrcht().Create(&hkrs).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrs))
	})
	t.Run("Create StructList LenIsZero", func(tt *testing.T) {
		var hkrs = []BreadHacker{}
		if err := iSMrcht().Create(&hkrs).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrs))
	})
	t.Run("Create StructList", func(tt *testing.T) {
		var hkrs = []BreadHacker{{}, {}, {}}
		Err(tt, iSMrcht().Create(&hkrs))
		tt.Log(Enc(hkrs))
	})
	t.Run("Create StructArray LenIsZero", func(tt *testing.T) {
		var hkrs [0]BreadHacker
		if err := iSMrcht().Create(&hkrs).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrs))
	})
	t.Run("Create StructArray", func(tt *testing.T) {
		var hkrs [3]BreadHacker
		Err(tt, iSMrcht().Create(&hkrs))
		tt.Log(Enc(hkrs))
	})
	t.Run("Create MapList Nil", func(tt *testing.T) {
		var hkrMapList []map[string]any
		if err := iSMrcht().Table(`tbl_bread_hacker`).
			Create(&hkrMapList).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrMapList))
	})
	t.Run("Create MapList LenIsZero", func(tt *testing.T) {
		var hkrMapList = []map[string]any{}
		if err := iSMrcht().Table(`tbl_bread_hacker`).
			Create(&hkrMapList).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrMapList))
	})
	t.Run("Create MapList", func(tt *testing.T) {
		var hkrMapList = []map[string]any{{}, {}, {}}
		if err := iSMrcht().Table(`tbl_bread_hacker`).
			Create(&hkrMapList).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrMapList))
	})
}

func testCreateMapExtremeCases(t *testing.T, txFn func() *gorm.DB) {
	// plugin create ok
	t.Run("Create PtrToZeroMap", func(tt *testing.T) {
		var hkrMap = map[string]any{}
		if err := txFn().Table(`tbl_bread_hacker`).
			Create(&hkrMap).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrMap))
	})
	// plugin create ok
	t.Run("Create ZeroMap", func(tt *testing.T) {
		var hkrMap = map[string]any{}
		if err := txFn().Table(`tbl_bread_hacker`).
			Create(hkrMap).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrMap))
	})
	// gorm not ok, plugin create ok
	t.Run("Create PtrToNilMap", func(tt *testing.T) {
		var hkrMap map[string]any
		if err := txFn().Table(`tbl_bread_hacker`).
			Create(&hkrMap).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrMap))
	})
	// gorm not support; throw error, so plugin BeforeCreate is not executed
	/*t.Run("Create NilPtrToNilMap", func(tt *testing.T) {
		var hkrMap *map[string]any
		if err := txFn().Table(`tbl_bread_hacker`).
			Create(hkrMap).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrMap))
	})*/
	// gorm not support; plugin not support as well; go syntax not support
	/*t.Run("Create NilMap", func(tt *testing.T) {
		var hkrMap map[string]any
		if err := txFn().Table(`tbl_bread_hacker`).
			Create(hkrMap).Error; err != nil {
			tt.Log(tt.Name(), err.Error())
			return
		}
		tt.Log(Enc(hkrMap))
	})*/
}

func TestCreateMapExtremeCases(t *testing.T) {
	var iSMrcht func() *gorm.DB
	t.Run("set merchant", func(tt *testing.T) {
		mrcht := RandomMerchant()
		Err(tt, iSqlite().Create(&mrcht))
		iSMrcht = func() *gorm.DB {
			return iSqlite().Set("merchant_id", mrcht.ID).
				Set(gormx.OptionKey, gormx.Option{
					DisableFieldDup:          false,
					EnableComplexFieldDup:    true,
					AfterCreateShowTenant:    true,
					BeforeCreateMapCallHooks: true,
					AfterCreateMapCallHooks:  true,
				})
		}
	})
	//testCreateMapExtremeCases(t, iSqlite0)
	testCreateMapExtremeCases(t, iSMrcht)
}
