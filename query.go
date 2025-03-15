package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/gormx/dynamicsql"
	"github.com/Juminiy/gormx/tenants"
	"gorm.io/gorm"
)

func (cfg *Config) BeforeQuery(tx *gorm.DB) {
	if tx.Error != nil || callback.SkipQuery.OK(tx) {
		return
	}
	sCfg := cfg.OptionConfig(tx)

	if sCfg.BeforeQueryOmitField {
		callback.BeforeQueryOmit(tx)
	}

	if sCfg.QueryDynamicSQL {
		dynamicsql.OmitEmptyClause(tx)
	}

	cfg.AddTenantClauses(tx, true)
}

func (cfg *Config) AfterQuery(tx *gorm.DB) {
	if tx.Error != nil || callback.SkipQuery.OK(tx) {
		return
	}
	sCfg := cfg.OptionConfig(tx)

	if sCfg.AfterFindMapCallHooks {
		callback.AfterFindMapCallHook(tx)
	}

	if !sCfg.AfterQueryShowTenant {
		cfg.TenantsSeq(tx)(func(tenant *tenants.Tenant) bool {
			deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tenant.Field.Name:   nil, // FieldName
				tenant.Field.DBName: nil, // DBName
			})
			return true
		})
	}
}
