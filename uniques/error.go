package uniques

import (
	"fmt"
	"github.com/Juminiy/gormx/deps"
)

type FieldDupError interface {
	error
	DBTable() string
	DupDBName() []string
	ScopeKeys() []string
}

func IsFieldDupError(err error) bool {
	return deps.IndI(err).Type == _fieldDupCountErrRType ||
		deps.IndI(err).Type == _fieldDupListCheckErrRType
}

var _fieldDupCountErrRType = deps.IndI(fieldDupCountErr{}).Type
var _fieldDupListCheckErrRType = deps.IndI(fieldDupListCheckErr{}).Type

type fieldDupCountErr struct {
	dbTable   string
	dbName    []string
	scopeKeys []string
}

func (e fieldDupCountErr) Error() string {
	fieldDupDesc := fmt.Sprintf("field dup count error, table:[%s] column:(%v)",
		e.dbTable, e.dbName)
	if len(e.scopeKeys) > 0 {
		return fmt.Sprintf("%s, in scope:(%v)", fieldDupDesc, e.scopeKeys)
	}
	return fieldDupDesc
}

func (e fieldDupCountErr) DBTable() string {
	return e.dbTable
}

func (e fieldDupCountErr) DupDBName() []string {
	return e.dbName
}

func (e fieldDupCountErr) ScopeKeys() []string { return e.scopeKeys }

type fieldDupListCheckErr struct {
	fieldDupCountErr
	dupValues []any
}

func (e fieldDupListCheckErr) Error() string {
	return fmt.Sprintf("field dup list check error, table:[%s] column:(%v), values:(%v)",
		e.dbTable, e.dbName, e.dupValues)
}
