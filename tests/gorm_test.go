package gorm_api

import (
	"encoding/json"
	"errors"
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/gormx/schemas"
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
		PluginName: "gormx",
		TagKey:     "mt",
	}))
	util.Must(tx.Use(&clauses.Config{
		PluginName:                 "clause_checker",
		AllowWriteClauseToRawOrRow: true,
		BeforePlugins:              []string{"gormx"},
	}))
	tx.Plugins["gormx"].(*gormx.Config).SchemasCfg().
		GraspSchema(tx.DB, &Product{}, &AppUser{}, &Consumer{}, &CalicoWeave{})
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
		if schemas.IsFieldDupError(err) ||
			errors.Is(err, gormx.ErrNotAllowTenantGlobalUpdate) ||
			errors.Is(err, gormx.ErrNotAllowTenantGlobalDelete) ||
			errors.Is(err, gorm.ErrRecordNotFound) {
			t.Log(err)
		} else {
			util.Must(err)
		}
	}
}

func TestInit(t *testing.T) {}
