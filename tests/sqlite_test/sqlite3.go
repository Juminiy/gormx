package sqlite3

import (
	"github.com/Juminiy/gormx/tests/db_decl"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Orm() *gorm.DB {
	return db_decl.Orm(sqlite.Open("kdb.db"))
}
