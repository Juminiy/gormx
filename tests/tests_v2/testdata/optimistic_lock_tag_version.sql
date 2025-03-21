----- IntVersion

--- DestPk Updates(Struct)
SELECT `version` FROM `tbl_chuck_block`
WHERE (`id`) IN ((17))
AND `tbl_chuck_block`.`deleted_at` = 0

UPDATE `tbl_chuck_block`
SET `version`=1,`min_size`=3202331964011572402
WHERE `tbl_chuck_block`.`version` = 0
AND `tbl_chuck_block`.`deleted_at` = 0
AND `id` = 17

--- DestPk Updates(Map)
SELECT `version` FROM `tbl_chuck_block`
WHERE (`id`) IN ((18))
AND `tbl_chuck_block`.`deleted_at` = 0

UPDATE `tbl_chuck_block`
SET `max_size`=2371010289426027463,`version`=version + 1
WHERE `tbl_chuck_block`.`version` = 0
AND `tbl_chuck_block`.`id` = 18
AND `tbl_chuck_block`.`deleted_at` = 0

--- ModelPk Model(Struct).Updates(Struct)
SELECT `version` FROM `tbl_chuck_block`
WHERE (`id`) IN ((19))
AND `tbl_chuck_block`.`deleted_at` = 0

UPDATE `tbl_chuck_block`
SET `version`=1,`idle_size`=8884958695632747392
WHERE `tbl_chuck_block`.`version` = 0
AND `tbl_chuck_block`.`deleted_at` = 0
AND `id` = 19

--- ModelPk Model(Struct).Updates(Map)
SELECT `version` FROM `tbl_chuck_block`
WHERE (`id`) IN ((20))
AND `tbl_chuck_block`.`deleted_at` = 0

UPDATE `tbl_chuck_block`
SET `version`=version + 1,`wise_desc`="menorah"
WHERE `tbl_chuck_block`.`version` = 0
AND `tbl_chuck_block`.`deleted_at` = 0
AND `id` = 20
