package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/deps"
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

	if tInfo := cfg.TenantsCfg().TenantInfo(tx); tInfo != nil {
		deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
			tInfo.Field.Name: tInfo.Field.Value, // FieldName
		})
	}
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
		if tInfo := cfg.TenantsCfg().TenantInfo(tx); tInfo != nil {
			deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tInfo.Field.Name: nil, // FieldName
			})
		}
	}
}
