package gormx_tests

import (
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"testing"
)

func TestColumnsInExpr(t *testing.T) {
	var bt BabyTrade
	var insExpr = clause.IN{
		Column: []string{"id", "deleted_at", "sim_uuid"},
		Values: []any{[]any{1, 0, uuid.NewString()}},
	}

	Err(t, iSqlite().Find(&bt, insExpr).Error)

	Err(t, iMySQL().Find(&bt, insExpr).Error)

	Err(t, iPg().Find(&bt, insExpr).Error)
}
