
SELECT * FROM `tbl_order` WHERE `tbl_order`.`deleted_at` IS NULL LIMIT 3

SELECT * FROM `tbl_order` WHERE `tbl_order`.`deleted_at` IS NULL LIMIT 3

SELECT * FROM `tbl_order` WHERE `tbl_order`.`user_id` = 666
AND `tbl_order`.`deleted_at` IS NULL LIMIT 3

SELECT * FROM `tbl_order` WHERE `tbl_order`.`user_id` = 666
AND `tbl_order`.`deleted_at` IS NULL LIMIT 3

SELECT * FROM `tbl_order` WHERE `tbl_order`.`user_id` = 666
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`deleted_at` IS NULL LIMIT 3

SELECT * FROM `tbl_order` WHERE `tbl_order`.`user_id` = 666
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`deleted_at` IS NULL LIMIT 3
