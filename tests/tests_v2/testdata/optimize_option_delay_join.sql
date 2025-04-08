[5.263ms] [rows:0]
SELECT * FROM `tbl_order`
WHERE (amount_total BETWEEN 10 AND 5000000)
AND `tbl_order`.`deleted_at` IS NULL
ORDER BY shipping_fee desc
LIMIT 10 OFFSET 500000
--- PASS: TestOptimizeSQLDelayJoin/NoPlugin (0.00s)

[0.933ms] [rows:0]
SELECT `tbl_order`.`id`,`tbl_order`.`created_at`,
`tbl_order`.`updated_at`,`tbl_order`.`deleted_at`,
`tbl_order`.`serial`,`tbl_order`.`user_id`,
`tbl_order`.`tenant_id`,`tbl_order`.`amount_total`,
`tbl_order`.`amount_discount`,`tbl_order`.`shipping_fee`,
`tbl_order`.`amount_actual`,`tbl_order`.`order_type`,
`tbl_order`.`order_status`,`tbl_order`.`pay_time`,
`tbl_order`.`pay_method`,`tbl_order`.`receiver_name`,
`tbl_order`.`receiver_phone`,`tbl_order`.`receiver_address`,
`tbl_order`.`shipped_time`,`tbl_order`.`finished_time`,
`tbl_order`.`logistics_id`,`tbl_order`.`logistics_name`,
`tbl_order`.`extras_info`
FROM `tbl_order` INNER JOIN (
    SELECT `id`,`serial` FROM `tbl_order`
    WHERE (amount_total BETWEEN 10 AND 5000000)
    AND `tbl_order`.`deleted_at` IS NULL
    ORDER BY shipping_fee desc
    LIMIT 10 OFFSET 500000
) AS tbl_order_pk
USING (`id`,`serial`)
WHERE `tbl_order`.`deleted_at` IS NULL
--- PASS: TestOptimizeSQLDelayJoin/IsPlugin (0.00s)