package gormx

import (
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/gormx/tenants"
	"github.com/Juminiy/gormx/uniques"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"iter"
	"slices"
)

func (cfg *Config) SchemasCfg() *schemas.Config {
	return cfg.schCfg
}

func (cfg *Config) UniquesCfg() *uniques.Config {
	return cfg.unqCfg
}

func (cfg *Config) TenantsCfg() map[string]*tenants.Config {
	return cfg.tetCfg
}

func (cfg *Config) TenantsSeq(tx *gorm.DB) iter.Seq[*tenants.Tenant] {
	return slices.Values(lo.Filter(
		lo.MapToSlice(cfg.TenantsCfg(), func(fieldTag string, tetCfg *tenants.Config) *tenants.Tenant {
			return tetCfg.TenantInfo(tx)
		}),
		func(item *tenants.Tenant, _ int) bool {
			return item != nil
		},
	))
}

func (cfg *Config) AddTenantClauses(tx *gorm.DB, forQuery bool) {
	cfg.TenantsSeq(tx)(func(tenant *tenants.Tenant) bool {
		tenant.AddClause(tx)
		if !forQuery {
			deps.Ind(tx.Statement.ReflectValue).SetField(map[string]any{
				tenant.Field.DBName: nil,
				tenant.Field.Name:   nil,
			})
		}
		return true
	})
}
