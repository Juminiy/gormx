package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/spf13/cast"
	"time"
)

type Time sql.NullTime

var ErrTimeFromJSON = errors.New("value is not null, time or timestamp from json")

var ErrTimeFromDB = errors.New("value convert is invalid from database")

func (t *Time) Scan(src any) error {
	return (*sql.NullTime)(t).Scan(src)
}

func (t Time) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	}
	return json.Marshal(nil)
}

func (t *Time) UnmarshalJSON(b []byte) error {
	if bStr := util.Bytes2StringNoCopy(b); InValidJSONValue(bStr) {
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
	} /*else if parsedTime, err := now.Parse(bStr); err == nil {
		t.Valid = true
		t.Time = parsedTime
	} else {
		nullTime := sql.NullTime{}
		if err := json.Unmarshal(b, &nullTime); err == nil {
			*t = Time(nullTime)
		}
	}*/

	t.Valid = false
	return ErrTimeFromJSON
}

type DateTime Time

func (t *DateTime) Scan(src any) error {
	return (*sql.NullTime)(t).Scan(src)
}

func (t DateTime) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time.Format(time.DateTime), nil
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return (Time)(t).MarshalJSON()
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	return (*Time)(t).UnmarshalJSON(b)
}

type Timestamp Time

func (t *Timestamp) Scan(src any) error {
	if unixSec, ok := src.(int64); ok {
		t.Time = time.Unix(unixSec, 0)
		t.Valid = true
		return nil
	}
	return ErrTimeFromDB
}

func (t Timestamp) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time.Unix(), nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return (Time)(t).MarshalJSON()
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	return (*Time)(t).UnmarshalJSON(b)
}
