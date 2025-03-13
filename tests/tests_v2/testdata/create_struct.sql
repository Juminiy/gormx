--- NoPlugin
INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-13 19:58:44.465","2025-03-13 19:58:44.465",NULL,"8da83d69-1708-42ce-bc93-0362dfbbac95",0,0,9128,77,251,9302,2,1,NULL,0,"Dixie Raynor","1445641758","107 Passageport, Milwaukee, Delaware 81956",NULL,NULL,0,"",NULL) RETURNING `id`

--- IsPlugin
SELECT count(*) FROM `tbl_order`
WHERE ((1!=1
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Amely Yundt","3474007559","3643 Manorsshire, Lexington-Fayette, New Mexico 14205")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((1434,42,374)))
AND `tbl_order`.`deleted_at` IS NULL

INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-13 19:58:44.47","2025-03-13 19:58:44.47",NULL,"11c52803-c786-49d5-aa3f-6b8788e92934",0,0,1434,42,374,1766,1,1,NULL,0,"Amely Yundt","3474007559","3643 Manorsshire, Lexington-Fayette, New Mexico 14205",NULL,NULL,0,"",NULL) RETURNING `id`

--- IsPlugin Scopes(user_id)
SELECT count(*) FROM `tbl_order`
WHERE ((1!=1
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((5003,67,313)))
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Raphaelle Crona","9479408352","438 Creekstad, Wichita, Indiana 45217")))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-13 20:24:12.927","2025-03-13 20:24:12.927",NULL,"2c39fdbd-a4ba-4a85-abbe-c780284bae18",666,0,5003,67,313,5249,2,1,NULL,0,"Raphaelle Crona","9479408352","438 Creekstad, Wichita, Indiana 45217",NULL,NULL,0,"",NULL) RETURNING `id`

--- IsPlugin Scopes(user_id, tenant_id)
SELECT count(*) FROM `tbl_order`
WHERE ((1!=1
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Carleton Lind","7978078285","3607 Crestshire, Chula Vista, North Dakota 32597")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((7994,88,309)))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`user_id` = 666
AND `tbl_order`.`tenant_id` = 888

INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-13 20:24:12.93","2025-03-13 20:24:12.93",NULL,"6069d86f-2fd1-45a8-8489-6f415ce10b3c",666,888,7994,88,309,8215,4,1,NULL,0,"Carleton Lind","7978078285","3607 Crestshire, Chula Vista, North Dakota 32597",NULL,NULL,0,"",NULL) RETURNING `id`
