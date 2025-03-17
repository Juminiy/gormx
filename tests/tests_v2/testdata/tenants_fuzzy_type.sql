SELECT count(*) FROM `tbl_product`
WHERE ((((1!=1
OR (`name`,`ean`,`upc`) IN (("Rubber Silver Fridge","3944 - Belgian Witbier","060639777970")))
OR (`p_category`,`p_brand`,`p_model`) IN ((10,"Thriller","silicon")))
OR (`ean`) IN (("3944 - Belgian Witbier")))
OR (`upc`) IN (("060639777970")))
AND `tbl_product`.`deleted_at` IS NULL
AND `tbl_product`.`merchant_id` = 888

INSERT INTO `tbl_product` (`created_at`,`updated_at`,`deleted_at`,`name`,`desc`,`price`,`total`,`status`,`p_category`,`p_brand`,`p_model`,`ean`,`upc`,`merchant_id`) VALUES
("2025-03-17 21:54:33.928","2025-03-17 21:54:33.928",NULL,"Rubber Silver Fridge","Enjoy the luxury of noise reduction with this product, crafted from paper. Its horrible design and water-resistant make it ideal for gardening.",4378,5884,0,10,"Thriller","silicon","3944 - Belgian Witbier","060639777970","888") RETURNING `id`

SELECT `tbl_product`.`id`,`tbl_product`.`created_at`,`tbl_product`.`updated_at`,`tbl_product`.`deleted_at`,`tbl_product`.`name`,`tbl_product`.`desc`,`tbl_product`.`price`,`tbl_product`.`total`,`tbl_product`.`status`,`tbl_product`.`p_category`,`tbl_product`.`p_brand`,`tbl_product`.`p_model`,`tbl_product`.`ean`,`tbl_product`.`merchant_id`
FROM `tbl_product`
WHERE `tbl_product`.`merchant_id` = 888
AND `tbl_product`.`deleted_at` IS NULL
AND `tbl_product`.`id` = 9
ORDER BY `tbl_product`.`id` LIMIT 1

SELECT count(*) FROM `tbl_product`
WHERE (1!=1
OR (`ean`) IN (("3942 - Belgian Wheat")))
AND `tbl_product`.`deleted_at` IS NULL
AND NOT `tbl_product`.`id` = 9
AND `tbl_product`.`merchant_id` = 888

UPDATE `tbl_product`
SET `ean`="3942 - Belgian Wheat",`p_brand`="Saga",`p_category`=7,`p_model`="",`updated_at`="2025-03-17 21:54:33.93"
WHERE `tbl_product`.`id` = 9
AND `tbl_product`.`merchant_id` = 888
AND `tbl_product`.`deleted_at` IS NULL

UPDATE `tbl_product`
SET `deleted_at`="2025-03-17 21:54:33.931"
WHERE `tbl_product`.`merchant_id` = 888
AND `tbl_product`.`id` = 9
AND `tbl_product`.`deleted_at` IS NULL RETURNING *
