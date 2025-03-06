package gorm_api

import (
	"database/sql"
	"encoding/json"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/kube/pkg/util"
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
	n := Name("null")
	bs, err := json.Marshal(n)
	Err(t, err)
	t.Logf("%s", bs)
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
			Time sql.NullTime
		}{ID: 10,
			Time: sql.NullTime{Time: time.Now()}}), // [3], not ok

		Enc(struct {
			ID   uint
			Time time.Time
		}{ID: 10,
			Time: time.Now()}), // [4]

		Enc(struct {
			ID   uint
			Time int64
		}{ID: 10,
			Time: time.Now().Unix(), // [5]
		}),

		Enc(struct {
			ID   uint
			Time string
		}{ID: 10,
			Time: time.Now().String(), // [6], not ok
		}),

		Enc(struct {
			ID   uint
			Time string
		}{ID: 10,
			Time: "2025-03-06 22:54:28", // [7], not ok
		}),
	}
	for i, timeRep := range vList {
		var v struct {
			ID   uint
			Time schemas.Time
		}
		err := json.Unmarshal([]byte(timeRep), &v)
		if err != nil {
			t.Logf("[%d], errJSON: [%s] error: [%s]", i, timeRep, err.Error())
		} else {
			t.Log(Enc(v))
		}
	}

}
