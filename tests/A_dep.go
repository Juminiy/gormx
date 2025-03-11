package gormx_tests

import (
	"encoding/json"
	"errors"
	"github.com/Juminiy/gormx"
	mysql8 "github.com/Juminiy/gormx/tests/mysql_test"
	postgres17 "github.com/Juminiy/gormx/tests/postgres_test"
	sqlite3 "github.com/Juminiy/gormx/tests/sqlite_test"
	"github.com/Juminiy/gormx/uniques"
	"gorm.io/gorm"
	"testing"
)

var (
	isqlite = sqlite3.Orm()
	imysql  = mysql8.Orm()
	ipg     = postgres17.Orm()
)

func iSqlite() *gorm.DB {
	return isqlite.Debug()
}

func iMySQL() *gorm.DB {
	return imysql.Debug()
}

func iPg() *gorm.DB {
	return ipg.Debug()
}

var Enc = func(v any) string {
	bs, _ := json.MarshalIndent(v, "", "  ")
	return string(bs)
}
var Dec = func(s string, v any) {
	_ = json.Unmarshal([]byte(s), v)
}
var Err = func(t *testing.T, err error) {
	if err != nil {
		if uniques.IsFieldDupError(err) ||
			errors.Is(err, gormx.ErrNotAllowTenantGlobalUpdate) ||
			errors.Is(err, gormx.ErrNotAllowTenantGlobalDelete) ||
			errors.Is(err, gorm.ErrRecordNotFound) {
			t.Log(err)
		} else {
			t.Error(err)
		}
	}
}
