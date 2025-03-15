--- NoPlugin
UPDATE `tbl_order` SET `deleted_at`="2025-03-15 10:01:11.899"
WHERE (`tbl_order`.`id`,`tbl_order`.`serial`)
IN ((149,"4ca872a3-9449-479d-8582-c8217c49bb7e"),
(150,"01e5e5d4-5671-4470-a2e7-537aef80c6fd"),
(151,"28f274e8-893d-4b8c-9dde-a14511e2327c"))
AND `tbl_order`.`deleted_at` IS NULL

--- IsPlugin NoScopes
UPDATE `tbl_order` SET `deleted_at`="2025-03-15 10:01:11.901"
WHERE (`tbl_order`.`id`,`tbl_order`.`serial`)
IN ((152,"17a63dad-b9c8-4780-88fe-fb885f7a8a7e"),
(153,"16c307b1-a04e-4d32-bb28-e780c66a7be5"),
(154,"761d526a-d8cf-4bcf-a3bc-235bc8a29da9"))
AND `tbl_order`.`deleted_at` IS NULL

--- IsPlugin Scopes(user_id)
UPDATE `tbl_order` SET `deleted_at`="2025-03-15 10:01:11.903"
WHERE `tbl_order`.`user_id` = 666
AND (`tbl_order`.`id`,`tbl_order`.`serial`)
IN ((155,"d7f08ea3-c20e-4374-ad2b-d6cbabf0afe0"),
(156,"5a0c3930-ddf0-4c41-b071-692f34480f02"),
(157,"26f0c4dd-eef7-41fd-a9c4-2dc41e8aba41"))
AND `tbl_order`.`deleted_at` IS NULL

--- IsPlugin Scopes(user_id, tenant_id)
UPDATE `tbl_order` SET `deleted_at`="2025-03-15 10:01:11.904"
WHERE `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666
AND (`tbl_order`.`id`,`tbl_order`.`serial`)
IN ((158,"e544a69f-6d4d-481f-9f0a-981ca6893f6e"),
(159,"87eae986-1dea-4980-809d-d25c51e6734c"),
(160,"f36b85df-6265-4f0e-8c42-aa9e497b1b51"))
AND `tbl_order`.`deleted_at` IS NULL
