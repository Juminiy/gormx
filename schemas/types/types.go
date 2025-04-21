package types

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"strconv"
)

func ValueFromDBError(typeDesc string) error {
	return fmt.Errorf("type: %s, value convert is illegal from database", typeDesc)
}

func ValueFromJSONError(typeDesc string) error {
	return fmt.Errorf("type: %s, value decode is illegal from json", typeDesc)
}

// InValidJSON
// - invalid-json will not throw error, ignore value only
// - illegal-json will throw error
func InValidJSON(b []byte) bool {
	return util.ElemIn(util.Bytes2StringNoCopy(b),
		`null`, `0`, `0.0`, `""`,
	)
}

func parseI64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func parseI32(s string) (int32, error) {
	i32, err := strconv.ParseInt(s, 10, 32)
	return int32(i32), err
}

/*
func TrimStrEscape(s string) string {
	return strings.TrimRight(strings.TrimLeft(s, `"`), `"`)
}*/
