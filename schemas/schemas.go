package schemas

import (
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"slices"
	"sync"
)

type Config struct {
	PluginName   string // default: gormx:schemas
	TagKey       string // default: gx
	TagUniqueKey string // default: unique

	cacheStore *sync.Map
}

func (cfg *Config) ParseSchema(tx *gorm.DB) {
	stmt := tx.Statement
	if len(stmt.Table) != 0 &&
		(stmt.Schema == nil || // no Schema
			(stmt.Schema != nil && (len(stmt.Schema.Name) == 0 || // has unnamed Schema
				func() bool {
					parsedSchema, ok := cfg.cacheStore.Load(cfg.graspSchemaKey(stmt.Table))
					return ok && parsedSchema.(*schema.Schema) != stmt.Schema
				}()))) { // destSchema != parsedSchema
		if zeroV, ok := cfg.cacheStore.Load(cfg.graspModelKey(stmt.Table)); ok {
			if err := stmt.Parse(zeroV); err != nil {
				tx.Logger.Error(stmt.Context, "use table parse schema error: %s", err.Error())
			}
		}
	}
}

func (cfg *Config) GraspSchema(tx *gorm.DB, zeroList ...any) {
	if cfg.cacheStore == nil {
		cfg.cacheStore = new(sync.Map)
	}
	slices.All(zeroList)(func(_ int, zeroV any) bool {
		stmt := tx.Statement
		err := stmt.Parse(zeroV)
		if err != nil {
			tx.Logger.Warn(stmt.Context, "user table grasp schema error: %s", err.Error())
		} else if stmt.Schema != nil {
			cfg.cacheStore.Store(cfg.graspSchemaKey(stmt.Schema.Table), stmt.Schema)
			cfg.cacheStore.Store(cfg.graspModelKey(stmt.Schema.Table), deps.IndI(zeroV).Interface())
		}
		return true
	})

}

func (cfg *Config) graspSchemaKey(tableName string) string {
	return util.StringJoin(":", cfg.PluginName, "grasp_schema", tableName)
}

func (cfg *Config) graspModelKey(tableName string) string {
	return util.StringJoin(":", cfg.PluginName, "grasp_model", tableName)
}
