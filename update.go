package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/gormx/schemas"
	"gorm.io/gorm"
)

func (cfg *Config) BeforeUpdate(tx *gorm.DB) {
	if tx.Error != nil {
		return
	}
	sCfg := SessionConfig(cfg, tx)

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
		(&schemas.Config{}).FieldDupCheck(tx, true)
		if tx.Error != nil {
			return
		}
	}

	if sCfg.UpdateMapCallHooks {
		callback.BeforeUpdateMapCallHook(tx)
	}

	cfg.AddTenantClause(tx, false)
}
