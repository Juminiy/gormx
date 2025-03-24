package gormx

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/explain"
	"gorm.io/gorm"
)

func (cfg *Config) BeforeRowOrRaw(tx *gorm.DB) {
	if tx.Error != nil || callback.SkipRowRaw.OK(tx) {
		return
	}

	if cfg.OptionConfig(tx).WriteClauseToRowOrRaw {
		callback.WriteToRowOrRaw(tx)
	}
}

func (cfg *Config) AfterRow(tx *gorm.DB) {
	if tx.Error != nil || callback.SkipRowRaw.OK(tx) {
		return
	}

	if cfg.OptionConfig(tx).ExplainQueryOrRow {
		explain.QueryOrRow(tx)
	}
}
