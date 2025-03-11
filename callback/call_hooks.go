package callback

import (
	"github.com/Juminiy/gormx/clauses"
	"github.com/Juminiy/gormx/deps"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"maps"
	"reflect"
	"slices"
)

// referred from: callbacks.callMethod
func CallHooks(db *gorm.DB, fc func(any, *gorm.DB) bool) {
	ntx := db.Session(&gorm.Session{NewDB: true})

	switch db.Statement.Dest.(type) {
	case map[string]any, *map[string]any:
		if structValue := deps.IndI(db.Statement.Model); structValue.CanAddr() { // *T -> T
			fc(structValue.Addr().Interface(), ntx)
		}

	case *[]map[string]any:
		structSlice := deps.IndI(db.Statement.Model).V // *[]*T -> []*T
		for i := 0; i < structSlice.Len(); i++ {
			if addrValue := reflect.Indirect(structSlice.Index(i)); addrValue.CanAddr() { // *T -> T
				fc(addrValue.Addr().Interface(), ntx)
			}
		}

	default: // ignore
	}

}

func setUpDestMapStmtModel(tx *gorm.DB, sch *schema.Schema) {
	//if tx.Statement.Model == tx.Statement.Dest
	switch deps.IndI(tx.Statement.Dest).T.Kind() {
	case reflect.Slice: // *[]map[string]any, []map[string]any
		// only for create
		tx.Statement.Model = sch.MakeSlice().Interface() // *[]*T

	case reflect.Map: // *map[string]any, map[string]any
		if modelInd := deps.IndI(tx.Statement.Model); modelInd.T.Kind() == reflect.Struct &&
			modelInd.CanAddr() { // Model is *T, **T, ...
			// do nothing
		} else { // for create
			tx.Statement.Model = reflect.New(sch.ModelType).Interface() // *T
			if modelInd.T.Kind() == reflect.Struct && !modelInd.IsZero() {
				deps.IndI(tx.Statement.Model).Set(modelInd.Value)
			}
		}

	default: // ignore
	}

	switch mapValue := tx.Statement.Dest.(type) {
	// Model *T
	case map[string]any:
		deps.IndI(tx.Statement.Model).StructSet(toFieldValue(sch, mapValue))

	case *map[string]any:
		deps.IndI(tx.Statement.Model).StructSet(toFieldValue(sch, *mapValue))

		// Model *[]*T
	case *[]map[string]any:
		structSlice := deps.IndI(tx.Statement.Model)
		slices.All(*mapValue)(func(_ int, m map[string]any) bool {
			newElem := reflect.New(sch.ModelType)             // *T
			deps.Ind(newElem).StructSet(toFieldValue(sch, m)) // *T <- m
			structSlice.SliceAppend(newElem.Interface())      // Model = append(Model, *T)
			return true
		})

	default: //ignore
	}
}

func scanModelToDestMap(tx *gorm.DB) {
	switch destValue := tx.Statement.Dest.(type) {
	case map[string]any:
		scanModelValueToDestValue(deps.IndI(tx.Statement.Model).StructToMap(), destValue)

	case *map[string]any:
		scanModelValueToDestValue(deps.IndI(tx.Statement.Model).StructToMap(), *destValue)

	case *[]map[string]any:
		slices.All(deps.IndI(tx.Statement.Model).SliceStructValues())(func(i int, m map[string]any) bool {
			scanModelValueToDestValue(m, (*destValue)[i])
			return true
		})

	default: // ignore
	}
}

func scanModelValueToDestValue(modelValue, destValue map[string]any) {
	maps.All(modelValue)(func(field string, modelFv any) bool {
		modelFieldValueIsZero := deps.ItemValueIsZero(modelFv)
		if destFv, ok := destValue[field]; ok && modelFieldValueIsZero {
			delete(destValue, field)
		} else if (!ok || deps.ItemValueIsZero(destFv)) &&
			deps.Comp(reflect.TypeOf(modelFv)) &&
			!modelFieldValueIsZero {
			destValue[field] = modelFv
		}
		return true
	})
}

func scanDestMapToModel(tx *gorm.DB) {
	// omit embedded fields
	switch destValue := tx.Statement.Dest.(type) {
	case map[string]any:
		deps.IndI(tx.Statement.Model).StructSet(destValue)

	case *map[string]any:
		deps.IndI(tx.Statement.Model).StructSet(*destValue)

	case *[]map[string]any:
		modelSlice := deps.IndI(tx.Statement.Model)
		slices.All(*destValue)(func(i int, m map[string]any) bool {
			deps.Ind(modelSlice.Index(i)).StructSet(m)
			return true
		})
	}
}

func toFieldValue(sch *schema.Schema, values map[string]any) map[string]any {
	return lo.MapKeys(values, func(_ any, columnOrField string) string {
		field := sch.LookUpField(columnOrField)
		if field != nil {
			return field.Name
		}
		return ""
	})
}

func toColumnValue(sch *schema.Schema, values map[string]any) map[string]any {
	return lo.MapKeys(values, func(_ any, columnOrField string) string {
		field := sch.LookUpField(columnOrField)
		if field != nil {
			return field.DBName
		}
		return ""
	})
}

func returningQuery(tx *gorm.DB, dest any) {
	if tx.Error != nil {
		return
	}
	ntx := tx.Session(&gorm.Session{NewDB: true, SkipHooks: true})

	ntx = ntx.Table(tx.Statement.Table)

	if sch := tx.Statement.Schema; sch != nil {
		slices.All(sch.QueryClauses)(func(_ int, c clause.Interface) bool {
			ntx.Statement.AddClause(c)
			return true
		})
	}

	if txClause, ok := clauses.WhereClause(tx); ok {
		ntx.Statement.AddClause(txClause)
	}

	if returning, ok := clauses.ReturningClause(tx); ok && len(returning.Columns) > 0 {
		slices.All(returning.Columns)(func(_ int, column clause.Column) bool {
			ntx.Statement.Selects = append(ntx.Statement.Selects, column.Name)
			return true
		})
	} else if len(tx.Statement.Selects) != 0 {
		ntx.Statement.Selects = tx.Statement.Selects
	}

	err := ntx.Find(dest).Error
	if err != nil {
		ntx.Logger.Error(ntx.Statement.Context, "before delete, do query, error: %s", err.Error())
	}
}
