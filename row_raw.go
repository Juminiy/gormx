package gormx

import (
	"github.com/Juminiy/gormx/callback"
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
