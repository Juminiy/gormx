package callback

import (
	"gorm.io/gorm"
)

func BeforeDeleteReturning(tx *gorm.DB) {
	returningQuery(tx, tx.Statement.Dest)
}
