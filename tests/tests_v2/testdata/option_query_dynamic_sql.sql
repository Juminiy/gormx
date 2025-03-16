--- NoPlugin
--- near "?": syntax error
SELECT * FROM `tbl_order`
WHERE (id >= 6 AND id <= 11)
AND id IN "1f8c0ae0-b0e2-4a8e-a17a-5bb84f5a8bc2"
AND pay_method IS NOT NULL
AND id = "de6b5b38-8084-4095-81b2-406b35962ac1"
AND serial = "9e61a00b-6301-410c-adb9-b52bb08bdd8e"
AND shipped_time IS NULL
AND order_type = 1
AND `tbl_order`.`deleted_at` IS NULL
ORDER BY created_at DESC,amount_total ASC
LIMIT 10 OFFSET 3

--- IsPlugin, NoOption
--- near "?": syntax error
SELECT * FROM `tbl_order`
WHERE (id >= 14 AND id <= 1)
AND id IN "619557e6-a9dc-40ff-a722-fe45dfe986e7"
AND pay_method IS NOT NULL
AND id = "bc85184a-1999-400a-98f8-b51f2424efce"
AND serial = "3089e930-5868-44af-bb88-c67641bda914"
AND shipped_time IS NULL
AND order_type = 2
AND `tbl_order`.`deleted_at` IS NULL
ORDER BY created_at DESC,amount_total ASC
LIMIT 10 OFFSET 3

---IsPlugin, Option
SELECT * FROM `tbl_order`
WHERE pay_method IS NOT NULL
AND shipped_time IS NULL
AND order_type = 1
AND `tbl_order`.`deleted_at` IS NULL
ORDER BY created_at DESC,amount_total ASC
LIMIT 10 OFFSET 3
