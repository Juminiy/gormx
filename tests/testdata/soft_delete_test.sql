UPDATE `tbl_baby_trade`
SET `cat_id`=254,`update_time`="2025-03-09 11:09:26.758"
WHERE `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514
AND `tbl_baby_trade`.`deleted_at` = 0

UPDATE `tbl_baby_trade` SET `deleted_at`=1741489766
WHERE `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514
AND `tbl_baby_trade`.`deleted_at` = 0
---

    --gorm_test.go:64: update tenant all rows or global update is not allowed
    --gorm_test.go:64: delete tenant all rows or global update is not allowed

UPDATE `tbl_baby_trade`
SET `deleted_at`=1741509794
WHERE `tbl_baby_trade`.`id` = 21
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514
AND `tbl_baby_trade`.`deleted_at` = 0
RETURNING *
