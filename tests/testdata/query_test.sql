SELECT count(*) FROM `tbl_baby_trade`
WHERE ((1!=1 OR (1=1 AND `auction_id` = 840)) OR (1=1 AND `cat_id` = 15))
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514

INSERT INTO `tbl_baby_trade`
(`create_time`,`update_time`,`deleted_at`,`sim_uuid`,`user_id`,`tenant_id`,`auction_id`,`cat_id`,`cat`,`buy_mount`,`day`) VALUES
("2025-03-09 10:55:35.241","2025-03-09 10:55:35.241",0,"284a8889-46bf-4e1e-babe-6a7f67ff53ec",114514,114514,840,15,0,29304,"1931-03-12 19:27:59.979540202 +0000 UTC")
RETURNING `id`

SELECT * FROM `tbl_baby_trade`
WHERE `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`id` = 14
ORDER BY `tbl_baby_trade`.`id`
LIMIT 1
---

SELECT * FROM `tbl_baby_trade`
WHERE `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`id` = 17
ORDER BY `tbl_baby_trade`.`id`
LIMIT 1

SELECT `tbl_baby_trade`.`create_time`,`tbl_baby_trade`.`update_time`,
`tbl_baby_trade`.`deleted_at`,
`tbl_baby_trade`.`id`,`tbl_baby_trade`.`sim_uuid`,
`tbl_baby_trade`.`user_id`,`tbl_baby_trade`.`tenant_id`,
`tbl_baby_trade`.`auction_id`,`tbl_baby_trade`.`cat_id`,
`tbl_baby_trade`.`cat`,`tbl_baby_trade`.`day`
FROM `tbl_baby_trade`
WHERE `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`id` = 18
ORDER BY `tbl_baby_trade`.`id`
LIMIT 1
