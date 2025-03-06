package schemas

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/jinzhu/now"
	"github.com/spf13/cast"
	"time"
)

type Time sql.NullTime

var ErrTimeNotValid = errors.New("value is not null or time.Time or sql.NullTime or timestamp")

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	}
	return json.Marshal(nil)
}

func (t *Time) UnmarshalJSON(b []byte) error {
	if bStr := util.Bytes2StringNoCopy(b); NotValidJSONValue(bStr) {
		t.Valid = false
		return nil
	} else if err := t.Time.UnmarshalJSON(b); err == nil {
		t.Valid = true
		return nil
	} else if bInt64 := cast.ToInt64(bStr); bInt64 > 0 {
		if unixTime := time.Unix(bInt64, 0); !unixTime.IsZero() {
			t.Valid = true
			t.Time = unixTime
			return nil
		}
	} else if parsedTime, err := now.Parse(bStr); err == nil {
		t.Valid = true
		t.Time = parsedTime
	} else {
		nullTime := sql.NullTime{}
		if err := json.Unmarshal(b, &nullTime); err == nil {
			*t = Time(nullTime)
		}
	}

	t.Valid = false
	return ErrTimeNotValid
}

func NotValidJSONValue(s string) bool {
	return util.ElemIn(s,
		`null`, ``, `0`, `0.0`,
		`"null"`, `""`, `"0"`, `"0.0"`,
	)
}

/*
func TrimStrEscape(s string) string {
	return strings.TrimRight(strings.TrimLeft(s, `"`), `"`)
}*/
