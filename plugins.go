package gormx

import (
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/gormx/tenants"
	"gorm.io/gorm"
)

func (cfg *Config) AddTenantClause(tx *gorm.DB, forQuery bool) {
	if tInfo := cfg.TenantsCfg().TenantInfo(tx); tInfo != nil {
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

func (cfg *Config) SchemasCfg() *schemas.Config {
	return cfg.schCfg
}

func (cfg *Config) TenantsCfg() *tenants.Config {
	return cfg.tetCfg
}
