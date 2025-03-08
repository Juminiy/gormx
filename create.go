package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/gormx/tenants"
	"gorm.io/gorm"
)

func (cfg *Config) BeforeCreate(tx *gorm.DB) {
	if tx.Error != nil || callback.SkipCreate.OK(tx) {
		return
	}
	sCfg := cfg.OptionConfig(tx)

	if !sCfg.DisableFieldDup {
		cfg.UniquesCfg().FieldDupCheck(tx, false, sCfg.EnableComplexFieldDup)
		if tx.Error != nil {
			return
		}
	}

	if sCfg.BeforeCreateMapCallHooks {
		callback.BeforeCreateMapCallHook(tx)
	}

	callback.BeforeCreateSetDefaultValuesToMap(tx)

	cfg.TenantsSeq(tx)(func(tenant *tenants.Tenant) bool {
		deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
			tenant.Field.Name: tenant.Field.Value, // FieldName
		})
		return true
	})
}

func (cfg *Config) AfterCreate(tx *gorm.DB) {
	if tx.Error != nil || callback.SkipCreate.OK(tx) {
		return
	}
	sCfg := cfg.OptionConfig(tx)

	callback.AfterCreateSetAutoIncPkToMap(tx)

	if sCfg.AfterCreateMapCallHooks {
		callback.AfterCreateMapCallHook(tx)
	}

	if !sCfg.AfterCreateShowTenant {
		cfg.TenantsSeq(tx)(func(tenant *tenants.Tenant) bool {
			deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tenant.Field.Name: nil, // FieldName
			})
			return true
		})
	}
}
