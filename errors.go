package gormx

import "errors"

var ErrNotAllowTenantGlobalUpdate = errors.New("update tenant all rows or global update is not allowed")
var ErrNotAllowTenantGlobalDelete = errors.New("delete tenant all rows or global update is not allowed")
