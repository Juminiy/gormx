package tenants

import (
	"database/sql"
	"encoding/json"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type ID tenantID

func (t ID) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.V)
	}
	return json.Marshal(nil)
}

func (t *ID) UnmarshalJSON(b []byte) error {
	if schemas.NotValidJSONValue(util.Bytes2StringNoCopy(b)) {
		t.Valid = false
		return nil
	} else if err := json.Unmarshal(b, &t.V); err == nil {
		t.Valid = true
		return nil
	} else {
		return err
	}
}

func (t ID) QueryClauses(f *schema.Field) []clause.Interface {
	return tenantID(t).Clauses(f)
}

func (t ID) UpdateClauses(f *schema.Field) []clause.Interface {
	return tenantID(t).Clauses(f)
}

func (t ID) DeleteClauses(f *schema.Field) []clause.Interface {
	return tenantID(t).Clauses(f)
}

type HideID tenantID

func (t HideID) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (t *HideID) UnmarshalJSON(b []byte) error {
	return nil
}

func (t HideID) QueryClauses(f *schema.Field) []clause.Interface {
	return tenantID(t).Clauses(f)
}

func (t HideID) UpdateClauses(f *schema.Field) []clause.Interface {
	return tenantID(t).Clauses(f)
}

func (t HideID) DeleteClauses(f *schema.Field) []clause.Interface {
	return tenantID(t).Clauses(f)
}

type tenantID sql.Null[uint]

func (t tenantID) Clauses(f *schema.Field) []clause.Interface {
	if t.Valid {
		fieldValue := schemas.FieldFromSchema(f)
		fieldValue.Value = t.V
		return []clause.Interface{&Tenant{Field: fieldValue}}
	}
	return nil
}
