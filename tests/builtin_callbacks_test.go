package gormx_tests

import (
	"gorm.io/gorm"
	"testing"
)

func txPure() *gorm.DB {
	return iSqlite()
}

// gorm:begin_transaction -> callbacks.BeginTransaction
// gorm:before_create -> callbacks.BeforeCreate
// gorm:save_before_associations -> callbacks.SaveBeforeAssociations .func1
// gorm:create -> callbacks.Create .func1
// gorm:save_after_associations -> callbacks.SaveAfterAssociations .func1
// gorm:after_create -> callbacks.AfterCreate
// gorm:commit_or_rollback_transaction -> callbacks.CommitOrRollbackTransaction
func TestBuiltinCreateCallbacks(t *testing.T) {
	Err(t, txPure().Create(&AppUser{}).Error)
}

// gorm:begin_transaction -> callbacks.BeginTransaction
// gorm:setup_reflect_value -> callbacks.SetupUpdateReflectValue
// gorm:before_update -> callbacks.BeforeUpdate
// gorm:save_before_associations -> callbacks.SaveBeforeAssociations .func1
// gorm:update -> callbacks.Update .func1
// gorm:save_after_associations -> callbacks.SaveAfterAssociations .func1
// gorm:after_update -> callbacks.AfterUpdate
// gorm:commit_or_rollback_transaction -> callbacks.CommitOrRollbackTransaction
func TestBuiltinUpdateCallbacks(t *testing.T) {
	Err(t, txPure().Updates(&AppUser{}).Error)
}

// gorm:begin_transaction -> callbacks.BeginTransaction
// gorm:before_delete -> callbacks.BeforeDelete
// gorm:delete_before_associations -> callbacks.DeleteBeforeAssociations
// gorm:delete -> callbacks.Delete .func1
// gorm:after_delete -> callbacks.AfterDelete
// gorm:commit_or_rollback_transaction -> callbacks.CommitOrRollbackTransaction
func TestBuiltinDeleteCallbacks(t *testing.T) {
	Err(t, txPure().Delete(&AppUser{}).Error)
}

// gorm:query -> callbacks.Query
// gorm:preload -> callbacks.Preload
// gorm:after_query -> callbacks.AfterQuery
func TestBuiltinQueryCallbacks(t *testing.T) {
	Err(t, txPure().Find(&AppUser{}).Error)
}

// gorm:row -> callbacks.RowQuery
func TestBuiltinRowCallbacks(t *testing.T) {
	Err(t, txPure().Raw(`SELECT 1=1 AS TrueValue`).Scan(new(int)).Error)
}

// gorm:raw -> callbacks.RawExec
func TestBuiltinRawCallbacks(t *testing.T) {
	Err(t, txPure().Exec(`SELECT name FROM sqlite_master WHERE type='table'`).Error)

	var tables []string
	Err(t, txPure().Raw(`SELECT name FROM sqlite_master WHERE type='table'`).Scan(&tables).Error)
	t.Log(Enc(tables))
}
