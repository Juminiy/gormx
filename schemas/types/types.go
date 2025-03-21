package types

import "github.com/Juminiy/kube/pkg/util"

func InValidJSONValue(s string) bool {
	return util.ElemIn(s,
		`null`, ``, `0`, `0.0`,
		`"null"`, `""`, `"0"`, `"0.0"`,
	)
}

/*
func TrimStrEscape(s string) string {
	return strings.TrimRight(strings.TrimLeft(s, `"`), `"`)
}*/
