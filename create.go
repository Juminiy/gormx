package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/gormx/tenants"
	"gorm.io/gorm"
)

func (cfg *Config) BeforeCreate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := SessionConfig(cfg, tx)

	if !sCfg.DisableFieldDup {
		(&schemas.Config{}).FieldDupCheck(tx, false)
		if tx.Error != nil {
			return
		}
	}

	if sCfg.BeforeCreateMapCallHooks {
		callback.BeforeCreateMapCallHook(tx)
	}

	callback.BeforeCreateSetDefaultValuesToMap(tx)

	if tInfo := tenants.Default().TenantInfo(tx); tInfo != nil {
		deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
			tInfo.Field.Name: tInfo.Field.Value, // FieldName
		})
	}
}

func (cfg *Config) AfterCreate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := SessionConfig(cfg, tx)

	callback.AfterCreateSetAutoIncPkToMap(tx)

	if sCfg.AfterCreateMapCallHooks {
		callback.AfterCreateMapCallHook(tx)
	}

	if !sCfg.AfterCreateShowTenant {
		if tInfo := tenants.Default().TenantInfo(tx); tInfo != nil {
			deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tInfo.Field.Name: nil, // FieldName
			})
		}
	}
}
