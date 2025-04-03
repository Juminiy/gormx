package gormx_testv2

import (
	"github.com/Juminiy/kube/pkg/util"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"testing"
)

func TestAsyncTxnAllTaskInFunc(t *testing.T) {
	err := iSqlite0().Transaction(func(tx *gorm.DB) error {
		eg, ctx := errgroup.WithContext(util.TODOContext())
		defer ctx.Done()

		var bread *BreadProduct
		eg.Go(func() error {
			bread = RandomBread()
			return tx.Create(&bread).Error
		})

		var sale *BreadSale
		eg.Go(func() error {
			sale = RandomBreadSale(RandomBread())
			return tx.Where("id = ?", 1).Updates(&sale).Error
		})

		var saleList []BreadSale
		eg.Go(func() error {
			return tx.Find(&saleList, "id BETWEEN ? AND ?", 10, 20).Error
		})

		return eg.Wait()
	})
	if err != nil {
		t.Logf("async transaction case (all async in a txn) error: %s", err.Error())
		return
	}
	t.Log("async transaction case (all sync in a txn) ok")
}

func TestAsyncTxnDryRun(t *testing.T) {
	tx := iSqlite0()
	eg, ctx := errgroup.WithContext(util.TODOContext())
	defer ctx.Done()

	var bread *BreadProduct
	eg.Go(func() error {
		bread = RandomBread()
		return tx.Create(&bread).Error
	})

	var sale *BreadSale
	eg.Go(func() error {
		sale = RandomBreadSale(RandomBread())
		return tx.Where("id = ?", 1).Updates(&sale).Error
	})

	var saleList []BreadSale
	eg.Go(func() error {
		return tx.Find(&saleList, "id BETWEEN ? AND ?", 10, 20).Error
	})

	err := eg.Wait()
	if err != nil {
		t.Logf("async transaction case (all async diff txn) error: %s", err.Error())
		return
	}
	t.Log("async transaction case (all sync in diff txn) ok")
}
