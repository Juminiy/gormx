package gormx_tests

import (
	"github.com/Juminiy/gormx/callback"
	expmaps "golang.org/x/exp/maps"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/soft_delete"
	"testing"
)

type Product struct {
	gorm.Model
	Name       string `mt:"unique:name"`
	Desc       string `mt:"unique:name"`
	NetContent string
	Code       uint `mt:"unique:code"`
	Price      int64
	TenantID   uint `gorm:"index;" mt:"tenant" json:"-"`
}

func TestCreateProduct(t *testing.T) {
	//Err(t, txMigrate().AutoMigrate(&Product{}))
}

type InnerType struct {
	Name string
	Desc string
}

type WrapType1 struct {
	gorm.Model
	InnerType
	WrapString string
	//WrapStruct        InnerType // Error
	//WrapPointerStruct *InnerType // Error
}

type WrapType2 struct {
	DeletedAt gorm.DeletedAt
	*InnerType
	WrapString string
}

type InnerType2 struct {
	InnerType
	InnerString string
}

type WrapType3 struct {
	DeletedAt soft_delete.DeletedAt
	InnerType2
}

func showSchema(schema *schema.Schema) string {
	return Enc(map[string]any{
		"Name":             schema.Name,
		"Table":            schema.Table,
		"DBNames":          schema.DBNames,
		"FieldsByName":     expmaps.Keys(schema.FieldsByName),
		"FieldsByBindName": expmaps.Keys(schema.FieldsByBindName),
		"FieldsByDBName":   expmaps.Keys(schema.FieldsByDBName),
	})
}

func TestSchema(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(&WrapType1{}, &WrapType2{}, &WrapType3{}))
	for _, ttx := range []*gorm.DB{
		_txTenant().Find(&WrapType1{}),
		_txTenant().Find(&WrapType2{}),
		_txTenant().Find(&WrapType3{})} {
		t.Log(showSchema(ttx.Statement.Schema))
	}
}

func txMigrate(tx ...*gorm.DB) *gorm.DB {
	txn := iSqlite()
	if len(tx) > 0 {
		txn = tx[0]
	}
	return callback.SkipRowRaw.Set(txn)
}
