package gorm_api

import (
	"database/sql"
	"github.com/Juminiy/gormx"
	"github.com/Juminiy/gormx/schemas"
	"github.com/Juminiy/gormx/tenants"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"testing"
	"time"
)

type Consumer struct {
	gorm.Model
	AppID    uint `gorm:"index"`
	TenantID uint `gorm:"index" mt:"tenant"`
	UserID   uint `gorm:"index"`
	VisitAt  gorm.DeletedAt
	LookupAt gorm.DeletedAt
	Region   string `gorm:"default:CN"`
}

func (c *Consumer) BeforeCreate(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		c.UserID = cast.ToUint(userID)
	}
	c.VisitAt = ValidTime(c.VisitAt, tx.NowFunc())
	c.LookupAt = ValidTime(c.LookupAt, tx.NowFunc().AddDate(1, 0, 0))
	return nil
}

func (c *Consumer) AfterCreate(tx *gorm.DB) error {
	c.UserID = 0
	return nil
}

func (c *Consumer) BeforeUpdate(tx *gorm.DB) error {
	tx.Logger.Info(tx.Statement.Context, "you are in PointerTo Consumer before update hooks")
	if userID, ok := tx.Get("user_id"); ok {
		tx.Statement.AddClause(ClauseUserID(tx.Statement, userID))
	}
	return nil
}

func (c *Consumer) AfterUpdate(tx *gorm.DB) error {
	tx.Logger.Info(tx.Statement.Context, "you are in PointerTo Consumer after update hooks")
	return nil
}

func (c *Consumer) BeforeDelete(tx *gorm.DB) error {
	if userID, ok := tx.Get("user_id"); ok {
		tx.Statement.AddClause(ClauseUserID(tx.Statement, userID))
	}
	return nil
}

func (c *Consumer) AfterDelete(tx *gorm.DB) error {
	return nil
}

func (c *Consumer) AfterFind(tx *gorm.DB) error {
	c.UserID = 0
	return nil
}

func ClauseUserID(stmt *gorm.Statement, userID any) clause.Interface {
	f := schemas.FieldFromSchema(stmt.Schema.FieldsByName["UserID"])
	f.Value = userID
	return &tenants.Tenant{Field: f}
}

func TestBuiltinHooks(t *testing.T) {
	Err(t, txMigrate().AutoMigrate(&Consumer{}))
}

func txHooks() *gorm.DB {
	return txMixed().Set(gormx.OptionKey, gormx.Option{
		BeforeCreateMapCallHooks: true,
		UpdateMapCallHooks:       true,
		AfterFindMapCallHooks:    true,
	})
}

// Deprecated: use clause.Interface instead
func DeletedAt(schema *schema.Schema) *schemas.Field { // maybe not required
	deletedAt := schema.LookUpField("DeletedAt")
	if deletedAt == nil {
		deletedAt = schema.LookUpField("deleted_at")
		if deletedAt == nil { // pkg soft_delete
			return nil
		}
	}
	return util.New(schemas.FieldFromSchema(deletedAt))
}

func ValidTime(src gorm.DeletedAt, dest time.Time) gorm.DeletedAt {
	if !src.Valid {
		return gorm.DeletedAt(sql.NullTime{Valid: true, Time: dest})
	}
	return src
}
