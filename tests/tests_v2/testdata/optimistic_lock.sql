--- NoPk Struct
UPDATE `tbl_order` SET `updated_at`="2025-03-19 22:27:44.194",`pay_time`="2025-03-19 22:27:44.194"
WHERE id = 528 AND `tbl_order`.`deleted_at` IS NULL
--- NoPk Map
UPDATE `tbl_order` SET `pay_time`="2025-03-19 22:27:44.194",`updated_at`="2025-03-19 22:27:44.194"
WHERE id = 528 AND `tbl_order`.`deleted_at` IS NULL

--- DestPk Updates(Struct)
SELECT `updated_at` FROM `tbl_order`
WHERE (`tbl_order`.`id`) IN ((529)) AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order` SET `updated_at`="2025-03-19 22:27:44.195",`pay_time`="2025-03-19 22:27:44.195"
WHERE `tbl_order`.`updated_at` = "2025-03-19 22:27:44.195"
AND `tbl_order`.`deleted_at` IS NULL AND `id` = 529

--- DestPk Updates(Map)
SELECT `updated_at` FROM `tbl_order`
WHERE (`id`) = 530 AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order` SET `pay_time`="2025-03-19 22:27:44.196",`updated_at`="2025-03-19 22:27:44.196"
WHERE `tbl_order`.`updated_at` = "2025-03-19 22:27:44.196"
AND `tbl_order`.`id` = 530 AND `tbl_order`.`deleted_at` IS NULL

--- ModelPk Updates(Struct)
SELECT `updated_at` FROM `tbl_order`
WHERE (`tbl_order`.`id`,`tbl_order`.`serial`) IN ((531,"aad662d2-ae25-48ad-a121-84b6d145045c"))
AND `tbl_order`.`deleted_at` IS NULL

UPDATE `tbl_order` SET `updated_at`="2025-03-19 22:27:44.198",`pay_time`="2025-03-19 22:27:44.197"
WHERE `tbl_order`.`updated_at` = "2025-03-19 22:27:44.197"
AND `tbl_order`.`deleted_at` IS NULL
AND `id` = 531 AND `serial` = "aad662d2-ae25-48ad-a121-84b6d145045c"

--- ModelPK Updates(Map)
UPDATE `tbl_order` SET `pay_time`="2025-03-19 22:27:44.199",`updated_at`="2025-03-19 22:27:44.199"
WHERE `tbl_order`.`deleted_at` IS NULL
AND `id` = 532 AND `serial` = "6f35f133-f6ec-48eb-a64b-01c5b45a9426"
