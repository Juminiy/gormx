package gormx_tests

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"testing"
)

// fixed
func TestFixBugAfterCreate(t *testing.T) {
	var cwMap = map[string]any{
		"Name": "sandbox-99",
	}
	Err(t, txFull().Table(`tbl_calico_weave`).Create(&cwMap).Error)
	t.Log(Enc(cwMap))
	// bug1: TenantID is hidden
	// bug2: AfterCreate Hooks not effected
}

// fixed
func TestFixBugBeforeUpdateTenantOmit(t *testing.T) {
	Err(t, txFull().Updates(&CalicoWeave{
		Model:    gorm.Model{ID: 2},
		TenantID: util.MaxUint,
		UserID:   util.MaxUint,
		Name:     "MyName-2",
	}).Error)

	Err(t, txFull().Table(`tbl_calico_weave`).Updates(map[string]any{
		"id":        "2",
		"tenant_id": util.MaxUint,
		"user_id":   util.MaxUint,
		"name":      "MyName-2",
	}).Error)
}

// fixed
func TestFixBugBeforeUpdateSetModelPkToClause(t *testing.T) {
	Err(t, txFull().Updates(&CalicoWeave{
		Model:    gorm.Model{ID: 2},
		TenantID: util.MaxUint,
		UserID:   util.MaxUint,
		Name:     "MyName-2",
	}).Error)
}

// zero data, ignore it
func TestNoScopeUniquesErrorOrExpr(t *testing.T) {
	Err(t, txPure().Create((&BabyTrade{}).RandomSet()).Error)
}
