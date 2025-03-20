--- DestPk Updates(Struct)
SELECT `updated_at` FROM `tbl_order`
WHERE (`tbl_order`.`id`) IN ((540))
AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order` SET `updated_at`="2025-03-19 22:44:28.114",`pay_time`="2025-03-19 22:44:28.114"
WHERE `tbl_order`.`updated_at` = "2025-03-19 22:44:28.114"
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 540

--- DestPk Updates(Map)
SELECT `updated_at` FROM `tbl_order`
WHERE (`id`) IN ((541))
AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order` SET `pay_time`="2025-03-19 22:44:28.116",`updated_at`="2025-03-19 22:44:28.116"
WHERE `tbl_order`.`updated_at` = "2025-03-19 22:44:28.115"
AND `tbl_order`.`id` = 541
AND `tbl_order`.`deleted_at` IS NULL

--- ModelPk Updates(Struct)
SELECT `updated_at` FROM `tbl_order`
WHERE (`tbl_order`.`id`,`tbl_order`.`serial`) IN ((542,"a527aa87-057f-4477-89ec-836e9ca46a50"))
AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order` SET `updated_at`="2025-03-19 22:44:28.117",`pay_time`="2025-03-19 22:44:28.117"
WHERE `tbl_order`.`updated_at` = "2025-03-19 22:44:28.116"
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 542
AND `serial` = "a527aa87-057f-4477-89ec-836e9ca46a50"

--- ModelPk Updates(Map)
SELECT `updated_at` FROM `tbl_order`
WHERE (`id`,`serial`) IN ((543,"a91e3c27-4caa-4d9a-9e01-0e50e5430fe7"))
AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order` SET `pay_time`="2025-03-19 22:44:28.12",`updated_at`="2025-03-19 22:44:28.12"
WHERE `tbl_order`.`updated_at` = "2025-03-19 22:44:28.119"
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 543
AND `serial` = "a91e3c27-4caa-4d9a-9e01-0e50e5430fe7"


--- IntVersion Updates(Map)
SELECT `version` FROM `tbl_chuck_block`
WHERE (`id`) IN ((14)) AND `tbl_chuck_block`.`deleted_at` = 0

UPDATE `tbl_chuck_block`
SET `min_size`=5813099675382541171,`version`=version + 1
WHERE `tbl_chuck_block`.`version` = 0
AND `tbl_chuck_block`.`id` = 14
AND `tbl_chuck_block`.`deleted_at` = 0

--- IntVersion Updates(Struct)
SELECT `version` FROM `tbl_chuck_block`
WHERE (`id`) IN ((14)) AND `tbl_chuck_block`.`deleted_at` = 0

UPDATE `tbl_chuck_block`
SET `version`=2,`max_size`=6697225754786805900
WHERE `tbl_chuck_block`.`version` = 1
AND `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 14
