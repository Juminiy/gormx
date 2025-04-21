package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type RMBCent int64

var ErrParseRMBCent = errors.New("invalid string parse to rmb cent")

var ErrRMBCentValueFromDB = ValueFromDBError("RMBCent")

var ErrRMBCentValueFromJSON = ValueFromJSONError("RMBCent")

func (c *RMBCent) Scan(src any) error {
	if cent, ok := src.(int64); ok {
		*c = RMBCent(cent)
		return nil
	}
	return ErrRMBCentValueFromDB
}

func (c RMBCent) Value() (driver.Value, error) {
	return int64(c), nil
}

func (c RMBCent) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *RMBCent) UnmarshalJSON(b []byte) error {
	if InValidJSON(b) {
		return nil
	}

	var rawStr string
	if err := json.Unmarshal(b, &rawStr); err != nil {
		return err
	} else if err = c.Parse(rawStr); err == nil {
		return nil
	}

	return ErrRMBCentValueFromJSON
}

func (c RMBCent) String() string {
	centStr := strconv.FormatInt(int64(c), 10)
	centLen := len(centStr)
	switch centLen {
	case 1:
		return "0.0" + centStr
	case 2:
		return "0." + centStr
	default:
		return centStr[:centLen-2] + "." + centStr[centLen-2:]
	}
}

func (c *RMBCent) Parse(centStr string) error {
	var centY, centC string
	if len(centStr) > 20 {
		return ErrParseRMBCent
	} else if idxP := strings.IndexByte(centStr, '.'); idxP == 0 || idxP == len(centStr)-1 {
		return ErrParseRMBCent
	} else if idxP == -1 {
		centY = centStr
	} else { // if 0 < idxP && idxP < len(centStr)-1
		centY, centC = centStr[:idxP], centStr[idxP+1:]
	}

	notDigit := func(b rune) bool { return b < '0' || b > '9' }
	if len(centC) > 2 {
		return ErrParseRMBCent
	} else if strings.ContainsFunc(centY, notDigit) || strings.ContainsFunc(centC, notDigit) {
		return ErrParseRMBCent
	}

	var cent int64
	cent100, err := parseI64(centY)
	if err != nil {
		return ErrParseRMBCent
	}
	cent = cent100 * 100

	if len(centC) == 1 {
		cent += int64(centC[0]-'0') * 10
	} else if len(centC) == 2 {
		cent += int64(centC[0]-'0')*10 + int64(centC[1]-'0')
	}

	*c = RMBCent(cent)
	return nil
}
