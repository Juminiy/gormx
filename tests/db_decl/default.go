package db_decl

import (
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/gormx/plugins"
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
	util.Must(plugins.OneError(
		tx.Use(&gormx.Config{
			PluginName:  "gormx",
			TagKey:      "mt",
			KnownModels: []any{},
			KnownScopes: map[string]string{
				"tenant": "tenant_id",
				"user":   "user_id",
			},
		})))
	tx.Default()
	return tx.DB.Debug()
}
