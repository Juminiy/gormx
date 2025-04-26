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

// UnmarshalJSON
// - "2006-01-02T15:04:05.000Z" | time.Time
// - 1136214246 		   		| int64
func (t *Time) UnmarshalJSON(b []byte) error {
	if bStr := util.Bytes2StringNoCopy(b); InValidJSON(b) {
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
	}

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

// UnmarshalJSON
// - "2006-01-02T15:04:05.000Z" | time.Time
// - 1136214246 		   		| int64
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

// UnmarshalJSON
// - "2006-01-02T15:04:05.000Z" | time.Time
// - 1136214246 		   		| int64
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	if InValidJSON(b) {
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

type AnyTime sql.NullInt64

func (t *AnyTime) Scan(src any) error {
	return (*Timestamp)(t).Scan(src)
}

func (t AnyTime) Value() (driver.Value, error) {
	return (Timestamp)(t).Value()
}

func (t AnyTime) MarshalJSON() ([]byte, error) {
	return (Timestamp)(t).MarshalJSON()
}

// UnmarshalJSON
// receive format:
//
//	JSONValueFormat								 	  | GoType
//
// - "2006-01-02T15:04:05.000Z"  				 	  | time.Time
// - 1136214246									 	  | int64
// - "1136214246"								 	  | string(int64Value)
// - {"Time":"2006-01-02T15:04:05.000Z","Valid":true} | sql.NullTime
// - {"Int64":1136214246,"Valid":true}			 	  | sql.NullInt64
func (t *AnyTime) UnmarshalJSON(b []byte) error {
	if InValidJSON(b) {
		t.Valid = false
		return nil
	}

	if (*Timestamp)(t).UnmarshalJSON(b) == nil {
		return nil
	}

	var i64Str string
	if json.Unmarshal(b, &i64Str) == nil {
		if i64v, err := parseI64(i64Str); err == nil && i64v > 0 {
			t.Int64 = i64v
			t.Valid = true
			return nil
		}
	}

	var nullTimeVal sql.NullTime
	if json.Unmarshal(b, &nullTimeVal) == nil && !nullTimeVal.Time.IsZero() {
		t.Int64 = nullTimeVal.Time.Unix()
		t.Valid = nullTimeVal.Valid
		return nil
	}

	var nullI64Val sql.NullInt64
	if json.Unmarshal(b, &nullI64Val) == nil && nullI64Val.Int64 > 0 {
		*t = AnyTime(nullI64Val)
		return nil
	}

	return ErrTimeFromJSON
}
