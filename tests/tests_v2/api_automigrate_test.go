package gormx_testv2

import (
	"encoding/json"
	"errors"
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/gormx/callback"
	sqlite3 "github.com/Juminiy/gormx/tests/sqlite_test"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"slices"
	"testing"
)

var (
	isqlite  = sqlite3.Orm() // isPlugin
	isqlite0 = sqlite3.Orm() // noPlugin
	//imysql   = mysql8.Orm()  // isPlugin
	//imysql0  = mysql8.Orm()  // noPlugin
	//ipg        = postgres17.Orm()
	_ModelList = []any{&Order{}, &Product{}, &ChuckBlock{}, &BreadProduct{}, &BreadSale{}, &BreadMerchant{}, &BreadHacker{}}
)

// no plugin
func iSqlite0() *gorm.DB {
	return isqlite0.Debug()
}

// raw plugin
func iSqlite() *gorm.DB {
	return isqlite.Debug()
}

// raw plugin with tenants value
func iSUser() *gorm.DB { return iSqlite().Set("user_id", 666) }

func iSUserTenant() *gorm.DB {
	return iSqlite().Set("user_id", 666).Set("tenant_id", 888)
}

// mysql
/*func iInnoDB() *gorm.DB {
	return imysql.Debug()
}

func iInnoDB0() *gorm.DB {
	return imysql0.Debug()
}*/

// postgresql
/*func iPG() *gorm.DB {
	return ipg.Debug()
}*/

func init() {
	slices.Values([]*gorm.DB{isqlite /*imysql*/ /* ipg*/})(func(db *gorm.DB) bool {
		util.Must(db.Use(&gormx.Config{
			PluginName:  "gormx",
			TagKey:      "x",
			KnownModels: _ModelList,
			KnownScopes: map[string]string{
				"tenant":   "tenant_id",
				"user":     "user_id",
				"project":  "project_id",
				"merchant": "merchant_id",
			},
		}))
		return true
	})
}

func TXAutoMigrate(iDb *gorm.DB) *gorm.DB {
	return callback.SkipRowRaw.Set(iDb)
}

func TestAAAInit(t *testing.T) {
	slices.Values([]*gorm.DB{iSqlite() /*iInnoDB(), iPG()*/})(func(db *gorm.DB) bool {
		util.Must(TXAutoMigrate(db).AutoMigrate(_ModelList...))
		return true
	})
}

func Err(t *testing.T, tx *gorm.DB) {
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			t.Log("not found any record, please insert some record")
			return
		}
		t.Error(err)
	}
}

func Enc(i any) string {
	b, err := json.MarshalIndent(i, "", "  ")
	util.Must(err)
	return string(b)
}
