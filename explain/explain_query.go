package explain

import (
	"encoding/json"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func QueryOrRow(tx *gorm.DB) {
	if tx.Statement.SQL.Len() > 0 {
		var (
			dialectorName = schemas.DialectorName(tx.Dialector)
			explainClause string
			explainResult any
		)
		switch dialectorName {
		case mysql.DefaultDriverName:
			explainClause, explainResult = "EXPLAIN", &[]MySQLPlan{}
		case "postgres":
			explainClause, explainResult = "EXPLAIN (FORMAT JSON)", util.New("")
		case sqlite.DriverName, "sqlite":
			explainClause, explainResult = "EXPLAIN", &[]SqliteVMInstruction{}
		default: // no support
			return
		}
		err := tx.Session(&gorm.Session{NewDB: true, SkipHooks: true}).
			Raw(explainClause+" "+tx.Statement.SQL.String(), tx.Statement.Vars...).
			Scan(explainResult).Error
		if err != nil {
			tx.Logger.Error(tx.Statement.Context, "explain sql query error: %s", err.Error())
			return
		}
		bs, err := json.Marshal(explainResult)
		if err != nil {
			tx.Logger.Error(tx.Statement.Context, "explain sql marshalJSON error: %s", err.Error())
			return
		}
		tx.Logger.Info(tx.Statement.Context, "explain sql plan: %s", string(bs))
	}
}
