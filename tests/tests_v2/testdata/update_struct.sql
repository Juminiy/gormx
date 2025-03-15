--- NoPlugin
UPDATE `tbl_order`
SET `created_at`="2025-03-15 13:39:19.147",`updated_at`="2025-03-15 13:39:19.15",`amount_total`=8550,`amount_discount`=31,`shipping_fee`=128,`amount_actual`=8647,`order_type`=2,`order_status`=2,`pay_time`="2025-03-16 13:39:19.15",`pay_method`=1,`receiver_name`="Brenda Robel",`receiver_phone`="3858642617",`receiver_address`="83776 Crescentmouth, Plano, Maine 69768"
WHERE `tbl_order`.`deleted_at` IS NULL
AND `id` = 213
AND `serial` = "0481e911-27dd-4440-9b44-68d851027a99"

UPDATE `tbl_order`
SET `created_at`="2025-03-15 13:39:19.147",`updated_at`="2025-03-15 13:39:19.15",`amount_total`=8550,`amount_discount`=31,`shipping_fee`=128,`amount_actual`=8647,`order_type`=2,`order_status`=2,`pay_time`="2025-03-16 13:39:19.15",`pay_method`=1,`receiver_name`="Brenda Robel",`receiver_phone`="3858642617",`receiver_address`="83776 Crescentmouth, Plano, Maine 69768",`shipped_time`="2025-03-22 13:39:19.15",`logistics_id`=1658,`logistics_name`="Redfin"
WHERE `tbl_order`.`deleted_at` IS NULL
AND `id` = 213
AND `serial` = "0481e911-27dd-4440-9b44-68d851027a99"

--- IsPlugin, NoScopes
SELECT count(*) FROM `tbl_order`
WHERE (((1!=1
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Oma Reichert","1120766625","869 North Knollstad, Oakland, Kentucky 54180")))
OR (`serial`) IN (("fcc7f44c-f55b-4d61-b556-17c1157b19dd")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((8815,55,358)))
AND `tbl_order`.`deleted_at` IS NULL
AND NOT ((`tbl_order`.`id` = 214 AND `tbl_order`.`serial` = "fcc7f44c-f55b-4d61-b556-17c1157b19dd") AND 1=1)

UPDATE `tbl_order`
SET `created_at`="2025-03-15 13:39:19.152",`updated_at`="2025-03-15 13:39:19.153",`amount_total`=8815,`amount_discount`=55,`shipping_fee`=358,`amount_actual`=9118,`order_type`=1,`order_status`=2,`pay_time`="2025-03-16 13:39:19.153",`pay_method`=1,`receiver_name`="Oma Reichert",`receiver_phone`="1120766625",`receiver_address`="869 North Knollstad, Oakland, Kentucky 54180"
WHERE `tbl_order`.`deleted_at` IS NULL
AND `id` = 214
AND `serial` = "fcc7f44c-f55b-4d61-b556-17c1157b19dd"

SELECT count(*) FROM `tbl_order`
WHERE ((((1!=1
OR (`serial`) IN (("fcc7f44c-f55b-4d61-b556-17c1157b19dd")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((8815,55,358)))
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((358,17044,"GetRaised")))
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Oma Reichert","1120766625","869 North Knollstad, Oakland, Kentucky 54180")))
AND `tbl_order`.`deleted_at` IS NULL
AND NOT ((`tbl_order`.`id` = 214 AND `tbl_order`.`serial` = "fcc7f44c-f55b-4d61-b556-17c1157b19dd") AND 1=1)

UPDATE `tbl_order`
SET `created_at`="2025-03-15 13:39:19.152",`updated_at`="2025-03-15 13:39:19.154",`amount_total`=8815,`amount_discount`=55,`shipping_fee`=358,`amount_actual`=9118,`order_type`=1,`order_status`=2,`pay_time`="2025-03-16 13:39:19.153",`pay_method`=1,`receiver_name`="Oma Reichert",`receiver_phone`="1120766625",`receiver_address`="869 North Knollstad, Oakland, Kentucky 54180",`shipped_time`="2025-03-22 13:39:19.153",`logistics_id`=17044,`logistics_name`="GetRaised"
WHERE `tbl_order`.`deleted_at` IS NULL
AND `id` = 214
AND `serial` = "fcc7f44c-f55b-4d61-b556-17c1157b19dd"

--- IsPlugin, Scopes(user_id)
SELECT count(*) FROM `tbl_order`
WHERE (((1!=1
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Betsy Abernathy","7268142312","294 Viaductbury, Cleveland, Michigan 35671")))
OR (`serial`) IN (("10a6aff9-b01d-4870-8f22-70eac487f8db")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((9268,10,267)))
AND `tbl_order`.`deleted_at` IS NULL
AND NOT ((`tbl_order`.`id` = 215 AND `tbl_order`.`serial` = "10a6aff9-b01d-4870-8f22-70eac487f8db") AND 1=1)
AND `tbl_order`.`user_id` = 666

UPDATE `tbl_order`
SET `created_at`="2025-03-15 13:39:19.154",`updated_at`="2025-03-15 13:39:19.155",`amount_total`=9268,`amount_discount`=10,`shipping_fee`=267,`amount_actual`=9525,`order_type`=2,`order_status`=2,`pay_time`="2025-03-16 13:39:19.155",`pay_method`=2,`receiver_name`="Betsy Abernathy",`receiver_phone`="7268142312",`receiver_address`="294 Viaductbury, Cleveland, Michigan 35671"
WHERE `tbl_order`.`user_id` = 666
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 215
AND `serial` = "10a6aff9-b01d-4870-8f22-70eac487f8db"

SELECT count(*) FROM `tbl_order`
WHERE ((((1!=1
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Betsy Abernathy","7268142312","294 Viaductbury, Cleveland, Michigan 35671")))
OR (`serial`) IN (("10a6aff9-b01d-4870-8f22-70eac487f8db")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((9268,10,267)))
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((267,11603,"Garmin")))
AND `tbl_order`.`deleted_at` IS NULL
AND NOT ((`tbl_order`.`id` = 215 AND `tbl_order`.`serial` = "10a6aff9-b01d-4870-8f22-70eac487f8db") AND 1=1)
AND `tbl_order`.`user_id` = 666

UPDATE `tbl_order`
SET `created_at`="2025-03-15 13:39:19.154",`updated_at`="2025-03-15 13:39:19.156",`amount_total`=9268,`amount_discount`=10,`shipping_fee`=267,`amount_actual`=9525,`order_type`=2,`order_status`=2,`pay_time`="2025-03-16 13:39:19.155",`pay_method`=2,`receiver_name`="Betsy Abernathy",`receiver_phone`="7268142312",`receiver_address`="294 Viaductbury, Cleveland, Michigan 35671",`shipped_time`="2025-03-22 13:39:19.155",`logistics_id`=11603,`logistics_name`="Garmin"
WHERE `tbl_order`.`user_id` = 666
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 215
AND `serial` = "10a6aff9-b01d-4870-8f22-70eac487f8db"

--- IsPlugin, Scopes(user_id, tenant_id)
SELECT count(*) FROM `tbl_order`
WHERE (((1!=1 OR (`serial`) IN (("9da01315-c43a-4985-98c0-ce9dfedf9b62")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((1827,92,441)))
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Orlo Aufderhar","6350679858","166 Lake Pointview, Reno, Alabama 66246")))
AND `tbl_order`.`deleted_at` IS NULL
AND NOT ((`tbl_order`.`id` = 216 AND `tbl_order`.`serial` = "9da01315-c43a-4985-98c0-ce9dfedf9b62") AND 1=1)
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666

UPDATE `tbl_order`
SET `created_at`="2025-03-15 13:39:19.157",`updated_at`="2025-03-15 13:39:19.157",`amount_total`=1827,`amount_discount`=92,`shipping_fee`=441,`amount_actual`=2176,`order_type`=4,`order_status`=2,`pay_time`="2025-03-16 13:39:19.157",`pay_method`=1,`receiver_name`="Orlo Aufderhar",`receiver_phone`="6350679858",`receiver_address`="166 Lake Pointview, Reno, Alabama 66246"
WHERE `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 216
AND `serial` = "9da01315-c43a-4985-98c0-ce9dfedf9b62"

SELECT count(*) FROM `tbl_order`
WHERE ((((1!=1 OR (`serial`) IN (("9da01315-c43a-4985-98c0-ce9dfedf9b62")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((1827,92,441)))
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((441,53234,"Allianz")))
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Orlo Aufderhar","6350679858","166 Lake Pointview, Reno, Alabama 66246")))
AND `tbl_order`.`deleted_at` IS NULL
AND NOT ((`tbl_order`.`id` = 216 AND `tbl_order`.`serial` = "9da01315-c43a-4985-98c0-ce9dfedf9b62") AND 1=1)
AND `tbl_order`.`user_id` = 666
AND `tbl_order`.`tenant_id` = 888

UPDATE `tbl_order`
SET `created_at`="2025-03-15 13:39:19.157",`updated_at`="2025-03-15 13:39:19.158",`amount_total`=1827,`amount_discount`=92,`shipping_fee`=441,`amount_actual`=2176,`order_type`=4,`order_status`=2,`pay_time`="2025-03-16 13:39:19.157",`pay_method`=1,`receiver_name`="Orlo Aufderhar",`receiver_phone`="6350679858",`receiver_address`="166 Lake Pointview, Reno, Alabama 66246",`shipped_time`="2025-03-22 13:39:19.158",`logistics_id`=53234,`logistics_name`="Allianz"
WHERE `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 216
AND `serial` = "9da01315-c43a-4985-98c0-ce9dfedf9b62"




----- Bug fix
SELECT count(*) FROM `tbl_order` WHERE (1!=1
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((386,8355,"Genability")))
AND `tbl_order`.`deleted_at` IS NULL
AND (`tbl_order`.`id`,`tbl_order`.`serial`) NOT IN ((259,"8413e9a4-25cd-45ce-b52f-e298c65c94dd"))

SELECT count(*) FROM `tbl_order` WHERE (1!=1
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((467,42735,"Informatica")))
AND `tbl_order`.`deleted_at` IS NULL
AND (`tbl_order`.`id`,`tbl_order`.`serial`) NOT IN ((260,"c65df3a4-1a4d-4c59-a16a-e45eeee3ca14"))
AND `tbl_order`.`user_id` = 666

SELECT count(*) FROM `tbl_order` WHERE (1!=1
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((449,51006,"Golden Helix")))
AND `tbl_order`.`deleted_at` IS NULL
AND (`tbl_order`.`id`,`tbl_order`.`serial`) NOT IN ((261,"043c6b94-ca2d-4c99-ad26-cd918090ed4f"))
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666

SELECT count(*) FROM `tbl_order`
WHERE ((((1!=1
OR (`serial`) IN (("2c39a79f-698f-472e-b8ef-a12d1e6fdcd3")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((5743,1,294)))
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((294,37848,"Healthline")))
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Jedidiah Weber","8881787854","95417 Inletside, Newark, Michigan 85772")))
AND `tbl_order`.`deleted_at` IS NULL
AND ((`tbl_order`.`id`,`tbl_order`.`serial`) NOT IN ((266,"2c39a79f-698f-472e-b8ef-a12d1e6fdcd3"))
AND NOT id = 266)
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666
