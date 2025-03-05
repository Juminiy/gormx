package gorm_api

import (
	"encoding/json"
	"errors"
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/kube/pkg/storage_api/gorm_api/multi_tenants"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

var _tx *DB

func init() {
	tx, err := New(gorm.Config{
		Dialector: sqlite.Open("kdb.db"),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:         "tbl_",
			SingularTable:       true,
			NameReplacer:        nil,
			NoLowerCase:         false,
			IdentifierMaxLength: 255,
		},
	})
	util.Must(err)
	util.Must(tx.Use(&gormx.Config{
		PluginName: "multi_tenants",
	}))
	util.Must(tx.Use(&clauses.Config{
		PluginName:                 "clause_checker",
		AllowWriteClauseToRawOrRow: true,
		BeforePlugins:              []string{"multi_tenants"},
	}))
	/*tx.Plugins["multi_tenants"].(*gormx.Config).
	GraspSchema(tx.DB, &Product{}, &AppUser{}, &Consumer{}, &CalicoWeave{})*/
	tx.DB = tx.Debug()
	_tx = tx
}

var Enc = func(v any) string {
	bs, _ := json.MarshalIndent(v, " ", "")
	return string(bs)
}
var Dec = func(s string, v any) {
	_ = json.Unmarshal([]byte(s), v)
}

var Err = func(t *testing.T, err error) {
	if err != nil {
		if multi_tenants.IsFieldDupError(err) ||
			errors.Is(err, multi_tenants.ErrDeleteTenantAllNotAllowed) ||
			errors.Is(err, multi_tenants.ErrUpdateTenantAllNotAllowed) ||
			errors.Is(err, gorm.ErrRecordNotFound) {
			t.Log(err)
		} else {
			util.Must(err)
		}
	}
}

func TestInit(t *testing.T) {}
