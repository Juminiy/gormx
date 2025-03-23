--- Count
SELECT count(*) FROM `tbl_bread_sale`
WHERE `tbl_bread_sale`.`merchant_id` = 13
AND `tbl_bread_sale`.`deleted_at` IS NULL


--- Pluck
SELECT `release_count` FROM `tbl_bread_sale`
WHERE `tbl_bread_sale`.`merchant_id` = 13
AND `tbl_bread_sale`.`deleted_at` IS NULL


--- Dest NotEq Schema.ModelType
SELECT `tbl_bread_sale`.`id`,`tbl_bread_sale`.`first_sale_time`,`tbl_bread_sale`.`last_sale_time`,`tbl_bread_sale`.`sale_count`
FROM `tbl_bread_sale`
WHERE `tbl_bread_sale`.`merchant_id` = 13
AND `tbl_bread_sale`.`deleted_at` IS NULL