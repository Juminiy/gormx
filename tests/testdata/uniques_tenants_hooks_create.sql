SELECT count(*) FROM `tbl_calico_weave`
WHERE (1!=1 OR (1=1 AND `name` = "MyName"))
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND NOT `tbl_calico_weave`.`id` = 2
AND `tbl_calico_weave`.`tenant_id` = 1919810


UPDATE `tbl_calico_weave`
SET `updated_at`="2025-03-08 19:15:16.336",`name`="MyName"
WHERE `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`user_id` = 114514
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND `id` = 2
---

SELECT count(*) FROM `tbl_calico_weave`
WHERE (1!=1 OR (1=1 AND `name` = "MyName"))
AND `tbl_calico_weave`.`deleted_at` IS NULL
AND NOT `tbl_calico_weave`.`id` = 2
AND `tbl_calico_weave`.`tenant_id` = 1919810

UPDATE `tbl_calico_weave`
SET `name`="MyName",`updated_at`="2025-03-08 19:15:16.337"
WHERE `tbl_calico_weave`.`id` = 2
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`user_id` = 114514
AND `tbl_calico_weave`.`deleted_at` IS NULL
---