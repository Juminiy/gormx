UPDATE `tbl_chuck_block` SET `max_size`=4551157041860130406
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 31
--- PASS: TestUpdateByImplicitPk/IsPlugin_DestPk_Updates(Struct)

UPDATE `tbl_chuck_block` SET `max_size`=3136754554848347098
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 32
--- PASS: TestUpdateByImplicitPk/IsPlugin_ModelPk_ModelPtr_Updates(Struct)

UPDATE `tbl_chuck_block` SET `max_size`=6220817006379722895
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 33
--- PASS: TestUpdateByImplicitPk/IsPlugin_ModelPk_ModelValue_Updates(Struct)

UPDATE `tbl_chuck_block` SET `id`=34,`max_size`=3285286720619353829
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 34
--- PASS: TestUpdateByImplicitPk/IsPlugin_ModelPkAndDestPk_ModelPtr_Updates(Struct)

UPDATE `tbl_chuck_block` SET `id`=35,`max_size`=4387253926117090357
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 35
--- PASS: TestUpdateByImplicitPk/IsPlugin_ModelPkAndDestPk_ModelValue_Updates(Struct)

UPDATE `tbl_chuck_block` SET `max_size`=6849951978164470281
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 36
--- PASS: TestUpdateByImplicitPk/IsPlugin_ModelPk_ModelPtr_Updates(Map)

UPDATE `tbl_chuck_block` SET `max_size`=5198644607570523146
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 37
--- PASS: TestUpdateByImplicitPk/IsPlugin_ModelPk_ModelValue_Updates(Map)

UPDATE `tbl_chuck_block` SET `max_size`=3413666573114149029
WHERE `tbl_chuck_block`.`id` = 38 AND `tbl_chuck_block`.`deleted_at` = 0
--- PASS: TestUpdateByImplicitPk/IsPlugin_DestPk_Updates(Map)

UPDATE `tbl_chuck_block` SET `max_size`=4298193345400595493
WHERE `tbl_chuck_block`.`id` = 39 AND `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 39
--- PASS: TestUpdateByImplicitPk/IsPlugin_ModelPkAndDestPk_ModelPtr_Updates(Map)

UPDATE `tbl_chuck_block` SET `max_size`=7522623964701405650
WHERE `tbl_chuck_block`.`id` = 40 AND `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 40
--- PASS: TestUpdateByImplicitPk/IsPlugin_ModelPkAndDestPk_ModelValue_Updates(Map)