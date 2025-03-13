--- NoPlugin
INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-13 20:43:47.827","2025-03-13 20:43:47.827",NULL,"439f036d-c649-49cf-b70b-a63bf18f13d7",0,0,8504,98,281,8687,3,1,NULL,0,"Jimmy Heathcote","2877245554","9969 New Spursport, Portland, Michigan 96661",NULL,NULL,0,"",NULL),
("2025-03-13 20:43:47.827","2025-03-13 20:43:47.827",NULL,"02777fef-8b8f-4efa-963d-69287f35921b",0,0,1600,57,295,1838,3,1,NULL,0,"Jarrett Tremblay","5603887876","35106 Forgestad, Fresno, California 52253",NULL,NULL,0,"",NULL),
("2025-03-13 20:43:47.827","2025-03-13 20:43:47.827",NULL,"b2861294-df75-4380-897c-c64271ed623a",0,0,1481,57,17,1441,4,1,NULL,0,"Arlene Ratke","1286720377","916 East Cliffsborough, Newark, North Dakota 49061",NULL,NULL,0,"",NULL) RETURNING `id`

--- IsPlugin, NoUniques, NoScopes
INSERT INTO `tbl_order` (`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-13 20:43:47.829","2025-03-13 20:43:47.829",NULL,"e42b76dd-2e7c-4f5a-ae70-88777a1f2d31",0,0,5829,64,323,6088,4,1,NULL,0,"Lavern Bahringer","4660793449","971 South Brookfort, Arlington, North Dakota 39354",NULL,NULL,0,"",NULL),
("2025-03-13 20:43:47.829","2025-03-13 20:43:47.829",NULL,"c87616f5-92d7-4f0a-aed3-83b251414000",0,0,5716,39,410,6087,1,1,NULL,0,"Stephan Altenwerth","5021515731","562 Streetmouth, Washington, Wisconsin 63944",NULL,NULL,0,"",NULL),
("2025-03-13 20:43:47.829","2025-03-13 20:43:47.829",NULL,"0e46eafb-bacc-4d14-b576-b0047e83e501",0,0,2509,65,345,2789,2,1,NULL,0,"Athena Pfannerstill","1936257493","95877 Courtsborough, Durham, New Hampshire 77266",NULL,NULL,0,"",NULL) RETURNING `id`

--- IsPlugin, Uniques, 1Scopes
SELECT count(*) FROM `tbl_order`
WHERE ((1!=1
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Juston Douglas","3806472428","231 West Ovalport, Winston-Salem, Maine 68617"),("Lindsay Wuckert","3999141090","603 Isleview, Seattle, Rhode Island 42465"),("Ava Sanford","6002674958","9300 Divideland, Hialeah, Massachusetts 13620")))
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((7564,70,95),(2137,61,170),(4592,69,450)))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-13 20:57:17.332","2025-03-13 20:57:17.332",NULL,"94b7fc57-4d2b-4600-baf6-fa79bab750d1",666,0,7564,70,95,7589,3,1,NULL,0,"Juston Douglas","3806472428","231 West Ovalport, Winston-Salem, Maine 68617",NULL,NULL,0,"",NULL),
("2025-03-13 20:57:17.332","2025-03-13 20:57:17.332",NULL,"4d2cb0b9-9ec0-41bf-977c-42ecb1d588e6",666,0,2137,61,170,2246,3,1,NULL,0,"Lindsay Wuckert","3999141090","603 Isleview, Seattle, Rhode Island 42465",NULL,NULL,0,"",NULL),
("2025-03-13 20:57:17.332","2025-03-13 20:57:17.332",NULL,"767e6dbb-e1c2-41fa-8b2e-38288864e5aa",666,0,4592,69,450,4973,2,1,NULL,0,"Ava Sanford","6002674958","9300 Divideland, Hialeah, Massachusetts 13620",NULL,NULL,0,"",NULL) RETURNING `id`

--- IsPlugin, Uniques, 2Scopes
SELECT count(*) FROM `tbl_order`
WHERE ((1!=1
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((4070,28,415),(7368,30,41),(2452,59,479)))
OR (`receiver_name`,`receiver_phone`,`receiver_address`) IN (("Mollie Ruecker","1739529553","9947 East Spurville, Scottsdale, Hawaii 91495"),("Cortney Murphy","5126748204","1815 Parksview, Aurora, North Dakota 64275"),("Cynthia Nienow","1512986349","24152 South Clubfort, Fort Worth, Mississippi 68119")))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order`
(`created_at`,`updated_at`,`deleted_at`,`serial`,`user_id`,`tenant_id`,`amount_total`,`amount_discount`,`shipping_fee`,`amount_actual`,`order_type`,`order_status`,`pay_time`,`pay_method`,`receiver_name`,`receiver_phone`,`receiver_address`,`shipped_time`,`finished_time`,`logistics_id`,`logistics_name`,`extras_info`) VALUES
("2025-03-13 20:57:17.334","2025-03-13 20:57:17.334",NULL,"f63efe41-a0b1-4d7f-b03d-2dcc64ba5430",666,888,4070,28,415,4457,1,1,NULL,0,"Mollie Ruecker","1739529553","9947 East Spurville, Scottsdale, Hawaii 91495",NULL,NULL,0,"",NULL),
("2025-03-13 20:57:17.334","2025-03-13 20:57:17.334",NULL,"c38b9feb-26e7-4a86-813b-8f76850e4caa",666,888,7368,30,41,7379,2,1,NULL,0,"Cortney Murphy","5126748204","1815 Parksview, Aurora, North Dakota 64275",NULL,NULL,0,"",NULL),
("2025-03-13 20:57:17.334","2025-03-13 20:57:17.334",NULL,"e882bbe3-446b-45a7-8628-e66c82cf439e",666,888,2452,59,479,2872,4,1,NULL,0,"Cynthia Nienow","1512986349","24152 South Clubfort, Fort Worth, Mississippi 68119",NULL,NULL,0,"",NULL) RETURNING `id`
