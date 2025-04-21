package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

var ErrExampleTypFromJSON = ValueFromJSONError("ExampleTyp")

var ErrExampleTypFromDB = ValueFromDBError("ExampleTyp")

type ExampleTyp struct {
	IntTyp       int
	StrTyp       string
	TimeTyp      time.Time
	NullTimeType sql.NullTime
	TimePtrTyp   *time.Time
	Balance      RMBCent
	BinSz        BinarySize
	RMap         map[string]string
}

type exampleTyp ExampleTyp

func (e *ExampleTyp) Scan(v any) error {
	if bs, ok := v.([]byte); ok {
		return json.Unmarshal(bs, e)
	}
	return ErrExampleTypFromDB
}

func (e ExampleTyp) Value() (driver.Value, error) {
	return json.Marshal(e)
}

// UnmarshalJSON
// can receive JSONBytes or EscapedJSONString
func (e *ExampleTyp) UnmarshalJSON(b []byte) error {
	if InValidJSON(b) {
		return nil
	}
	var eVal exampleTyp
	if err := json.Unmarshal(b, &eVal); err == nil {
		*e = ExampleTyp(eVal)
		return nil
	}
	var eStr string
	if err := json.Unmarshal(b, &eStr); err == nil {
		if err := json.Unmarshal([]byte(eStr), &eVal); err == nil {
			*e = ExampleTyp(eVal)
			return nil
		}
	}
	return ErrExampleTypFromJSON // parse typedValue error, return error
	//return nil                   // or parse typedValue error, ignore error
}

func (e ExampleTyp) MarshalJSON() ([]byte, error) {
	if e.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(exampleTyp(e))
}

func (e ExampleTyp) IsZero() bool {
	return e.IntTyp == 0 && len(e.StrTyp) == 0 &&
		e.TimeTyp.IsZero() && !e.NullTimeType.Valid &&
		e.TimePtrTyp == nil && e.Balance == 0 &&
		e.BinSz == 0 && len(e.RMap) == 0

}
