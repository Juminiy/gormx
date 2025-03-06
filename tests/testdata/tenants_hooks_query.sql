SELECT `tbl_calico_weave`.`id`,
`tbl_calico_weave`.`created_at`,`tbl_calico_weave`.`updated_at`,
`tbl_calico_weave`.`deleted_at`,`tbl_calico_weave`.`tenant_id`,
`tbl_calico_weave`.`user_id`,`tbl_calico_weave`.`name`,`tbl_calico_weave`.`desc`,
`tbl_calico_weave`.`pumping`,`tbl_calico_weave`.`elephant`,
`tbl_calico_weave`.`loc_id`,`tbl_calico_weave`.`app_id`,
`tbl_calico_weave`.`app_me`,`tbl_calico_weave`.`app_yr`
FROM `tbl_calico_weave`
WHERE `tbl_calico_weave`.`id` = 3
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL
ORDER BY `tbl_calico_weave`.`id`
LIMIT 1

SELECT `tbl_calico_weave`.`id`,
`tbl_calico_weave`.`created_at`,`tbl_calico_weave`.`updated_at`,
`tbl_calico_weave`.`deleted_at`,`tbl_calico_weave`.`tenant_id`,
`tbl_calico_weave`.`user_id`,`tbl_calico_weave`.`name`,`tbl_calico_weave`.`desc`,
`tbl_calico_weave`.`pumping`,`tbl_calico_weave`.`elephant`,
`tbl_calico_weave`.`loc_id`,`tbl_calico_weave`.`app_id`,
`tbl_calico_weave`.`app_me`,`tbl_calico_weave`.`app_yr`
FROM `tbl_calico_weave`
WHERE `tbl_calico_weave`.`id` IN (1,2)
AND `tbl_calico_weave`.`tenant_id` = 1919810
AND `tbl_calico_weave`.`deleted_at` IS NULL
