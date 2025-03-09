package gormx_tests

import (
	"github.com/Juminiy/gormx/plugins"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/gorm"
	"sync"
	"testing"
)

var once sync.Once

func skipTx() *gorm.DB {
	once.Do(func() {
		util.Must(plugins.OneError(
			txPure().Callback().Create().Remove("gormx:before_create"),
			txPure().Callback().Create().Remove("gormx:after_create"),
			txPure().Callback().Create().Remove("gormx:before_query"),
			txPure().Callback().Create().Remove("gormx:after_query"),
			txPure().Callback().Create().Remove("gormx:before_update"),
			//txPure().Callback().Create().Remove("gormx:after_update"),
			txPure().Callback().Create().Remove("gormx:before_delete"),
			//txPure().Callback().Create().Remove("gormx:after_delete"),
			txPure().Callback().Create().Remove("clauses:before_row"),
			txPure().Callback().Create().Remove("clauses:before_raw"),
			txPure().Callback().Create().Remove("clauses:before_query"),
			txPure().Callback().Create().Remove("clauses:before_update"),
			txPure().Callback().Create().Remove("clauses:before_delete"),
		))
	})
	return txPure()
}

func TestSkipPlugin(t *testing.T) {
	consumerMap := map[string]any{
		"AppID": 10,
	}
	Err(t, skipTx().
		Model(&Consumer{}).
		Create(&consumerMap).Error)
	t.Log(Enc(consumerMap))

	// panic
	/*consumerMap2 := map[string]any{
		"AppID": 20,
	}
	Err(t, skipTx().
		Table(`tbl_consumer`).
		Create(&consumerMap2).Error)
	t.Log(Enc(consumerMap2))*/
}
