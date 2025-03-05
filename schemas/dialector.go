package schemas

import (
	"github.com/Juminiy/kube/pkg/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DialectorNotSupportReturningClause(dialector gorm.Dialector) bool {
	return util.ElemIn(dialector.Name(), mysql.DefaultDriverName)
}
