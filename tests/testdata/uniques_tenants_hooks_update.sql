SELECT count(*) FROM `tbl_calico_weave`
WHERE (1!=1 OR (1=1 AND `name` = "MyName-2"))
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND `tbl_calico_weave`.`tenant_id` = 1919810

UPDATE `tbl_calico_weave`
SET `updated_at`="2025-03-08 18:09:48.118",`name`="MyName-2"
WHERE `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`user_id` = 114514
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND `id` = 2
---

SELECT count(*) FROM `tbl_calico_weave`
WHERE (1!=1 OR (1=1 AND `name` = "MyName-2"))
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND NOT `tbl_calico_weave`.`id` = "2"
AND `tbl_calico_weave`.`tenant_id` = 1919810

UPDATE `tbl_calico_weave`
SET `name`="MyName-2",`updated_at`="2025-03-08 18:09:48.118"
WHERE `tbl_calico_weave`.`id` = "2"
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`user_id` = 114514
AND `tbl_calico_weave`.`deleted_at` IS NULL
