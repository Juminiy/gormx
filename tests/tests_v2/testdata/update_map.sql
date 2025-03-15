--- NoPlugin
UPDATE `tbl_order`
SET `logistics_id`=35895,`logistics_name`="Redfin",`shipped_time`="2025-03-22 14:30:57.268",`shipping_fee`=160,`updated_at`="2025-03-15 14:30:57.268"
WHERE `tbl_order`.`deleted_at` IS NULL
AND `id` = 238
AND `serial` = "6ac00138-e235-4705-9f1b-2b7bfb96c8c0"

--- IsPlugin, NoScopes
SELECT count(*) FROM `tbl_order`
WHERE (1!=1
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((230,32606,"Embark")))
AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order`
SET `logistics_id`=32606,`logistics_name`="Embark",`shipped_time`="2025-03-22 14:39:16.914",`shipping_fee`=230,`updated_at`="2025-03-15 14:39:16.914"
WHERE `tbl_order`.`deleted_at` IS NULL
AND `id` = 243
AND `serial` = "d0575d27-f7f2-46cd-aed3-c9d196a4170a"


--- IsPlugin, Scopes(user_id)
SELECT count(*) FROM `tbl_order`
WHERE (1!=1
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((88,10326,"DataMade")))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`user_id` = 666

UPDATE `tbl_order`
SET `logistics_id`=10326,`logistics_name`="DataMade",`shipped_time`="2025-03-22 14:39:16.916",`shipping_fee`=88,`updated_at`="2025-03-15 14:39:16.916"
WHERE `tbl_order`.`user_id` = 666
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 244
AND `serial` = "1bcc8d6f-586e-4837-937e-2bc6810d0639"

--- IsPlugin, Scopes(user_id, tenant_id)
SELECT count(*) FROM `tbl_order`
WHERE (1!=1
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((418,21710,"Rivet Software")))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`user_id` = 666
AND `tbl_order`.`tenant_id` = 888

UPDATE `tbl_order`
SET `logistics_id`=21710,`logistics_name`="Rivet Software",`shipped_time`="2025-03-22 14:39:16.917",`shipping_fee`=418,`updated_at`="2025-03-15 14:39:16.918"
WHERE `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 245
AND `serial` = "51372207-6a78-4c0a-a50e-e4e847066e6f"





----- Bug fix
SELECT count(*) FROM `tbl_order` WHERE (1!=1
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((427,42011,"Foursquare")))
AND `tbl_order`.`deleted_at` IS NULL
AND (`tbl_order`.`id`,`tbl_order`.`serial`) NOT IN ((263,"a4429797-fde7-4a85-8948-50ff2f41f2d5"))

SELECT count(*) FROM `tbl_order` WHERE (1!=1
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((48,29594,"DataWeave")))
AND `tbl_order`.`deleted_at` IS NULL
AND (`tbl_order`.`id`,`tbl_order`.`serial`) NOT IN ((264,"ba7779d7-430c-43af-8f3f-57becd8b13f2"))
AND `tbl_order`.`user_id` = 666

SELECT count(*) FROM `tbl_order` WHERE (1!=1
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((268,22441,"Intermap Technologies")))
AND `tbl_order`.`deleted_at` IS NULL
AND (`tbl_order`.`id`,`tbl_order`.`serial`) NOT IN ((265,"40a75d7a-c78c-48b0-a031-5a544232639e"))
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666

SELECT count(*) FROM `tbl_order` WHERE (1!=1
OR (`shipping_fee`,`logistics_id`,`logistics_name`) IN ((366,43345,"Trintech")))
AND `tbl_order`.`deleted_at` IS NULL
AND ((`tbl_order`.`id`,`tbl_order`.`serial`) NOT IN ((267,"84d6fe4f-b0f0-4699-8c32-c78ff196b086"))
AND NOT id = 267)
AND `tbl_order`.`tenant_id` = 888 AND `tbl_order`.`user_id` = 666
