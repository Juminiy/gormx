package gormx_testv2

import (
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/gormx/callback"
	mysql8 "github.com/Juminiy/gormx/tests/mysql_test"
	postgres17 "github.com/Juminiy/gormx/tests/postgres_test"
	sqlite3 "github.com/Juminiy/gormx/tests/sqlite_test"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"slices"
	"testing"
)

var (
	isqlite    = sqlite3.Orm()
	imysql     = mysql8.Orm()
	ipg        = postgres17.Orm()
	_ModelList = []any{&Order{}}
)

func iSqlite() *gorm.DB {
	return isqlite.Debug()
}

func iInnoDB() *gorm.DB {
	return imysql.Debug()
}

func iPG() *gorm.DB {
	return ipg.Debug()
}

func init() {
	slices.Values([]*gorm.DB{isqlite, imysql, ipg})(func(db *gorm.DB) bool {
		util.Must(db.Use(&gormx.Config{
			PluginName:  "gormx",
			TagKey:      "x",
			KnownModels: _ModelList,
			KnownScopes: map[string]string{
				"tenant":  "tenant_id",
				"user":    "user_id",
				"project": "project_id",
			},
		}))
		return true
	})
}

func TXAutoMigrate(iDb *gorm.DB) *gorm.DB {
	return callback.SkipRowRaw.Set(iDb)
}

func TestAAAInit(t *testing.T) {
	slices.Values([]*gorm.DB{iSqlite(), iInnoDB(), iPG()})(func(db *gorm.DB) bool {
		util.Must(TXAutoMigrate(db).AutoMigrate(_ModelList...))
		return true
	})
}
