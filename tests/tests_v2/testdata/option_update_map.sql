--- 1. NoPlugin, 2. IsPlugin, NoOption, 3. IsPlugin, IsOption

----- OmitUnknownKey
--- no such column: no_column
UPDATE `tbl_order`
SET `no_column`=1
WHERE id = 1

--- no such column: no_column
UPDATE `tbl_order`
SET `no_column`=1,`updated_at`="2025-03-16 12:04:19.63"
WHERE id = 1
AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order` SET `updated_at`="2025-03-16 12:04:19.63"
WHERE id = 1
AND `tbl_order`.`deleted_at` IS NULL



----- OmitZeroElem
--- no such column: logistics
UPDATE `tbl_order`
SET `logistics`="",`logistics_name`="",`shipping_fee`=0
WHERE id = 1

--- no such column: logistics
UPDATE `tbl_order`
SET `logistics`="",`logistics_name`="",`shipping_fee`=0,`updated_at`="2025-03-16 12:27:49.183"
WHERE id = 1
AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order` SET `updated_at`="2025-03-16 12:27:49.183"
WHERE id = 1
AND `tbl_order`.`deleted_at` IS NULL



----- SetPkToClause
--- WHERE conditions required
UPDATE `tbl_order`
SET `id`=1,`logistics`="",`logistics_name`="",`serial`="ff1bffdf-423d-471d-ac39-e108452d210d",`shipping_fee`=0

--- WHERE conditions required
UPDATE `tbl_order`
SET `id`=1,`logistics`="",`logistics_name`="",`serial`="ff1bffdf-423d-471d-ac39-e108452d210d",`shipping_fee`=0,
`updated_at`="2025-03-16 12:47:28.775"
WHERE `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order`
SET `updated_at`="2025-03-16 12:47:28.776"
WHERE `tbl_order`.`id` = 1
AND `tbl_order`.`serial` = "ff1bffdf-423d-471d-ac39-e108452d210d"
AND `tbl_order`.`deleted_at` IS NULL



--- CallHooks
UPDATE `tbl_order` SET `shipped_time`="2025-03-16 13:14:51.116" WHERE id = 1

UPDATE `tbl_order` SET `shipped_time`="2025-03-16 13:14:51.117",
`updated_at`="2025-03-16 13:14:51.117"
WHERE id = 1
AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order` SET `shipped_time`="2025-03-16 13:14:51.118",
`updated_at`="2025-03-16 13:14:51.118"
WHERE id = 1
AND `serial` = "17f3e4d1-f493-4a12-b110-9e4523949a5f"
AND `tbl_order`.`deleted_at` IS NULL



----- ClauseRetuning
UPDATE `tbl_order`
SET `amount_discount`=amount_discount + 10,`order_status`=2,`pay_method`=2,
`pay_time`='2025-03-16 13:47:29.663',`updated_at`='2025-03-16 13:47:29.672'
WHERE `tbl_order`.`deleted_at` IS NULL
AND `id` = 1

UPDATE `tbl_order`
SET `amount_discount`=amount_discount + 10,`order_status`=2,`pay_method`=2,
`pay_time`='2025-03-16 13:47:29.663',`updated_at`='2025-03-16 13:47:29.713'
WHERE `tbl_order`.`deleted_at` IS NULL
AND `id` = 1

UPDATE `tbl_order`
SET `amount_discount`=amount_discount + 10,`order_status`=2,`pay_method`=2,
`pay_time`='2025-03-16 13:47:29.663',`updated_at`='2025-03-16 13:47:29.761'
WHERE `tbl_order`.`deleted_at` IS NULL
AND `id` = 1

SELECT * FROM `tbl_order` WHERE `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 1
