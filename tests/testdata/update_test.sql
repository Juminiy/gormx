SELECT count(*) FROM `tbl_baby_trade`
WHERE ((1!=1 OR (1=1 AND `auction_id` = 10)) OR (1=1 AND `cat_id` = 11))
AND `tbl_baby_trade`.`deleted_at` = 0
AND NOT `tbl_baby_trade`.`id` = 1

UPDATE `tbl_baby_trade`
SET `auction_id`=10,`cat_id`=11,`update_time`="2025-03-08 21:48:29.596"
WHERE `tbl_baby_trade`.`id` = 1
AND `tbl_baby_trade`.`deleted_at` = 0
---

SELECT count(*) FROM `tbl_baby_trade`
WHERE ((1!=1 OR (1=1 AND `auction_id` = 10)) OR (1=1 AND `cat_id` = 11))
AND `tbl_baby_trade`.`deleted_at` = 0
AND NOT `tbl_baby_trade`.`id` = 1

UPDATE `tbl_baby_trade`
SET `auction_id`=10,`buy_mount`=0,`cat_id`=11,`update_time`="2025-03-08 22:11:50.567"
WHERE `tbl_baby_trade`.`id` = 1
AND `tbl_baby_trade`.`deleted_at` = 0
---

SELECT count(*) FROM `tbl_baby_trade`
WHERE ((1!=1 OR (1=1 AND `auction_id` = 10)) OR (1=1 AND `cat_id` = 11))
AND `tbl_baby_trade`.`deleted_at` = 0
AND NOT (`tbl_baby_trade`.`id` = 1 AND `tbl_baby_trade`.`sim_uuid` = "bc06526c-57b1-4579-831a-8354b5672d87" AND 1=1)

UPDATE `tbl_baby_trade`
SET `auction_id`=10,`cat_id`=11,`update_time`="2025-03-08 22:33:57.864"
WHERE `tbl_baby_trade`.`id` = 1
AND `tbl_baby_trade`.`sim_uuid` = "bc06526c-57b1-4579-831a-8354b5672d87"
AND `tbl_baby_trade`.`deleted_at` = 0
---

SELECT count(*) FROM `tbl_baby_trade`
WHERE ((1!=1 OR (1=1 AND `auction_id` = 10)) OR (1=1 AND `cat_id` = 11))
AND `tbl_baby_trade`.`deleted_at` = 0
AND NOT ((`tbl_baby_trade`.`id` = 1 AND `tbl_baby_trade`.`sim_uuid` = "bc06526c-57b1-4579-831a-8354b5672d87") AND 1=1)

UPDATE `tbl_baby_trade`
SET `update_time`="2025-03-08 22:37:41.08",`auction_id`=10,`cat_id`=11
WHERE `tbl_baby_trade`.`deleted_at` = 0
AND `id` = 1
AND `sim_uuid` = "bc06526c-57b1-4579-831a-8354b5672d87"
