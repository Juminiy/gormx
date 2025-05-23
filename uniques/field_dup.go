package uniques

import (
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	expmaps "golang.org/x/exp/maps"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"reflect"
	"slices"
	"strings"
	"sync"
)

type Config struct {
	Name         string   // default: gormx:uniques
	TagKey       string   // default: gormx
	TagUniqueKey string   // default: unique
	TxKeys       []string // txKey: elem of gormx.Config .KnownScopes

	cacheStore *sync.Map
}

func (cfg *Config) Initial() *Config {
	cfg.cacheStore = new(sync.Map)
	return cfg
}

type FieldDup struct {
	Clauses     []clause.Interface
	DBTable     string
	FieldColumn map[string]string
	ColumnField map[string]string
	Groups      map[string][]string // Groups[group_name] -> FieldList

	cfg *Config
}

// FieldDupInfo
// Create(&Struct), Create(&[]Struct), Create(&[N]Struct)
// Create(&map[string]any{}), Create(&[]map[string]any{})
// Create map[string]any ~ Map[K]V, K(string) is FieldName, V(any) is FieldValue
// Updates(&Struct)
// Updates(&map[string]any{})
// Updates map[string]any ~ Map[K]V, K(string) is ColumnName, V(any) is FieldValue
func (cfg *Config) FieldDupInfo(tx *gorm.DB) *FieldDup {
	if sch := tx.Statement.Schema; sch != nil {
		schFieldDupKey := util.StringJoin(":", cfg.Name, cfg.TagUniqueKey, sch.Table)
		if fieldDup, ok := cfg.cacheStore.Load(schFieldDupKey); ok {
			return fieldDup.(*FieldDup)
		} else if fieldDup := cfg.fieldDupInfo(tx); fieldDup != nil {
			fieldDup.cfg = cfg
			cfg.cacheStore.Store(schFieldDupKey, fieldDup)
			return fieldDup
		}
	}
	return nil
}

func (cfg *Config) fieldDupInfo(tx *gorm.DB) *FieldDup {
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
			ForUpdate:   true,
		}).simple(tx)

	case *map[string]any:
		(&rowValues{
			ColumnValue: *columnValue,
			FieldDup:    d,
			ForUpdate:   true,
		}).simple(tx)

	default:
		rval := deps.IndI(dest)
		switch rval.Type.Kind() {
		case reflect.Struct:
			(&rowValues{
				FieldValue: rval.StructValues(),
				FieldDup:   d,
				ForUpdate:  true,
			}).simple(tx)

		case reflect.Map:
			(&rowValues{
				ColumnValue: rval.MapValues(),
				FieldDup:    d,
				ForUpdate:   true,
			}).simple(tx)

		default: // ignore case
		}
	}
}

func (d *FieldDup) doCount(tx *gorm.DB, exprI clause.Expression, forUpdate bool) {
	scopeKeys := sets.New[string]()
	// new session and copy settings
	ntx := tx.Session(&gorm.Session{NewDB: true, SkipHooks: true})
	// where clause
	// 1. copy scopes session values: tenant_id, user_id, project_id, ...
	// 2. do not skip Query Count(*) callbacks /*callback.SkipQuery.Set(ntx).*/
	tx.Statement.Settings.Range(func(key, value any) bool {
		if keyStr, ok := key.(string); ok {
			if scopeKey, ok := prefixOrSuffixIn(keyStr, d.cfg.TxKeys...); ok {
				scopeKeys.Insert(scopeKey)
				ntx = ntx.Set(keyStr, value)
			}
		}
		return true
	})
	/*ntx.Statement.Schema = tx.Statement.Schema*/

	// where clause 2. orExpr
	ntx = ntx.Table(d.DBTable).Where(exprI)

	// where clause 3. soft_delete
	// and other clauses, maybe unused
	slices.All(d.Clauses)(func(_ int, c clause.Interface) bool {
		ntx.Statement.AddClause(c)
		return true
	})

	// where clause for update 4. NOT(tx.Clause)
	if forUpdate {
		var exprs []clause.Expression
		// implicit where clause (primaryKey)
		if expr, ok := clauses.StmtPrimaryKeyClause(tx.Statement); ok {
			exprs = append(exprs, expr)
		}
		// explicit where clause
		if txClause, ok := clauses.WhereClause(tx.Statement); ok {
			exprs = append(exprs, txClause)
		}
		if len(exprs) > 0 {
			ntx.Statement.AddClause(clause.Where{
				Exprs: []clause.Expression{clause.Not(exprs...)},
			})
		}
	}

	// do Count
	var cnt int64
	err := ntx.Count(&cnt).Error
	if err != nil {
		ntx.Logger.Error(ntx.Statement.Context, "before %s, do field uniques check, error: %s", func() string {
			if forUpdate {
				return "update"
			} else {
				return "create"
			}
		}(), err.Error())
		return
	}
	if cnt > 0 {
		fdErr := fieldDupCountErr{
			dbTable:   d.DBTable,
			dbName:    expmaps.Keys(d.ColumnField),
			scopeKeys: scopeKeys.UnsortedList(),
		}
		_ = tx.AddError(fdErr)
	}
}

func prefixOrSuffixIn(s string, keys ...string) (scopeKey string, prefixOrSuffix bool) {
	slices.All(keys)(func(_ int, key string) bool {
		if strings.HasPrefix(s, key) || strings.HasSuffix(s, key) {
			scopeKey, prefixOrSuffix = key, true
			return false
		}
		return true
	})
	return
}
