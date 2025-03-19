package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
)

func (cfg *Config) BeforeUpdate(tx *gorm.DB) {
	if tx.Error != nil || callback.SkipUpdate.OK(tx) {
		return
	}
	sCfg := cfg.OptionConfig(tx)

	if !sCfg.AllowTenantGlobalUpdate && !tx.AllowGlobalUpdate {
		if clauses.NoWhereClause(tx) {
			_ = tx.AddError(ErrNotAllowTenantGlobalUpdate)
			return
		}
	}

	if sCfg.UpdateMapOmitUnknownKey {
		callback.BeforeUpdateMapDeleteUnknownColumn(tx)
	}

	if sCfg.UpdateMapOmitZeroElemKey {
		callback.BeforeUpdateMapDeleteZeroValueColumn(tx)
	}

	if sCfg.UpdateOptimisticLock {
		cfg.OptLockCfg().OptimisticLock(tx)
	}

	if sCfg.UpdateMapSetPkToClause {
		callback.BeforeUpdateMapDeletePkAndSetPkToClause(tx)
	}

	if !sCfg.DisableFieldDup {
		cfg.UniquesCfg().FieldDupCheck(tx, true, false)
		if tx.Error != nil {
			return
		}
	}

	if sCfg.UpdateMapCallHooks {
		callback.BeforeUpdateMapCallHook(tx)
	}

	cfg.AddTenantClauses(tx, false)
}

func (cfg *Config) AfterUpdate(tx *gorm.DB) {
	if tx.Error != nil || callback.SkipUpdate.OK(tx) {
		return
	}

	if cfg.OptionConfig(tx).AfterUpdateReturning &&
		util.MapOk(tx.Statement.Clauses, clauses.Returning) &&
		schemas.DialectorNotSupportReturningClause(tx.Dialector) {
		callback.AfterUpdateReturning(tx)
	}
}
