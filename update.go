package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/clauses"
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

	if sCfg.UpdateMapSetPkToClause {
		callback.BeforeUpdateMapDeletePkAndSetPkToClause(tx)
	}

	if sCfg.UpdateMapOmitZeroElemKey {
		callback.BeforeUpdateMapDeleteZeroValueColumn(tx)
	}

	if sCfg.UpdateMapOmitUnknownKey {
		callback.BeforeUpdateMapDeleteUnknownColumn(tx)
	}

	if !sCfg.DisableFieldDup {
		cfg.UniquesCfg().FieldDupCheck(tx, true, sCfg.EnableComplexFieldDup)
		if tx.Error != nil {
			return
		}
	}

	if sCfg.UpdateMapCallHooks {
		callback.BeforeUpdateMapCallHook(tx)
	}

	cfg.AddTenantClause(tx, false)
}
