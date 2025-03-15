--- NoPlugin
INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-14 19:07:13.495","2025-03-14 19:07:13.495",NULL,"f9460104-cc33-4a93-8c6e-c4c04ed648cc",0,0,9427,53,11,9385,2,1,NULL,0,"Kirstin Mitchell","3468414319","57289 North Coursetown, Virginia Beach, Nevada 90085",NULL,NULL,0,"",NULL) RETURNING `id`

UPDATE `tbl_order` SET `deleted_at`="2025-03-14 19:07:13.496"
WHERE `tbl_order`.`id` = 97
AND `tbl_order`.`deleted_at` IS NULL

--- IsPlugin NoScopes
SELECT count(*) FROM `tbl_order` WHERE ((1!=1
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Berenice Kirlin","2212338653","61183 New Courtshaven, Boise, Texas 19064")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((7148,30,193)))
AND `tbl_order`.`deleted_at` IS NULL

INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-14 19:07:13.497","2025-03-14 19:07:13.497",NULL,"7034d9cf-2eb8-4c01-93ae-2a0359469be0",0,0,7148,30,193,7311,1,1,NULL,0,"Berenice Kirlin","2212338653","61183 New Courtshaven, Boise, Texas 19064",NULL,NULL,0,"",NULL) RETURNING `id`

UPDATE `tbl_order` SET `deleted_at`="2025-03-14 19:07:13.498"
WHERE `tbl_order`.`id` = 98
AND `tbl_order`.`deleted_at` IS NULL

--- IsPlugin Scopes(user_id)
SELECT count(*) FROM `tbl_order` WHERE ((1!=1
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((2108,29,112)))
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Trystan Hintz","3499722107","7938 Inletside, Stockton, New Mexico 27281")))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-14 19:07:13.498","2025-03-14 19:07:13.498",NULL,"f6585508-dd53-4c62-8c03-c4af2e1d5473",666,0,2108,29,112,2191,3,1,NULL,0,"Trystan Hintz","3499722107","7938 Inletside, Stockton, New Mexico 27281",NULL,NULL,0,"",NULL) RETURNING `id`

UPDATE `tbl_order` SET `deleted_at`="2025-03-14 19:07:13.499"
WHERE `tbl_order`.`id` = 99
AND `tbl_order`.`user_id` = 666
AND `tbl_order`.`deleted_at` IS NULL

--- IsPlugin Scopes(tenant_id)
SELECT count(*) FROM `tbl_order` WHERE ((1!=1
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((5047,56,217)))
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Rosalinda Ebert","8471266065","570 Turnpikemouth, Glendale, Delaware 50462")))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-14 19:07:13.5","2025-03-14 19:07:13.5",NULL,"9a217f2d-f8ee-43c0-a147-be6db78213d1",666,888,5047,56,217,5208,3,1,NULL,0,"Rosalinda Ebert","8471266065","570 Turnpikemouth, Glendale, Delaware 50462",NULL,NULL,0,"",NULL) RETURNING `id`

UPDATE `tbl_order` SET `deleted_at`="2025-03-14 19:07:13.501"
WHERE `tbl_order`.`id` = 100
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666
AND `tbl_order`.`deleted_at` IS NULL
