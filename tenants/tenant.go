package tenants

import (
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Name         string // default: gormx:tenants
	TagKey       string // default: gormx
	TagTenantKey string // default: tenant
	TxTenantKey  string // default: tenant_id
	TxTenantsKey string // default: tenant_ids
	TxSkipKey    string // default: skip_tenant
}

type Tenant struct {
	Field schemas.Field
}

func (cfg *Config) TenantInfo(tx *gorm.DB) *Tenant {
	tenantInfoKey := util.StringJoin(":", cfg.Name, cfg.TagKey, cfg.TagTenantKey)
	if tInfo, ok := tx.Get(tenantInfoKey); ok {
		return tInfo.(*Tenant)
	}
	tInfo := cfg.tenantInfo(tx)
	if tInfo != nil {
		tx.Set(tenantInfoKey, tInfo)
	}
	return tInfo
}

func (cfg *Config) tenantInfo(tx *gorm.DB) *Tenant {
	tid, hastid := tx.Get(cfg.TxTenantKey)
	tids, hastids := tx.Get(cfg.TxTenantsKey)
	_, skiptid := tx.Get(cfg.TxSkipKey)
	if (!hastid && !hastids) || // tx no tenant_id or no tenant_ids set
		skiptid { // tx skip tenant_id and tenant_ids
		return nil
	}

	sch := tx.Statement.Schema
	if sch == nil { // no schema
		return nil
	}
	tidField, ok := lo.Find(sch.Fields, func(item *schema.Field) bool {
		if mt, ok := item.Tag.Lookup(cfg.TagKey); ok && mt == cfg.TagTenantKey {
			return true
		}
		return false
	})
	if !ok {
		return nil
	}

	field := schemas.FieldFromSchema(tidField)
	if hastid {
		field.Value = tid
	}
	if hastids {
		field.Values = deps.AS(tids)
	}
	return &Tenant{Field: field}
}
