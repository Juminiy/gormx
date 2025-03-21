UPDATE `tbl_chuck_block` SET `max_size`=3682892642066195465
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 21
--- PASS: TestUpdateByImplicitPk/NoPlugin_DestPk_Updates(Struct)

UPDATE `tbl_chuck_block` SET `max_size`=6090406364517231709
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 22
--- PASS: TestUpdateByImplicitPk/NoPlugin_ModelPk_ModelPtr_Updates(Struct)

UPDATE `tbl_chuck_block` SET `max_size`=3771426930665327058
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 23
--- PASS: TestUpdateByImplicitPk/NoPlugin_ModelPk_ModelValue_Updates(Struct)

UPDATE `tbl_chuck_block` SET `id`=24,`max_size`=5064481392337075047
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 24
--- PASS: TestUpdateByImplicitPk/NoPlugin_ModelPkAndDestPk_ModelPtr_Updates(Struct)

UPDATE `tbl_chuck_block` SET `id`=25,`max_size`=2779704806201686504
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 25
--- PASS: TestUpdateByImplicitPk/NoPlugin_ModelPkAndDestPk_ModelValue_Updates(Struct)

UPDATE `tbl_chuck_block` SET `max_size`=4740364344376120767
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 26
--- PASS: TestUpdateByImplicitPk/NoPlugin_ModelPk_ModelPtr_Updates(Map)

UPDATE `tbl_chuck_block` SET `max_size`=2937300446561575155
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 27
--- PASS: TestUpdateByImplicitPk/NoPlugin_ModelPk_ModelValue_Updates(Map)

--- WHERE conditions required
UPDATE `tbl_chuck_block` SET `id`=28,`max_size`=4719907600833420063
--- NoPlugin not support UpdateMap Pk in Map

UPDATE `tbl_chuck_block` SET `id`=29,`max_size`=3427064225313845555
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 29
--- PASS: TestUpdateByImplicitPk/NoPlugin_ModelPkAndDestPk_ModelPtr_Updates(Map)

UPDATE `tbl_chuck_block` SET `id`=30,`max_size`=8447249890238289265
WHERE `tbl_chuck_block`.`deleted_at` = 0 AND `id` = 30
--- PASS: TestUpdateByImplicitPk/NoPlugin_ModelPkAndDestPk_ModelValue_Updates(Map)