package clauses

import (
	"github.com/Juminiy/gormx/plugins"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"slices"
)

type Config struct {
	PluginName string

	AllowWriteClauseToRawOrRow    bool // effect on raw and row where clause, suggest to Wrap sql by: SELECT * FROM (your sql string)
	CheckAndOmitNotExistingColumn bool // effect on clause.OrderBy, clause.OrderByColumn

	BeforePlugins []string
	AfterPlugins  []string
}

func (cfg *Config) Name() string {
	return cfg.PluginName
}

func (cfg *Config) Initialize(tx *gorm.DB) error {
	if len(cfg.PluginName) == 0 {
		return plugins.NoPluginName
	}

	if registerRawRowErr := plugins.OneError(
		tx.Callback().Raw().Before("gorm:raw").
			Register(plugins.CallbackName(cfg.PluginName, true, 'E'), cfg.RowRawClause),

		tx.Callback().Row().Before("gorm:row").
			Register(plugins.CallbackName(cfg.PluginName, true, 'R'), cfg.RowRawClause),
	); registerRawRowErr != nil {
		return registerRawRowErr
	}

	var registerBeforeErr error
	var hasBefore bool
	slices.All(cfg.BeforePlugins)(func(_ int, s string) bool {
		if util.MapOk(tx.Plugins, s) {
			hasBefore = true
			registerBeforeErr = plugins.OneError(
				tx.Callback().Delete().
					Before(plugins.CallbackName(s, true, 'D')).
					Register(plugins.CallbackName(cfg.PluginName, true, 'D'), cfg.Clause),

				tx.Callback().Update().
					Before(plugins.CallbackName(s, true, 'U')).
					Register(plugins.CallbackName(cfg.PluginName, true, 'U'), cfg.Clause),

				tx.Callback().Query().
					Before(plugins.CallbackName(s, true, 'Q')).
					Register(plugins.CallbackName(cfg.PluginName, true, 'Q'), cfg.Clause),
			)
			if registerBeforeErr != nil {
				return false
			}
		}
		return true
	})
	if registerBeforeErr != nil || hasBefore {
		return registerBeforeErr
	}

	return plugins.OneError(
		tx.Callback().Delete().Before("gorm:delete").
			Register(plugins.CallbackName(cfg.PluginName, true, 'D'), cfg.Clause),

		tx.Callback().Update().Before("gorm:update").
			Register(plugins.CallbackName(cfg.PluginName, true, 'U'), cfg.Clause),

		tx.Callback().Query().Before("gorm:query").
			Register(plugins.CallbackName(cfg.PluginName, true, 'Q'), cfg.Clause),
	)
}

const SkipRawOrRow = "skip_raw_row"
