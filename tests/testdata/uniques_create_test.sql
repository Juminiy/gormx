--- Struct
---no scope
SELECT count(*) FROM `tbl_baby_trade`
WHERE (1!=1 OR (1=1 AND `sim_uuid` = "1cd4f9a5-a576-4573-97f0-0d25a0c90387"))
AND `tbl_baby_trade`.`deleted_at` = 0

INSERT INTO `tbl_baby_trade`
(`create_time`,`update_time`,`deleted_at`,`sim_uuid`,`user_id`,`tenant_id`,`auction_id`,`cat_id`,`cat`,`buy_mount`,`day`) VALUES
("2025-03-09 13:39:23.628","2025-03-09 13:39:23.628",0,"1cd4f9a5-a576-4573-97f0-0d25a0c90387",0,0,934,0,0,28011,"1939-03-16 20:01:46.586876096 +0000 UTC")
RETURNING `id`

---one scope
SELECT count(*) FROM `tbl_baby_trade`
WHERE (((1!=1
OR (1=1 AND `sim_uuid` = "fef98892-22c8-40ed-ba9e-cb36b053c5de"))
OR ((1=1 AND `auction_id` = 754) AND `cat_id` = 10))
OR ((1=1 AND `cat_id` = 10) AND `cat` = 1))
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`user_id` = 114514

INSERT INTO `tbl_baby_trade`
(`create_time`,`update_time`,`deleted_at`,`sim_uuid`,`user_id`,`tenant_id`,`auction_id`,`cat_id`,`cat`,`buy_mount`,`day`) VALUES
("2025-03-09 13:39:23.629","2025-03-09 13:39:23.629",0,"fef98892-22c8-40ed-ba9e-cb36b053c5de",114514,0,754,10,1,46772,"1932-11-16 23:51:02.327133761 +0000 UTC")
RETURNING `id`

---two scope
SELECT count(*) FROM `tbl_baby_trade`
WHERE (((1!=1
OR ((1=1 AND `cat_id` = 8) AND `cat` = 1))
OR (1=1 AND `sim_uuid` = "59ee3b8e-1da6-43cf-ba19-451261046149"))
OR ((1=1 AND `auction_id` = 646) AND `cat_id` = 8))
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514

INSERT INTO `tbl_baby_trade`
(`create_time`,`update_time`,`deleted_at`,`sim_uuid`,`user_id`,`tenant_id`,`auction_id`,`cat_id`,`cat`,`buy_mount`,`day`) VALUES
("2025-03-09 13:39:23.63","2025-03-09 13:39:23.63",0,"59ee3b8e-1da6-43cf-ba19-451261046149",114514,114514,646,8,1,33056,"1912-10-17 11:37:45.553369649 +0000 UTC")
RETURNING `id`



--- Map
---no scope
SELECT count(*) FROM `tbl_baby_trade`
WHERE (((1!=1
OR (1=1 AND `sim_uuid` = "cd84cfe7-2a7a-4ab5-ab53-bd6ff7746b2a"))
OR ((1=1 AND `auction_id` = 1021) AND `cat_id` = 425))
OR ((1=1 AND `cat_id` = 425) AND `cat` = 72))
AND `tbl_baby_trade`.`deleted_at` = 0

INSERT INTO `tbl_baby_trade`
(`auction_id`,`buy_mount`,`cat`,`cat_id`,`create_time`,`day`,`sim_uuid`,`update_time`) VALUES
(1021,122,72,425,"2025-03-09 14:08:04.554","1930-02-01 13:40:52.77165735 +0000 UTC","cd84cfe7-2a7a-4ab5-ab53-bd6ff7746b2a","2025-03-09 14:08:04.554")
RETURNING `id`


---one scope
SELECT count(*) FROM `tbl_baby_trade`
WHERE (((1!=1
OR (1=1 AND `sim_uuid` = "4f010f47-5077-4182-911f-39fbceb33b23"))
OR ((1=1 AND `auction_id` = 64) AND `cat_id` = 68))
OR ((1=1 AND `cat_id` = 68) AND `cat` = 56))
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`user_id` = 114514

INSERT INTO `tbl_baby_trade`
(`auction_id`,`buy_mount`,`cat`,`cat_id`,`create_time`,`day`,`sim_uuid`,`update_time`,`user_id`) VALUES
(64,128,56,68,"2025-03-09 14:08:04.556","2018-05-26 12:14:20.083766944 +0000 UTC","4f010f47-5077-4182-911f-39fbceb33b23","2025-03-09 14:08:04.556",114514)
RETURNING `id`


---two scopes
SELECT count(*) FROM `tbl_baby_trade`
WHERE (((1!=1
OR (1=1 AND `sim_uuid` = "2c61651d-d688-4607-a184-b772a72b0818"))
OR ((1=1 AND `auction_id` = 911) AND `cat_id` = 361))
OR ((1=1 AND `cat_id` = 361) AND `cat` = 196))
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514

INSERT INTO `tbl_baby_trade`
(`auction_id`,`buy_mount`,`cat`,`cat_id`,`create_time`,`day`,`sim_uuid`,`tenant_id`,`update_time`,`user_id`) VALUES
(911,1613,196,361,"2025-03-09 14:08:04.557","1961-08-22 10:31:27.732091288 +0000 UTC","2c61651d-d688-4607-a184-b772a72b0818",114514,"2025-03-09 14:08:04.557",114514)
RETURNING `id`



---StructList
SELECT count(*) FROM `tbl_baby_trade`
WHERE ((1!=1 OR (((1!=1
OR ((1=1 AND `auction_id` = 548) AND `cat_id` = 15))
OR ((1=1 AND `cat_id` = 15) AND `cat` = 1))
OR (1=1 AND `sim_uuid` = "136a9664-803f-44af-8f31-154e9d3a75fd")))
OR ((1!=1 OR ((1=1 AND `auction_id` = 179) AND `cat_id` = 1))
OR (1=1 AND `sim_uuid` = "2d399e71-a72a-435a-b1b1-a9069d395e37")))
AND `tbl_baby_trade`.`deleted_at` = 0

INSERT INTO `tbl_baby_trade`
(`create_time`,`update_time`,`deleted_at`,`sim_uuid`,`user_id`,`tenant_id`,`auction_id`,`cat_id`,`cat`,`buy_mount`,`day`) VALUES
("2025-03-09 14:35:50.697","2025-03-09 14:35:50.697",0,"136a9664-803f-44af-8f31-154e9d3a75fd",0,0,548,15,1,64980,"1902-12-02 00:16:47.231538879 +0000 UTC"),
("2025-03-09 14:35:50.697","2025-03-09 14:35:50.697",0,"2d399e71-a72a-435a-b1b1-a9069d395e37",0,0,179,1,0,64612,"1982-08-05 02:43:37.64237476 +0000 UTC")
RETURNING `id`

---MapList
SELECT count(*) FROM `tbl_baby_trade`
WHERE ((1!=1 OR (((1!=1
OR ((1=1 AND `cat_id` = 305) AND `cat` = 76))
OR (1=1 AND `sim_uuid` = "88334ee4-67fd-492e-8551-8af907ce948d"))
OR ((1=1 AND `auction_id` = 481) AND `cat_id` = 305)))
OR (((1!=1 OR (1=1 AND `sim_uuid` = "7cbba8fb-66e5-435d-988c-2d486cdf713e"))
OR ((1=1 AND `auction_id` = 410) AND `cat_id` = 60))
OR ((1=1 AND `cat_id` = 60) AND `cat` = 11)))
AND `tbl_baby_trade`.`deleted_at` = 0

INSERT INTO `tbl_baby_trade`
(`auction_id`,`buy_mount`,`cat`,`cat_id`,`create_time`,`day`,`sim_uuid`,`update_time`) VALUES
(481,1286,76,305,"2025-03-09 14:39:58.031","1971-08-25 23:43:48.601870685 +0000 UTC","88334ee4-67fd-492e-8551-8af907ce948d","2025-03-09 14:39:58.031"),
(410,3612,11,60,"2025-03-09 14:39:58.031","1935-01-03 08:39:38.613988717 +0000 UTC","7cbba8fb-66e5-435d-988c-2d486cdf713e","2025-03-09 14:39:58.031")
RETURNING `id`



---MapList, two scope, field dup
SELECT count(*) FROM `tbl_baby_trade`
WHERE ((1!=1 OR (((1!=1
OR (1=1 AND `sim_uuid` = "1e53bf2e-24ca-4d6e-aecf-c50ff88f1962"))
OR ((1=1 AND `auction_id` = 54) AND `cat_id` = 366))
OR ((1=1 AND `cat_id` = 366) AND `cat` = 162)))
OR (((1!=1
OR (1=1 AND `sim_uuid` = "bb2c38b0-ef6e-4af8-96e5-a8f60e7ee1b9"))
OR ((1=1 AND `auction_id` = 132) AND `cat_id` = 135))
OR ((1=1 AND `cat_id` = 135) AND `cat` = 40)))
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514

INSERT INTO `tbl_baby_trade`
(`auction_id`,`buy_mount`,`cat`,`cat_id`,`create_time`,`day`,`sim_uuid`,`tenant_id`,`update_time`,`user_id`) VALUES
(54,2718,162,366,"2025-03-09 16:11:49.884","1902-06-13 22:38:40.367356453 +0000 UTC","1e53bf2e-24ca-4d6e-aecf-c50ff88f1962",114514,"2025-03-09 16:11:49.884",114514),
(132,1855,40,135,"2025-03-09 16:11:49.884","1968-03-19 22:02:51.226301578 +0000 UTC","bb2c38b0-ef6e-4af8-96e5-a8f60e7ee1b9",114514,"2025-03-09 16:11:49.884",114514)
RETURNING `id`
---

SELECT count(*) FROM `tbl_product` WHERE
(((((1!=1 OR ((1!=1
OR ((1=1 AND `name` = "Milk") AND `desc` = "Fresh milk"))
OR (1=1 AND `code` = 100001)))
OR ((1!=1 OR ((1=1 AND `name` = "Bread") AND `desc` = "Whole wheat bread"))
OR (1=1 AND `code` = 100002)))
OR ((1!=1 OR ((1=1 AND `name` = "Rice") AND `desc` = "Long grain rice"))
OR (1=1 AND `code` = 100003)))
OR ((1!=1 OR (1=1 AND `code` = 100004))
OR ((1=1 AND `name` = "Eggs") AND `desc` = "Free-range eggs")))
OR ((1!=1 OR ((1=1 AND `name` = "Chicken") AND `desc` = "Fresh chicken breast"))
OR (1=1 AND `code` = 100006)))
AND `tbl_product`.`deleted_at` IS NULL
AND `tbl_product`.`tenant_id` = 114514
