package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/deps"
	"gorm.io/gorm"
)

func (cfg *Config) BeforeQuery(tx *gorm.DB) {
	if tx.Error != nil || SkipQuery.OK(tx) {
		return
	}
	if cfg.OptionConfig(tx).BeforeQueryOmitField {
		callback.BeforeQueryOmit(tx)
	}

	cfg.AddTenantClause(tx, true)
}

func (cfg *Config) AfterQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := cfg.OptionConfig(tx)

	if sCfg.AfterFindMapCallHooks {
		callback.AfterFindMapCallHook(tx)
	}

	if !sCfg.AfterQueryShowTenant {
		if tInfo := cfg.TenantsCfg().TenantInfo(tx); tInfo != nil {
			deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tInfo.Field.Name: nil, // FieldName
			})
		}
	}
}
