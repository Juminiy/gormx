package gormx

import (
	"github.com/Juminiy/gormx/optlock"
	"github.com/Juminiy/gormx/plugins"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/gormx/tenants"
	"github.com/Juminiy/gormx/uniques"
	"github.com/samber/lo"
	expmaps "golang.org/x/exp/maps"
	"gorm.io/gorm"
)

type Config struct {
	PluginName  string            // no default value, "" will be error
	TagKey      string            // default: gormx
	KnownModels []any             // must know your schemas(models, tables), or plugins will be folly
	KnownScopes map[string]string // must know your isolation scope (tenant,user) fieldTag -> txKey, or plugins will do nothing on scopes

	Option *Option // config Option can be overwritten by session, default all false

	schCfg     *schemas.Config            // schemas.Config
	unqCfg     *uniques.Config            // uniques.Config
	tetCfg     map[string]*tenants.Config // tenants.Config
	optLockCfg *optlock.Config            // optlock.Config
}

func (cfg *Config) Name() string {
	return cfg.PluginName
}

func (cfg *Config) Initialize(tx *gorm.DB) error {
	if len(cfg.PluginName) == 0 {
		return plugins.NoPluginName
	}
	if len(cfg.TagKey) == 0 {
		cfg.TagKey = "gormx"
	}

	cfg.Option = &Option{}

	cfg.schCfg = &schemas.Config{
		Name:   cfg.PluginName + ":schemas",
		TagKey: cfg.TagKey,
	}
	cfg.unqCfg = (&uniques.Config{
		Name:         cfg.PluginName + ":uniques",
		TagKey:       cfg.TagKey,
		TagUniqueKey: "unique",
		TxKeys:       expmaps.Values(cfg.KnownScopes),
	}).Initial()
	cfg.tetCfg = lo.MapValues(cfg.KnownScopes, func(txKey string, fieldTag string) *tenants.Config {
		return &tenants.Config{
			Name:         cfg.PluginName + "_scope:" + fieldTag,
			TagKey:       cfg.TagKey,
			TagTenantKey: fieldTag,
			TxTenantKey:  txKey,
			TxTenantsKey: txKey + "_list",
			TxSkipKey:    "skip_" + txKey,
		}
	})
	cfg.optLockCfg = &optlock.Config{
		Name:       cfg.PluginName + ":optlock",
		TagKey:     cfg.TagKey,
		TagOptLock: "version",
	}
	cfg.SchemasCfg().GraspSchema(tx, cfg.KnownModels...)
	return plugins.OneError(
		tx.Callback().Create().Before("gorm:before_create").
			Register(plugins.CallbackName(cfg.PluginName, true, 'C'), cfg.BeforeCreate),
		tx.Callback().Create().After("gorm:after_create").
			Register(plugins.CallbackName(cfg.PluginName, false, 'C'), cfg.AfterCreate),

		tx.Callback().Query().Before("gorm:query").
			Register(plugins.CallbackName(cfg.PluginName, true, 'Q'), cfg.BeforeQuery),
		tx.Callback().Query().After("gorm:after_query").
			Register(plugins.CallbackName(cfg.PluginName, false, 'Q'), cfg.AfterQuery),

		tx.Callback().Update().Before("gorm:setup_reflect_value").
			Register(plugins.CallbackName(cfg.PluginName, true, 'U'), cfg.BeforeUpdate),
		tx.Callback().Update().After("gorm:after_update").
			Register(plugins.CallbackName(cfg.PluginName, false, 'U'), cfg.AfterUpdate),

		tx.Callback().Delete().Before("gorm:before_delete").
			Register(plugins.CallbackName(cfg.PluginName, true, 'D'), cfg.BeforeDelete),
		/*tx.Callback().Delete().After("gorm:after_delete").
		Register(plugins.CallbackName(cfg.PluginName, false, 'D'), cfg.AfterDelete),*/

		tx.Callback().Row().Before("gorm:row").
			Register(plugins.CallbackName(cfg.PluginName, true, 'R'), cfg.BeforeRowOrRaw),
		tx.Callback().Row().After("gorm:row").
			Register(plugins.CallbackName(cfg.PluginName, false, 'R'), cfg.AfterRow),
		tx.Callback().Raw().Before("gorm:raw").
			Register(plugins.CallbackName(cfg.PluginName, true, 'E'), cfg.BeforeRowOrRaw),
	)
}

type Option struct {
	DisableFieldDup          bool // effect on create and update
	EnableComplexFieldDup    bool // effect on create
	AfterCreateShowTenant    bool // effect on create
	BeforeCreateMapCallHooks bool // effect on before create map
	AfterCreateMapCallHooks  bool // effect on after create map, it's a low efficiency option, not to open it unless strongly need

	AllowTenantGlobalDelete bool // effect on delete, if false: raise error when clause only have (tenant) and (soft_delete)
	BeforeDeleteReturning   bool // effect on delete, use with clause.Returning, when database not support Returning

	AllowTenantGlobalUpdate  bool // effect on update, if false: raise error when clause only have (tenant) and (soft_delete)
	UpdateMapOmitZeroElemKey bool // effect on update map
	UpdateMapOmitUnknownKey  bool // effect on update map
	UpdateMapSetPkToClause   bool // effect on update map
	UpdateMapCallHooks       bool // effect on update map
	UpdateOptimisticLock     bool // effect on update
	AfterUpdateReturning     bool // effect on update, use with clause.Returning, when database not support Returning

	BeforeQueryOmitField bool // effect on query, use with tag `gorm:"->:false"`
	AfterQueryShowTenant bool // effect on query
	QueryDynamicSQL      bool // effect on query
	ExplainQueryOrRow    bool // effect on query or row
	PluckQueryByPkClause bool // effect on query one column
	// Deprecated
	AfterFindMapCallHooks bool // effect on query map, after the evaluation, it's not a common and general case, but also to waste of time

	WriteClauseToRowOrRaw bool // effect on row or raw
}

const OptionKey = "session:option_config"

func (cfg *Config) OptionConfig(tx *gorm.DB) Option {
	cfg.SchemasCfg().ParseSchema(tx)
	if v, ok := tx.Get(OptionKey); ok {
		if vRecv, ok := v.(Option); ok {
			return vRecv
		} else if pRecv, ok := v.(*Option); ok && pRecv != nil {
			return *pRecv
		}
	}
	return *cfg.Option
}
