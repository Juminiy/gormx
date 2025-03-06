package uniques

import (
	"fmt"
	"github.com/Juminiy/gormx/deps"
)

type FieldDupError interface {
	error
	DBTable() string
	TenantDBName() string
	TenantValue() any
	DupDBName() []string
}

func IsFieldDupError(err error) bool {
	return deps.IndI(err).Type == _fieldDupErrRType
}

var _fieldDupErrRType = deps.IndI(fieldDupErr{}).Type

type fieldDupErr struct {
	dbTable      string
	tenantDBName string
	tenantValue  any
	dbName       []string
}

func (e fieldDupErr) Error() string {
	fieldDupDesc := fmt.Sprintf("field dup error, table:[%s] column:%v",
		e.dbTable, e.dbName)
	if len(e.tenantDBName) > 0 && e.tenantValue != nil {
		return fmt.Sprintf("%s, in tenant:([%s]:[%v])",
			fieldDupDesc, e.tenantDBName, e.tenantValue)
	}
	return fieldDupDesc
}

func (e fieldDupErr) DBTable() string {
	return e.dbTable
}

func (e fieldDupErr) TenantDBName() string {
	return e.tenantDBName
}

func (e fieldDupErr) TenantValue() any {
	return e.tenantValue
}

func (e fieldDupErr) DupDBName() []string {
	return e.dbName
}
