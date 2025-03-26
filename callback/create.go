package callback

import (
	"github.com/Juminiy/gormx/deps"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/schema"
	"reflect"
	"slices"
	"time"
)

// referred from: callbacks.ConvertToCreateValues
func BeforeCreateSetDefaultValuesToMap(tx *gorm.DB) {
	// Field: gorm.Model.CreatedAt, gorm.Model.UpdatedAt, tag with default
	if sch, ok := hasSchemaAndDestIsMap(tx); ok {
		selectColumns, restricted := tx.Statement.SelectAndOmitColumns(true, false)
		setUp := setUpMapValues{
			sch:           sch,
			curTime:       tx.Statement.DB.NowFunc(),
			selectColumns: selectColumns,
			restricted:    restricted,
		}

		switch mapValue := tx.Statement.Dest.(type) {
		case map[string]any:
			setUp.Do(mapValue)

		case *map[string]any:
			setUp.Do(*mapValue)

		case *[]map[string]any:
			slices.All(*mapValue)(func(_ int, m map[string]any) bool {
				setUp.Do(m)
				return true
			})

		default: // ignore
		}
	}
}

type setUpMapValues struct {
	sch           *schema.Schema
	curTime       time.Time
	selectColumns map[string]bool
	restricted    bool
}

func (setUp *setUpMapValues) Do(mapValue map[string]any) {
	slices.All(setUp.sch.DBNames)(func(_ int, dbName string) bool {
		if field := setUp.sch.FieldsByDBName[dbName]; field.Creatable && // can create set
			!util.MapOk(mapValue, field.Name) && // AND mapValue no fieldName
			!util.MapOk(mapValue, dbName) && // AND mapValue no dbName
			((field.HasDefaultValue && field.DefaultValueInterface != nil) || // AND(`default`)
				field.NotNull || // AND (OR `not null`)
				field.AutoCreateTime > 0 || field.AutoUpdateTime > 0) { // AND (OR autoCreateTime OR autoUpdateTime)
			if v, ok := setUp.selectColumns[dbName]; (ok && v) ||
				(!ok && (!setUp.restricted || field.AutoCreateTime > 0 || field.AutoUpdateTime > 0)) {
				if field.DefaultValueInterface != nil {
					mapValue[field.Name] = field.DefaultValueInterface
				} else if field.AutoCreateTime > 0 || field.AutoUpdateTime > 0 {
					mapValue[field.Name] = setUp.curTime
				}
			}
		}
		return true
	})
}

func AfterCreateSetAutoIncPkToMap(tx *gorm.DB) {
	// write back MapType's autoIncrement primaryKey values
	if sch, ok := hasSchemaAndDestIsMap(tx); ok {
		autoIncPk := lo.Filter(sch.PrimaryFields, func(item *schema.Field, _ int) bool {
			return item.AutoIncrement
		})

		// this func can be many choices: addAutoIncPkNameByDBName
		autoIncPkFunc := replaceAutoIncPkDBNameToName

		// Create Map gorm can write back primaryKey values
		// but Map[key] is DBName(ColumnName) not Name(FieldName)
		// Map Type Support in gorm.Scan
		// 1. Create(map[string]any{})
		// 2. Create(&map[string]any{})
		// 3. Create(&[]map[string]any{})
		switch mapValue := tx.Statement.Dest.(type) {
		case map[string]any:
			autoIncPkFunc(autoIncPk, mapValue, mapValue)

		case *map[string]any:
			autoIncPkFunc(autoIncPk, *mapValue, *mapValue)

		case *[]map[string]any:
			mapSz := len(*mapValue) / 2
			dstMap, srcMap := (*mapValue)[:mapSz], (*mapValue)[mapSz:]
			slices.All(dstMap)(func(i int, m map[string]any) bool {
				autoIncPkFunc(autoIncPk, m, srcMap[i])
				return true
			})
			tx.Statement.ReflectValue.Set(reflect.ValueOf(dstMap))

		default: // ignore
		}
	}
}

// dest.(type) must be map[string]any, *map[string]any, *[]map[string]any
func hasSchemaAndDestIsMap(tx *gorm.DB) (sch *schema.Schema, ok bool) {
	sch = tx.Statement.Schema
	if sch == nil {
		return
	}
	switch destMap := tx.Statement.Dest.(type) {
	case map[string]any:
		if deps.IndI(tx.Statement.Dest).IsNil() {
			// var mapValue map[string]any
			return nil, false
		}
		// var mapValue = map[string]any{}
		// mapValue := map[string]any
		return sch, true
	case *map[string]any:
		if mapRv := deps.IndI(tx.Statement.Dest); mapRv.IsNil() {
			mapRv.SetI(map[string]any{})
		}
		return sch, true
	case *[]map[string]any:
		return sch, destMap != nil && len(*destMap) > 0
	case []map[string]any:
		return sch, len(destMap) > 0
	default:
		return nil, false
	}
}

// Replace Create Map Key:
// (Map[DBName] -> Value) To (Map[Name] -> Value)
func replaceAutoIncPkDBNameToName(autoIncPk []*schema.Field, dstMap, srcMap map[string]any) {
	slices.All(autoIncPk)(func(_ int, field *schema.Field) bool {
		if srcV, ok := srcMap[field.DBName]; ok { // DBName called ColumnName
			delete(dstMap, field.DBName)
			dstMap[field.Name] = srcV
		} else if srcV, ok = srcMap["@"+field.DBName]; ok { // @DBName called NamedColumnName
			delete(dstMap, "@"+field.DBName)
			dstMap[field.Name] = srcV
		}
		return true
	})
}

// Add Create Map Key:
// (Map[Name] -> Value) By (Map[DBName] -> Value)
func AddAutoIncPkNameByDBName(autoIncPk []*schema.Field, dstMap, srcMap map[string]any) {
	slices.All(autoIncPk)(func(_ int, field *schema.Field) bool {
		if srcV, ok := srcMap[field.DBName]; ok { // DBName called ColumnName
			//delete(dstMap, field.DBName)
			dstMap[field.Name] = srcV
		} else if srcV, ok = srcMap["@"+field.DBName]; ok { // @DBName called NamedColumnName
			//delete(dstMap, "@"+field.DBName)
			dstMap[field.Name] = srcV
		}
		return true
	})
}

// referred from: callbacks.BeforeCreate
func BeforeCreateMapCallHook(db *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(db); ok &&
		!db.Statement.SkipHooks && sch.BeforeCreate {
		setUpDestMapStmtModel(db, sch)
		CallHooks(db, func(v any, tx *gorm.DB) bool {
			if beforeCreateI, ok := v.(callbacks.BeforeCreateInterface); ok {
				_ = db.AddError(beforeCreateI.BeforeCreate(tx))
				return true
			}
			return false
		})
		scanModelToDestMap(db)
	}
}

// referred from: callbacks.AfterCreate
func AfterCreateMapCallHook(db *gorm.DB) {
	if sch, ok := hasSchemaAndDestIsMap(db); ok &&
		!db.Statement.SkipHooks && sch.AfterCreate {
		scanDestMapToModel(db)
		CallHooks(db, func(v any, tx *gorm.DB) bool {
			if afterCreateI, ok := v.(callbacks.AfterCreateInterface); ok {
				_ = db.AddError(afterCreateI.AfterCreate(tx))
				return true
			}
			return false
		})
		scanModelToDestMap(db)
	}
}
