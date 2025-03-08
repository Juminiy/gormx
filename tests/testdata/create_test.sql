INSERT INTO `tbl_baby_trade`
(`create_time`,`update_time`,`deleted_at`,`sim_uuid`,`user_id`,`tenant_id`,`auction_id`,`cat_id`,`cat`,`buy_mount`,`day`) VALUES
("2025-03-08 21:14:45.705","2025-03-08 21:14:45.705",0,"bc06526c-57b1-4579-831a-8354b5672d87",0,0,461,1082,0,5635,"20250102")
RETURNING `id`
---

INSERT INTO `tbl_baby_trade`
(`create_time`,`update_time`,`deleted_at`,`sim_uuid`,`user_id`,`tenant_id`,`auction_id`,`cat_id`,`cat`,`buy_mount`,`day`) VALUES
("2025-03-08 21:17:04.874","2025-03-08 21:17:04.874",0,"a90a21eb-228e-4303-9502-f50d939c75cf",114514,114514,985,1436,0,5817,"20250102")
RETURNING `id`
---

SELECT count(*) FROM `tbl_baby_trade`
WHERE ((1!=1 OR (1=1 AND `auction_id` = 930)) OR (1=1 AND `cat_id` = 2979))
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514

INSERT INTO `tbl_baby_trade` (`create_time`,`update_time`,`deleted_at`,`sim_uuid`,`user_id`,`tenant_id`,`auction_id`,`cat_id`,`cat`,`buy_mount`,`day`) VALUES
("2025-03-08 21:24:21.422","2025-03-08 21:24:21.422",0,"4fe07e74-35e3-4af2-aa4e-4efa2da4d72e",114514,114514,930,2979,0,1125,"20250102")
RETURNING `id`
---


