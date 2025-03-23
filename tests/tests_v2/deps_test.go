package gormx_testv2

import (
	"database/sql"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/gormx/schemas/types"
	"testing"
	"time"
)

func TestItemValueIsZero(t *testing.T) {
	var nilMap map[string]string
	var nilSlice []int
	for _, vi := range []any{
		nil,         // nil
		true, false, // bool
		0, -666, // int
		uint(0), uint(888), // uint
		"", "no-zero", // string
		0.0, 0.01, // float
		nilSlice, []int{}, []int{1, 2, 3}, // slice
		[0]int{}, [3]int{}, [3]int{1, 2, 3}, // array
		nilMap, map[string]string{}, map[string]string{"1": "1"}, // map
		struct{}{}, // empty struct
		struct {
			Int    int
			String string
			Float  float64
		}{}, // some struct
		time.Time{}, time.Now(), // time
		sql.NullString{}, sql.NullString{Valid: true}, // sql.NullString
		sql.NullInt64{}, sql.NullInt64{Valid: true}, // sql.NullInt64
		sql.NullTime{}, sql.NullTime{Valid: true}, // sql.NullTime
		sql.NullFloat64{}, sql.NullFloat64{Valid: true}, // sql.NullFloat64
		sql.NullBool{}, sql.NullBool{Valid: true}, // sql.NullBool
		types.RMBCent(0), types.RMBCent(1), // types.*
		types.Time{}, types.Time{Valid: true},
		types.DateTime{}, types.DateTime{Valid: true},
		types.Timestamp{}, types.Timestamp{Valid: true},
		types.BinarySize(0), types.BinarySize(1),
	} {
		t.Logf("Value: %10v, IsZero: %5t", vi, deps.ItemValueIsZero(vi))
	}
}
