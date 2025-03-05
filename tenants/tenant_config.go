package tenants

type Config struct {
	PluginName   string // default: gormx
	TagKey       string // default: gx
	TagTenantKey string // default: tenant
	TxTenantKey  string // default: tenant_id
	TxTenantsKey string // default: tenant_ids
	TxSkipKey    string // default: skip_tenant
}

func Default() *Config {
	return &Config{
		PluginName:   "gormx",
		TagKey:       "gx",
		TagTenantKey: "tenant",
		TxTenantKey:  "tenant_id",
		TxTenantsKey: "tenant_ids",
		TxSkipKey:    "skip_tenant",
	}
}
