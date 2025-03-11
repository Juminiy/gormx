package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
)

func (cfg *Config) BeforeDelete(tx *gorm.DB) {
	if tx.Error != nil || callback.SkipDelete.OK(tx) {
		return
	}
	sCfg := cfg.OptionConfig(tx)

	if !sCfg.AllowTenantGlobalDelete && !tx.AllowGlobalUpdate {
		if clauses.NoWhereClause(tx) {
			_ = tx.AddError(ErrNotAllowTenantGlobalDelete)
			return
		}
	}

	cfg.AddTenantClauses(tx, false)

	if sCfg.BeforeDeleteReturning &&
		util.MapOk(tx.Statement.Clauses, clauses.Returning) &&
		schemas.DialectorNotSupportReturningClause(tx.Dialector) {
		callback.BeforeDeleteReturning(tx)
	}
}
