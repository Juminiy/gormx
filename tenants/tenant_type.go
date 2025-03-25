package tenants

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/gormx/schemas/types"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

/*
	func (t ID) QueryClauses(f *schema.Field) []clause.Interface {
		return t.tenantID().Clauses(f)
	}

	func (t ID) UpdateClauses(f *schema.Field) []clause.Interface {
		return t.tenantID().Clauses(f)
	}

	func (t ID) DeleteClauses(f *schema.Field) []clause.Interface {
		return t.tenantID().Clauses(f)
	}
*/
type ID sql.NullInt64

func (t ID) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Int64)
	}
	return json.Marshal(nil)
}

func (t *ID) UnmarshalJSON(b []byte) error {
	if types.InValidJSONValue(util.Bytes2StringNoCopy(b)) {
		t.Valid = false
		return nil
	} else if err := json.Unmarshal(b, &t.Int64); err == nil {
		t.Valid = true
		return nil
	} else {
		return err
	}
}

func (t *ID) Scan(value any) error {
	return (*sql.NullInt64)(t).Scan(value)
}

func (t ID) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Int64, nil
}

func (t ID) tenantID() tenantID {
	return tenantID{
		Valid: t.Valid,
		V:     t.Int64,
	}
}

/*
	func (t HideID) QueryClauses(f *schema.Field) []clause.Interface {
		return ID(t).tenantID().Clauses(f)
	}

	func (t HideID) UpdateClauses(f *schema.Field) []clause.Interface {
		return ID(t).tenantID().Clauses(f)
	}

	func (t HideID) DeleteClauses(f *schema.Field) []clause.Interface {
		return ID(t).tenantID().Clauses(f)
	}
*/
type HideID ID

func (t HideID) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (t *HideID) UnmarshalJSON(b []byte) error {
	return nil
}

func (t *HideID) Scan(value any) error {
	return (*sql.NullInt64)(t).Scan(value)
}

func (t HideID) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Int64, nil
}

/*
	func (t SID) QueryClauses(f *schema.Field) []clause.Interface {
		return t.tenantID().Clauses(f)
	}

	func (t SID) UpdateClauses(f *schema.Field) []clause.Interface {
		return t.tenantID().Clauses(f)
	}

	func (t SID) DeleteClauses(f *schema.Field) []clause.Interface {
		return t.tenantID().Clauses(f)
	}
*/
type SID sql.NullString

func (t SID) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.String)
	}
	return json.Marshal(nil)
}

func (t *SID) UnmarshalJSON(b []byte) error {
	if types.InValidJSONValue(util.Bytes2StringNoCopy(b)) {
		t.Valid = false
		return nil
	} else if err := json.Unmarshal(b, &t.String); err == nil {
		t.Valid = true
		return nil
	} else {
		return err
	}
}

func (t *SID) Scan(value any) error {
	return (*sql.NullString)(t).Scan(value)
}

func (t SID) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.String, nil
}

func (t SID) tenantID() tenantID {
	return tenantID{
		Valid: t.Valid,
		V:     t.String,
	}
}

type tenantID sql.Null[any]

func (t tenantID) Clauses(f *schema.Field) []clause.Interface {
	if t.Valid {
		fieldValue := schemas.FieldFromSchema(f)
		fieldValue.Value = t.V
		return []clause.Interface{&Tenant{Field: fieldValue}}
	}
	return nil
}
