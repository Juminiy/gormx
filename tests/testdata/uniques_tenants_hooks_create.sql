SELECT count(*) FROM `tbl_calico_weave`
WHERE (((1!=1
OR ((1=1 AND `loc_id` = 4362667370422918555) AND `app_id` = 100))
OR (((1=1 AND `app_id` = 100) AND `app_me` = "Donavon Considine") AND `app_yr` = "Sea glass collecting"))
OR (1=1 AND `name` = "sandbox-1"))
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave`
(`created_at`,`updated_at`,`deleted_at`,`tenant_id`,`user_id`,`name`,`desc`,`pumping`,`elephant`,`loc_id`,`app_id`,`app_me`,`app_yr`,`app_secret`)
VALUES ("2025-03-06 21:28:54.877","2025-03-06 21:28:54.877",NULL,1919810,114514,"sandbox-1","",5.5,3.3,4362667370422918555,100,"Donavon Considine","Sea glass collecting","�Uq���ݞ���""�ZӦ") RETURNING `id`

SELECT count(*) FROM `tbl_calico_weave`
WHERE (((1!=1
OR (1=1 AND `name` = "sandbox-2"))
OR ((1=1 AND `loc_id` = 2831154127680706952) AND `app_id` = 44))
OR (((1=1 AND `app_id` = 44) AND `app_me` = "Maud Quitzon") AND `app_yr` = "Flower collecting and pressing"))
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave` (`app_id`,`app_me`,`app_secret`,`app_yr`,`created_at`,`elephant`,`loc_id`,`name`,`pumping`,`tenant_id`,`updated_at`,`user_id`)
VALUES
(44,"Maud Quitzon","�v�h���M���'b�","Flower collecting and pressing","2025-03-06 21:28:54.878",3.3,2831154127680706952,"sandbox-2",5.5,1919810,"2025-03-06 21:28:54.878",114514)
RETURNING `id`

SELECT count(*) FROM `tbl_calico_weave`
WHERE (((1!=1
OR (1=1 AND `name` = "sandbox-3"))
OR ((1=1 AND `loc_id` = 9801296249076260640) AND `app_id` = 116))
OR (((1=1 AND `app_id` = 116) AND `app_me` = "Hailey Mosciski") AND `app_yr` = "Hooping"))
AND `tbl_calico_weave`.`deleted_at` IS NULL

INSERT INTO `tbl_calico_weave` (`app_id`,`app_me`,`app_secret`,`app_yr`,`created_at`,`elephant`,`loc_id`,`name`,`pumping`,`tenant_id`,`updated_at`,`user_id`)
VALUES (116,"Hailey Mosciski","�T��ɥ�S�k�9,s��","Hooping","2025-03-06 21:28:54.879",3.3,9801296249076260640,"sandbox-3",5.5,1919810,"2025-03-06 21:28:54.879",114514)
RETURNING `id`

