package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/gormx/tenants"
	"gorm.io/gorm"
)

func (cfg *Config) BeforeQuery(tx *gorm.DB) {
	if tx.Error != nil {
		//|| _SkipQueryCallback.Ok(tx)
		return
	}
	if SessionConfig(cfg, tx).BeforeQueryOmitField {
		callback.BeforeQueryOmit(tx)
	}

	cfg.AddTenantClause(tx, true)
}

func (cfg *Config) AfterQuery(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := SessionConfig(cfg, tx)

	if sCfg.AfterFindMapCallHooks {
		callback.AfterFindMapCallHook(tx)
	}

	if !sCfg.AfterQueryShowTenant {
		if tInfo := tenants.Default().TenantInfo(tx); tInfo != nil {
			deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tInfo.Field.Name: nil, // FieldName
			})
		}
	}
}

func (cfg *Config) AddTenantClause(tx *gorm.DB, forQuery bool) {
	if tInfo := tenants.Default().TenantInfo(tx); tInfo != nil {
		tInfo.AddClause(tx)
		if !forQuery {
			deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tInfo.Field.DBName: nil,
				tInfo.Field.Name:   nil,
			})
			/*tx.Statement.Omit(tInfo.Field.DBName) // not effected*/
		}
	}
}
