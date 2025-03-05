package tenants

import (
	"database/sql"
	"encoding/json"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type TenantID sql.Null[uint]

func (t TenantID) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.V)
	}
	return json.Marshal(nil)
}

func (t *TenantID) UnmarshalJSON(b []byte) error {
	if util.Bytes2StringNoCopy(b) == "null" {
		t.Valid = false
		return nil
	} else if err := json.Unmarshal(b, &t); err == nil {
		t.Valid = true
		return nil
	} else {
		return err
	}
}

func (t TenantID) QueryClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{&Tenant{Field: schemas.Field{
		Name:    f.Name,
		DBTable: f.Schema.Table,
		DBName:  f.DBName,
		Value:   t.V,
		Values:  nil,
	}}}
}

func (t TenantID) UpdateClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{&Tenant{Field: schemas.Field{
		Name:    f.Name,
		DBTable: f.Schema.Table,
		DBName:  f.DBName,
		Value:   t.V,
		Values:  nil,
	}}}
}

func (t TenantID) DeleteClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{&Tenant{Field: schemas.Field{
		Name:    f.Name,
		DBTable: f.Schema.Table,
		DBName:  f.DBName,
		Value:   t.V,
		Values:  nil,
	}}}
}
