--- Update(Struct)
SELECT count(*) FROM `tbl_calico_weave`
WHERE (1!=1 OR (1=1 AND `name` = "MyName"))
AND `tbl_calico_weave`.`deleted_at` IS NULL

UPDATE `tbl_calico_weave`
SET `updated_at`="2025-03-06 21:22:33.682",`name`="MyName"
WHERE `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`user_id` = 114514
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND `id` = 2

-- Model(table).Updates(Map)
SELECT count(*) FROM `tbl_calico_weave`
WHERE (1!=1 OR (1=1 AND `name` = "MyName"))
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND NOT `tbl_calico_weave`.`id` = 2

UPDATE `tbl_calico_weave`
SET `name`="MyName",`updated_at`="2025-03-06 21:22:33.683"
WHERE `tbl_calico_weave`.`id` = 2
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`user_id` = 114514
AND `tbl_calico_weave`.`deleted_at` IS NULL
