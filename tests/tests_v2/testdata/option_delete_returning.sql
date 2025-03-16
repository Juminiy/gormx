UPDATE `tbl_order`
SET `deleted_at`='2025-03-16 13:51:46.627'
WHERE `tbl_order`.`id` = 1
AND `tbl_order`.`deleted_at` IS NULL

SELECT * FROM `tbl_order`
WHERE `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`id` = 1

