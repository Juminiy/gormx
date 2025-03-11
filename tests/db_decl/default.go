package db_decl

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Orm(dialector gorm.Dialector) *gorm.DB {
	tx, err := New(gorm.Config{
		Dialector: dialector,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:         "tbl_",
			SingularTable:       true,
			NameReplacer:        nil,
			NoLowerCase:         false,
			IdentifierMaxLength: 255,
		},
		PrepareStmt: true,
	})
	util.Must(err)
	return tx.Default().Session(&gorm.Session{NewDB: true})
}
