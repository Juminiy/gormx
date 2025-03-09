package gormx_tests

import (
	"gorm.io/gorm/clause"
	"testing"
)

func TestSqliteReturning(t *testing.T) {
	var bt BabyTrade
	Err(t, iSqlite().Clauses(clause.Returning{}).Delete(&bt, 1).Error)
	t.Log(Enc(bt))
}

func TestMySQLNoReturning(t *testing.T) {
	var bt BabyTrade
	Err(t, iMySQL().Clauses(clause.Returning{}).Delete(&bt, 1).Error)
	t.Log(Enc(bt))
}
