--- Create(Struct)
SELECT count(*) FROM `tbl_calico_weave`
WHERE (((1!=1
OR (1=1 AND `name` = "sandbox-1"))
OR ((1=1 AND `loc_id` = 9492730750828437655) AND `app_id` = 111))
OR (((1=1 AND `app_id` = 111) AND `app_me` = "Adriel Lockman") AND `app_yr` = "Badminton"))
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave`
(`created_at`,`updated_at`,`deleted_at`,`tenant_id`,`user_id`,`name`,`desc`,`pumping`,`elephant`,`loc_id`,`app_id`,`app_me`,`app_yr`,`app_secret`)
VALUES ("2025-03-06 21:03:35.295","2025-03-06 21:03:35.295",NULL,1919810,114514,"sandbox-1","",5.5,3.3,9492730750828437655,111,"Adriel Lockman","Badminton","�Uq���ݞ���""�ZӦ")
RETURNING `id`

--- Model(Struct).Create(Map)
SELECT count(*) FROM `tbl_calico_weave`
WHERE (((1!=1
OR (1=1 AND `name` = "sandbox-2"))
OR ((1=1 AND `loc_id` = 17986825825728510162) AND `app_id` = 48))
OR (((1=1 AND `app_id` = 48) AND `app_me` = "Anabel Gaylord") AND `app_yr` = "Mathematics"))
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave`
(`app_id`,`app_me`,`app_secret`,`app_yr`,`created_at`,`elephant`,`loc_id`,`name`,`pumping`,`tenant_id`,`updated_at`,`user_id`)
VALUES (48,"Anabel Gaylord","�v�h���M���'b�","Mathematics","2025-03-06 21:03:35.296",3.3,17986825825728510162,"sandbox-2",5.5,1919810,"2025-03-06 21:03:35.296",114514)
RETURNING `id`

--- Table(table).Create(Map)
SELECT count(*) FROM `tbl_calico_weave`
WHERE (((1!=1
OR (1=1 AND `name` = "sandbox-3"))
OR ((1=1 AND `loc_id` = 15034027079695119020) AND `app_id` = 119))
OR (((1=1 AND `app_id` = 119) AND `app_me` = "Greg VonRueden") AND `app_yr` = "Movie and movie memorabilia collecting"))
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave`
(`app_id`,`app_me`,`app_secret`,`app_yr`,`created_at`,`elephant`,`loc_id`,`name`,`pumping`,`tenant_id`,`updated_at`,`user_id`)
VALUES (119,"Greg VonRueden","�T��ɥ�S�k�9,s��","Movie and movie memorabilia collecting","2025-03-06 21:03:35.297",3.3,15034027079695119020,"sandbox-3",5.5,1919810,"2025-03-06 21:03:35.297",114514)
RETURNING `id`

--- Create([]Struct)
SELECT count(*) FROM `tbl_calico_weave`
WHERE ((1!=1
OR (1!=1 OR (1=1 AND `name` = "handsome-5")))
OR (1!=1 OR (1=1 AND `name` = "handsome-6")))
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave`
(`created_at`,`updated_at`,`deleted_at`,`tenant_id`,`user_id`,`name`,`desc`,`pumping`,`elephant`,`loc_id`,`app_id`,`app_me`,`app_yr`,`app_secret`) VALUES
("2025-03-06 21:09:36.272","2025-03-06 21:09:36.272",NULL,1919810,114514,"handsome-5","",5.5,3.3,0,0,"","",""),
("2025-03-06 21:09:36.272","2025-03-06 21:09:36.272",NULL,1919810,114514,"handsome-6","",5.5,3.3,0,0,"","","")
RETURNING `id`

--- Table(table).Create([]Map)
SELECT count(*) FROM `tbl_calico_weave`
WHERE ((1!=1
OR (1!=1 OR (1=1 AND `name` = "handsome-3")))
OR (1!=1 OR (1=1 AND `name` = "handsome-4")))
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave`
(`created_at`,`elephant`,`name`,`pumping`,`tenant_id`,`updated_at`,`user_id`) VALUES
("2025-03-06 21:09:36.274",3.3,"handsome-3",5.5,1919810,"2025-03-06 21:09:36.274",114514),
("2025-03-06 21:09:36.274",3.3,"handsome-4",5.5,1919810,"2025-03-06 21:09:36.274",114514)
RETURNING `id`
