package uniques

import (
	"github.com/Juminiy/gormx/callback"
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"reflect"
	"slices"
	"strings"
	"sync"
)

type Config struct {
	Name         string // default: gormx:dup
	TagKey       string // default: gormx
	TagUniqueKey string // default: unique

	cacheStore *sync.Map
}

type FieldDup struct {
	Clauses     []clause.Interface
	DBTable     string
	FieldColumn map[string]string
	ColumnField map[string]string
	Groups      map[string][]string // Groups[key] -> FieldGroup
}

// FieldDupInfo
// Create(&Struct), Create(&[]Struct), Create(&[N]Struct)
// Create(&map[string]any{}), Create(&[]map[string]any{})
// Create map[string]any ~ Map[K]V, K(string) is FieldName, V(any) is FieldValue
// Updates(&Struct)
// Updates(&map[string]any{})
// Updates map[string]any ~ Map[K]V, K(string) is ColumnName, V(any) is FieldValue
func (cfg *Config) FieldDupInfo(tx *gorm.DB) *FieldDup {
	sch := tx.Statement.Schema
	if sch == nil {
		return nil
	}

	columnField := make(map[string]string, len(sch.DBNames)/4)
	groups := make(map[string][]string, len(sch.DBNames)/4)
	slices.All(sch.Fields)(func(_ int, field *schema.Field) bool {
		if mt, ok := field.Tag.Lookup(cfg.TagKey); ok {
			if keys, ok := util.MapElemOk(deps.Tag(mt), cfg.TagUniqueKey); ok {
				columnField[field.DBName] = field.Name
				slices.All(strings.Split(keys, ","))(func(_ int, key string) bool {
					if key == "-" { // ignore field
						return false
					} else if len(key) > 0 {
						groups[key] = append(groups[key], field.Name)
					} else {
						groups[field.Name] = []string{field.Name}
					}
					return true
				})

			}
		}
		return true
	})
	if len(groups) == 0 {
		return nil
	}

	return &FieldDup{
		Clauses:     sch.QueryClauses,
		DBTable:     sch.Table,
		FieldColumn: util.MapVK(columnField),
		ColumnField: columnField,
		Groups:      groups,
	}
}

func (cfg *Config) FieldDupCheck(tx *gorm.DB, forUpdate, arrOrSlice bool) {
	dupInfo := cfg.FieldDupInfo(tx)
	if dupInfo == nil {
		return
	}
	if forUpdate {
		dupInfo.Update(tx) // update map, struct
		return
	}
	if util.ElemIn(tx.Statement.ReflectValue.Kind(), reflect.Array, reflect.Slice) &&
		!arrOrSlice {
		return
	}
	dupInfo.Create(tx) // create
}

func (d *FieldDup) Create(tx *gorm.DB) {
	rval := deps.Ind(tx.Statement.ReflectValue)
	switch rval.Type.Kind() {
	case reflect.Struct:
		(&rowValues{
			FieldValue: rval.StructValues(),
			FieldDup:   d,
		}).simple(tx)

	case reflect.Map:
		(&rowValues{
			FieldValue: rval.MapValues(),
			FieldDup:   d,
		}).simple(tx)

	case reflect.Slice, reflect.Array:
		(&rowsValues{
			FieldDup: d,
		}).complex(tx)

	default: // ignore case
	}
}

func (d *FieldDup) Update(tx *gorm.DB) {
	dest := tx.Statement.Dest
	switch columnValue := dest.(type) {
	case map[string]any:
		(&rowValues{
			ColumnValue: columnValue,
			FieldDup:    d,
		}).simple(tx)

	case *map[string]any:
		(&rowValues{
			ColumnValue: *columnValue,
			FieldDup:    d,
		}).simple(tx)

	default:
		rval := deps.IndI(dest)
		switch rval.Type.Kind() {
		case reflect.Struct:
			(&rowValues{
				FieldValue: rval.StructValues(),
				FieldDup:   d,
			}).simple(tx)

		case reflect.Map:
			(&rowValues{
				ColumnValue: rval.MapValues(),
				FieldDup:    d,
			}).simple(tx)

		default: // ignore case
		}
	}
}

func (d *FieldDup) doCount(tx *gorm.DB, orExpr clause.Expression) {
	ntx := tx.Session(&gorm.Session{NewDB: true, SkipHooks: true})

	// where clause 1. orExpr
	ntx = callback.SkipQuery.Set(ntx).
		Table(d.DBTable).Where(orExpr)

	// where clause 2. tenant_id, user_id, ...
	// where clause 3. soft_delete
	// and other clauses
	slices.All(d.Clauses)(func(_ int, c clause.Interface) bool {
		ntx.Statement.AddClause(c)
		return true
	})

	// where clause for update 4. NOT(tx.Clause)
	if txClause, ok := clauses.WhereClause(tx); ok {
		ntx.Statement.AddClause(clause.Where{
			Exprs: []clause.Expression{clause.Not(txClause)},
		})
	}

	// do Count
	var cnt int64
	err := ntx.Count(&cnt).Error
	if err != nil {
		tx.Logger.Error(tx.Statement.Context, "before create or update, do field duplicated check, error: %s", err.Error())
		return
	}
	if cnt > 0 {
		fdErr := fieldDupErr{
			dbTable: d.DBTable,
			dbName:  util.MapKeys(d.ColumnField),
		}
		/*if d.Tenant != nil {
			fdErr.tenantDBName = d.Tenant.Field.DBName
			fdErr.tenantValue = d.Tenant.Field.Value
		}*/
		_ = tx.AddError(fdErr)
	}
}
