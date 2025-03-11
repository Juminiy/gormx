--- delete returning
SELECT * FROM `tbl_baby_trade`
WHERE `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`id` = 32021
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514

UPDATE `tbl_baby_trade` SET `deleted_at`=1741693454
WHERE `tbl_baby_trade`.`id` = 32021
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514
AND (`tbl_baby_trade`.`id`,`tbl_baby_trade`.`sim_uuid`) IN ((32021,'2420bfab-4dbe-4bc6-9717-e875066d4165'))
AND `tbl_baby_trade`.`deleted_at` = 0
---

--- update returning
SELECT count(*) FROM `tbl_baby_trade`
WHERE (((1!=1 OR (`sim_uuid`) IN ((NULL))) OR (`auction_id`,`cat_id`) IN ((10,11))) OR (`cat_id`,`cat`) IN ((11,NULL)))
AND `tbl_baby_trade`.`deleted_at` = 0
AND NOT (`tbl_baby_trade`.`id` = 32023
AND `tbl_baby_trade`.`sim_uuid` = 'bc06526c-57b1-4579-831a-8354b5672d87' AND 1=1)
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514

SELECT * FROM `tbl_baby_trade`
WHERE `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`id` = 32023
AND `tbl_baby_trade`.`sim_uuid` = 'bc06526c-57b1-4579-831a-8354b5672d87'
AND `tbl_baby_trade`.`user_id` = 114514
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`deleted_at` = 0

UPDATE `tbl_baby_trade`
SET `auction_id`=10,`buy_mount`=0,`cat_id`=11,`update_time`='2025-03-11 19:59:48.216'
WHERE `tbl_baby_trade`.`id` = 32023
AND `tbl_baby_trade`.`sim_uuid` = 'bc06526c-57b1-4579-831a-8354b5672d87'
AND `tbl_baby_trade`.`user_id` = 114514
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`deleted_at` = 0
