--- Delete(Struct)
UPDATE `tbl_calico_weave`
SET `deleted_at`="2025-03-06 22:08:38.371"
WHERE id = 2
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`user_id` = 114514
AND `tbl_calico_weave`.`deleted_at` IS NULL

--- Delete([]Struct)
UPDATE `tbl_calico_weave`
SET `deleted_at`="2025-03-06 22:08:38.372"
WHERE id = 2
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL

--- Table(table).Delete(Map)
UPDATE `tbl_calico_weave`
SET `deleted_at`="2025-03-06 22:08:38.372"
WHERE id = 2
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL
