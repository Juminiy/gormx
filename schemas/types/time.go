package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"github.com/Juminiy/kube/pkg/util"
	"time"
)

type Time sql.NullTime

var ErrTimeFromDB = ValueFromDBError("Time")

var ErrTimeFromJSON = ValueFromJSONError("Time")

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
	} else if bI64, err := parseI64(bStr); err == nil && bI64 > 0 {
		if unixTime := time.Unix(bI64, 0); !unixTime.IsZero() {
			t.Time = unixTime
			t.Valid = true
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

type Timestamp sql.NullInt64

func (t *Timestamp) Scan(src any) error {
	if unixSec, ok := src.(int64); ok {
		t.Int64 = unixSec
		t.Valid = true
		return nil
	}
	return ErrTimeFromDB
}

func (t Timestamp) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Int64, nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return Time{Time: time.Unix(t.Int64, 0), Valid: t.Valid}.MarshalJSON()
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	if bStr := util.Bytes2StringNoCopy(b); InValidJSONValue(bStr) {
		t.Valid = false
		return nil
	}

	if err := json.Unmarshal(b, &t.Int64); err == nil {
		t.Valid = true
		return nil
	}
	timeVal := time.Time{}
	if err := timeVal.UnmarshalJSON(b); err == nil {
		t.Int64 = timeVal.Unix()
		t.Valid = true
		return nil
	}

	t.Valid = false
	return ErrTimeFromJSON
}
