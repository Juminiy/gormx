package gormx_tests

import (
	"github.com/Juminiy/gormx/tenants"
	"gorm.io/gorm"
	"testing"
)

type Profile01 struct {
	gorm.Model
	Name     string
	TenantID uint   `mt:"tenant"`
	SimUUID  string `mt:"sim"`
	UserID   int    `mt:"user"`
}

type Profile0 struct {
	gorm.Model
	Name string
	tenants.ID
}

type Profile1 struct {
	gorm.Model
	Name     string
	TenantID tenants.ID
}

type Profile2 struct {
	gorm.Model
	Name     string
	TenantID tenants.ID
	UserID   tenants.ID
}

type Profile3 struct {
	gorm.Model
	Name      string
	TenantID  tenants.ID
	UserID    tenants.ID
	ProjectID tenants.ID
}

func TestNewTenantsType(t *testing.T) {

}
