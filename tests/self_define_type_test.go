package gormx_tests

import (
	"database/sql"
	"encoding/json"
	"github.com/Juminiy/gormx/schemas/types"
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
	"testing"
	"time"
)

type Name string

func (n Name) MarshalJSON() ([]byte, error) {
	if n == "" || n == "null" {
		return nil, nil
	}
	return util.String2BytesNoCopy(string(n)), nil
}

func (n *Name) UnmarshalJSON(b []byte) error {
	if str := util.Bytes2StringNoCopy(b); str != "null" {
		*n = Name(str)
	}
	return nil
}

func TestMagicType(t *testing.T) {
}

func TestTimeAlias(t *testing.T) {
	vList := []string{
		Enc(struct {
			ID   uint
			Time string
		}{ID: 10,
			Time: "null"}), // [0]

		Enc(struct {
			ID   uint
			Time string
		}{ID: 10,
			Time: ""}), // [1]

		Enc(struct {
			ID   uint
			Time int
		}{ID: 10,
			Time: 0}), // [2]

		Enc(struct {
			ID   uint
			Time float64
		}{ID: 10,
			Time: 0.0, // [3]
		}),

		Enc(struct {
			ID   uint
			Time sql.NullTime
		}{ID: 10,
			Time: sql.NullTime{Time: time.Now()}}), // [4], not ok

		Enc(struct {
			ID   uint
			Time time.Time
		}{ID: 10,
			Time: time.Now()}), // [5]

		Enc(struct {
			ID   uint
			Time int64
		}{ID: 10,
			Time: time.Now().Unix(), // [6]
		}),

		Enc(struct {
			ID   uint
			Time string
		}{ID: 10,
			Time: time.Now().String(), // [7], not ok
		}),

		Enc(struct {
			ID   uint
			Time string
		}{ID: 10,
			Time: "2025-03-06 22:54:28", // [8], not ok
		}),
	}
	for i, timeRep := range vList {
		var v struct {
			ID   uint
			Time types.Time
		}
		err := json.Unmarshal([]byte(timeRep), &v)
		if err != nil {
			t.Logf("[%d], errJSON: [%s] error: [%s]", i, timeRep, err.Error())
		} else {
			t.Log(Enc(v))
		}
	}

}

func testTypesTimeType[TimeType any](t *testing.T) {
	t.Logf("test for: %s", reflect.TypeFor[TimeType]().String())
	var apiDataS = [][]byte{
		[]byte(`{
	"Fa": "2006-01-02T15:04:05.000Z"
}`), []byte(`{
	"Fb": 1136214246
}`), []byte(`{
	"Fc": "1136214246"
}`), []byte(`{
	"Fd": {"Time":"2006-01-02T15:04:05.000Z","Valid":true}
}`), []byte(`{
	"Fe": {"Int64":1136214246,"Valid":true}
}`),
	}

	/*var timeNameCalled = map[string]string{
		"Fa": "time.Time String",
		"Fb": "int64 timestamp",
		"Fc": "string int64Timestamp",
		"Fd": "sql.NullTime",
		"Fe": "sql.NullInt64",
	}*/

	for _, apiData := range apiDataS {
		var tTimeVal struct {
			Fa TimeType `json:",omitzero"`
			Fb TimeType `json:",omitzero"`
			Fc TimeType `json:",omitzero"`
			Fd TimeType `json:",omitzero"`
			Fe TimeType `json:",omitzero"`
		}
		if err := json.Unmarshal(apiData, &tTimeVal); err != nil {
			t.Logf("error: %s", err.Error())
		} else {
			bs, _ := json.Marshal(tTimeVal)
			t.Logf("%s", bs)
		}
	}
}

func TestTypesTimeType(t *testing.T) {
	t.Run("types.Time", func(tt *testing.T) {
		testTypesTimeType[types.Time](tt)
	})
	t.Run("types.Timestamp", func(tt *testing.T) {
		testTypesTimeType[types.Timestamp](tt)
	})
	t.Run("types.DateTime", func(tt *testing.T) {
		testTypesTimeType[types.DateTime](tt)
	})
	t.Run("types.AnyTime", func(tt *testing.T) {
		testTypesTimeType[types.AnyTime](tt)
	})
}

func testTypesTimeFormat[TimeType any](t *testing.T) {
	t.Logf("test for: %s", reflect.TypeFor[TimeType]().String())
	var apiDataS = [][]byte{
		[]byte(`{
	"Fa": "01/02 03:04:05PM '06 -0700"
}`), []byte(`{
	"Fa": "Mon Jan _2 15:04:05 2006"
}`), []byte(`{
	"Fa": "Mon Jan _2 15:04:05 MST 2006"
}`), []byte(`{
	"Fa": "Mon Jan 02 15:04:05 -0700 2006"
}`), []byte(`{
	"Fa": "02 Jan 06 15:04 MST"
}`), []byte(`{
	"Fa": "02 Jan 06 15:04 -0700"
}`), []byte(`{
	"Fa": "Monday, 02-Jan-06 15:04:05 MST"
}`), []byte(`{
	"Fa": "Mon, 02 Jan 2006 15:04:05 MST"
}`), []byte(`{
	"Fa": "Mon, 02 Jan 2006 15:04:05 -0700"
}`), []byte(`{
	"Fa": "2006-01-02T15:04:05Z07:00"
}`), []byte(`{
	"Fa": "2006-01-02T15:04:05.999999999Z07:00"
}`), []byte(`{
	"Fa": "Jan _2 15:04:05"
}`), []byte(`{
	"Fa": "Jan _2 15:04:05.000"
}`), []byte(`{
	"Fa": "Jan _2 15:04:05.000000"
}`), []byte(`{
	"Fa": "Jan _2 15:04:05.000000000"
}`), []byte(`{
	"Fa": "2006-01-02 15:04:05"
}`),
	}

	/*var timeNameCalled = map[string]string{
		"Fa": "time.Time String",
		"Fb": "int64 timestamp",
		"Fc": "string int64Timestamp",
		"Fd": "sql.NullTime",
		"Fe": "sql.NullInt64",
	}*/

	for _, apiData := range apiDataS {
		var tTimeVal struct {
			Fa TimeType `json:",omitzero"`
			Fb TimeType `json:",omitzero"`
			Fc TimeType `json:",omitzero"`
			Fd TimeType `json:",omitzero"`
			Fe TimeType `json:",omitzero"`
		}
		if err := json.Unmarshal(apiData, &tTimeVal); err != nil {
			t.Logf("error: %s", err.Error())
		} else {
			bs, _ := json.Marshal(tTimeVal)
			t.Logf("%s", bs)
		}
	}
}

func TestTypesTimeFormat(t *testing.T) {
	t.Run("types.Time", func(tt *testing.T) {
		testTypesTimeFormat[types.Time](tt)
	})
	t.Run("types.Timestamp", func(tt *testing.T) {
		testTypesTimeFormat[types.Timestamp](tt)
	})
	t.Run("types.DateTime", func(tt *testing.T) {
		testTypesTimeFormat[types.DateTime](tt)
	})
	t.Run("types.AnyTime", func(tt *testing.T) {
		testTypesTimeFormat[types.AnyTime](tt)
	})
}
