=== RUN   TestCreateMap/CreateMap_NoPlugin
INSERT INTO `tbl_order`
(`amount_discount`,`amount_total`,`order_type`,`shipping_fee`) VALUES
(73,6080,4,399) RETURNING `id`
    api_create_test.go:111: {
          "AmountDiscount": 73,
          "AmountTotal": 6080,
          "OrderType": "虚拟商品订单",
          "ShippingFee": 399,
          "id": 43
        }
--- PASS: TestCreateMap/CreateMap_NoPlugin (0.00s)

=== RUN   TestCreateMap/CreateMap_IsPlugin:_Uniques,_NoScopes
SELECT count(*) FROM `tbl_order`
WHERE (1!=1 OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((9970,30,226)))
AND `tbl_order`.`deleted_at` IS NULL

INSERT INTO `tbl_order` (`amount_discount`,`amount_total`,`created_at`,`order_type`,`shipping_fee`,`updated_at`) VALUES (30,9970,"2025-03-13 22:13:26.89",1,226,"2025-03-13 22:13:26.89") RETURNING `id`
    api_create_test.go:116: {
          "AmountDiscount": 30,
          "AmountTotal": 9970,
          "CreatedAt": "2025-03-13T22:13:26.890135+08:00",
          "ID": 44,
          "OrderType": "普通订单",
          "ShippingFee": 226,
          "UpdatedAt": "2025-03-13T22:13:26.890135+08:00"
        }
--- PASS: TestCreateMap/CreateMap_IsPlugin:_Uniques,_NoScopes (0.00s)

=== RUN   TestCreateMap/CreateMap_IsPlugin:_Uniques,_1Scopes(user_id)
SELECT count(*) FROM `tbl_order`
WHERE (1!=1 OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((6620,100,141)))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order`
(`amount_discount`,`amount_total`,`created_at`,`order_type`,`shipping_fee`,`updated_at`,`user_id`)
VALUES (100,6620,"2025-03-13 22:13:26.89",1,141,"2025-03-13 22:13:26.89",666) RETURNING `id`
    api_create_test.go:122: {
          "AmountDiscount": 100,
          "AmountTotal": 6620,
          "CreatedAt": "2025-03-13T22:13:26.890915+08:00",
          "ID": 45,
          "OrderType": "普通订单",
          "ShippingFee": 141,
          "UpdatedAt": "2025-03-13T22:13:26.890915+08:00"
        }
--- PASS: TestCreateMap/CreateMap_IsPlugin:_Uniques,_1Scopes(user_id) (0.00s)

=== RUN   TestCreateMap/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id)
SELECT count(*) FROM `tbl_order`
WHERE (1!=1 OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((8383,8,84)))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order` (`amount_discount`,`amount_total`,`created_at`,`order_type`,`shipping_fee`,`tenant_id`,`updated_at`,`user_id`) VALUES
(8,8383,"2025-03-13 22:13:26.891",3,84,888,"2025-03-13 22:13:26.891",666) RETURNING `id`
    api_create_test.go:128: {
          "AmountDiscount": 8,
          "AmountTotal": 8383,
          "CreatedAt": "2025-03-13T22:13:26.891483+08:00",
          "ID": 46,
          "OrderType": "秒杀订单",
          "ShippingFee": 84,
          "TenantID": 888,
          "UpdatedAt": "2025-03-13T22:13:26.891483+08:00",
          "UserID": 666
        }
--- PASS: TestCreateMap/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id) (0.00s)

=== RUN   TestCreateMap/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id);_Reinforced:_before_call_hooks
SELECT count(*) FROM `tbl_order` WHERE (1!=1 OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((8253,44,77))) AND `tbl_order`.`deleted_at` IS NULL AND `tbl_order`.`tenant_id` = 888 AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order` (`amount_actual`,`amount_discount`,`amount_total`,`created_at`,`order_status`,`order_type`,`serial`,`shipping_fee`,`tenant_id`,`updated_at`,`user_id`) VALUES (8286,44,8253,"2025-03-13 22:13:26.892",1,2,"3c0a9f63-0d9e-4803-a0ef-d09eb987b970",77,888,"2025-03-13 22:13:26.892",666) RETURNING `id`
    api_create_test.go:134: {
          "AmountActual": 8286,
          "AmountDiscount": 44,
          "AmountTotal": 8253,
          "CreatedAt": "2025-03-13T22:13:26.892057+08:00",
          "ID": 47,
          "OrderStatus": "待支付",
          "OrderType": "团购订单",
          "Serial": "3c0a9f63-0d9e-4803-a0ef-d09eb987b970",
          "ShippingFee": 77,
          "TenantID": 888,
          "UpdatedAt": "2025-03-13T22:13:26.892057+08:00",
          "UserID": 666
        }
--- PASS: TestCreateMap/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id);_Reinforced:_before_call_hooks (0.00s)

=== RUN   TestCreateMap/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id);_Reinforced:_before_and_after_call_hooks
SELECT count(*) FROM `tbl_order` WHERE (1!=1 OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((8067,68,82))) AND `tbl_order`.`deleted_at` IS NULL AND `tbl_order`.`tenant_id` = 888 AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order` (`amount_actual`,`amount_discount`,`amount_total`,`created_at`,`order_status`,`order_type`,`serial`,`shipping_fee`,`tenant_id`,`updated_at`,`user_id`) VALUES (8081,68,8067,"2025-03-13 22:13:26.892",1,1,"a19e711e-5e36-43a3-acb2-09366f76ae1e",82,888,"2025-03-13 22:13:26.892",666) RETURNING `id`
    api_create_test.go:140: {
          "AmountActual": 8081,
          "AmountDiscount": 68,
          "AmountTotal": 8067,
          "CreatedAt": "2025-03-13T22:13:26.892583+08:00",
          "ID": 48,
          "OrderStatus": "待支付",
          "OrderType": "普通订单",
          "Serial": "a19e711e-5e36-43a3-acb2-09366f76ae1e",
          "ShippingFee": 82,
          "TenantID": 888,
          "UpdatedAt": "2025-03-13T22:13:26.892583+08:00",
          "UserID": 666
        }
--- PASS: TestCreateMap/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id);_Reinforced:_before_and_after_call_hooks (0.00s)