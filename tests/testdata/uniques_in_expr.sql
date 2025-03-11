SELECT count(*) FROM `tbl_baby_trade`
WHERE (((1!=1
OR (`sim_uuid`) IN (("aa12e128-9cd7-4546-a59a-3723ba202955")))
OR (`auction_id`,`cat_id`) IN ((328,3)))
OR (`cat_id`,`cat`) IN ((3,0)))
AND `tbl_baby_trade`.`deleted_at` = 0
AND `tbl_baby_trade`.`tenant_id` = 114514
AND `tbl_baby_trade`.`user_id` = 114514

SELECT count(*) FROM `tbl_product`
WHERE ((1!=1
OR (`name`,`desc`) IN (("Milk","Fresh milk"),("Bread","Whole wheat bread"),("Rice","Long grain rice"),("Eggs","Free-range eggs"),("Chicken","Fresh chicken breast")))
OR (`code`) IN ((100001),(100002),(100003),(100004),(100006)))
AND `tbl_product`.`deleted_at` IS NULL
AND `tbl_product`.`tenant_id` = 114514

SELECT count(*) FROM `tbl_product`
WHERE ((1!=1
OR (`name`,`desc`) IN (("Beer","Local lager beer"),("Noodles","Instant noodles"),("Shampoo","Herbal shampoo"),("Toothpaste","Mint toothpaste")))
OR (`code`) IN ((100007),(100008),(100009),(100010)))
AND `tbl_product`.`deleted_at` IS NULL
AND `tbl_product`.`tenant_id` = 114514
