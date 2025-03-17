package gormx_testv2

import (
	"errors"
	"github.com/Juminiy/gormx"
	"testing"
)

var CComplexCheck = gormx.Option{EnableComplexFieldDup: true}
var ErrUniquesNotPreCheckList = errors.New("uniques plugin create not pre check list")

func TestCreateListPreCheck(t *testing.T) {
	t.Run("NoPlugin", func(tt *testing.T) {
		order := RandomOrder()
		orders := []*Order{order, order, order}
		Err(tt, iSqlite0().Create(&orders))
	})
	t.Run("IsPlugin, NoCheck, NoScopes", func(tt *testing.T) {
		order := RandomOrder()
		orders := []*Order{order, order, order}
		Err(tt, iSqlite().Create(&orders))
	})
	t.Run("IsPlugin, IsCheck, NoScopes", func(tt *testing.T) {
		order := RandomOrder()
		orders := []*Order{order, order, order}
		if err := iSqlite().Set(gormx.OptionKey, CComplexCheck).Create(&orders).Error; err == nil {
			tt.Error(ErrUniquesNotPreCheckList)
		} else {
			tt.Log(err.Error())
		}
	})
	t.Run("IsPlugin, IsCheck, Scopes(user_id)", func(tt *testing.T) {
		order := RandomOrder()
		orders := []*Order{order, order, order}
		if err := iSUser().Set(gormx.OptionKey, CComplexCheck).Create(&orders).Error; err == nil {
			tt.Error(ErrUniquesNotPreCheckList)
		} else {
			tt.Log(err.Error())
		}
	})
	t.Run("IsPlugin, IsCheck, Scopes(user_id, tenant_id)", func(tt *testing.T) {
		order := RandomOrder()
		orders := []*Order{order, order, order}
		if err := iSUserTenant().Set(gormx.OptionKey, CComplexCheck).Create(&orders).Error; err == nil {
			tt.Error(ErrUniquesNotPreCheckList)
		} else {
			tt.Log(err.Error())
		}
	})
}
