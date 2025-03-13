=== RUN   TestCreateMapList/CreateMap_NoPlugin
INSERT INTO `tbl_order` (`amount_discount`,`amount_total`,`order_type`,`shipping_fee`) VALUES (73,8128,1,70),(10,1010,4,477),(37,1507,2,173) RETURNING `id`
    api_create_test.go:156: [
          {
            "AmountDiscount": 73,
            "AmountTotal": 8128,
            "OrderType": "普通订单",
            "ShippingFee": 70
          },
          {
            "AmountDiscount": 10,
            "AmountTotal": 1010,
            "OrderType": "虚拟商品订单",
            "ShippingFee": 477
          },
          {
            "AmountDiscount": 37,
            "AmountTotal": 1507,
            "OrderType": "团购订单",
            "ShippingFee": 173
          },
          {
            "id": 79
          },
          {
            "id": 80
          },
          {
            "id": 81
          }
        ]
--- PASS: TestCreateMapList/CreateMap_NoPlugin (0.00s)

=== RUN   TestCreateMapList/CreateMap_IsPlugin:_Uniques,_NoScopes
INSERT INTO `tbl_order` (`amount_discount`,`amount_total`,`created_at`,`order_type`,`shipping_fee`,`updated_at`) VALUES
(90,8759,"2025-03-13 22:42:34.382",1,256,"2025-03-13 22:42:34.382"),
(39,8675,"2025-03-13 22:42:34.382",2,283,"2025-03-13 22:42:34.382"),
(66,8901,"2025-03-13 22:42:34.382",4,170,"2025-03-13 22:42:34.382")
RETURNING `id`
    api_create_test.go:161: [
          {
            "AmountDiscount": 90,
            "AmountTotal": 8759,
            "CreatedAt": "2025-03-13T22:42:34.382304+08:00",
            "ID": 82,
            "OrderType": "普通订单",
            "ShippingFee": 256,
            "UpdatedAt": "2025-03-13T22:42:34.382304+08:00"
          },
          {
            "AmountDiscount": 39,
            "AmountTotal": 8675,
            "CreatedAt": "2025-03-13T22:42:34.382304+08:00",
            "ID": 83,
            "OrderType": "团购订单",
            "ShippingFee": 283,
            "UpdatedAt": "2025-03-13T22:42:34.382304+08:00"
          },
          {
            "AmountDiscount": 66,
            "AmountTotal": 8901,
            "CreatedAt": "2025-03-13T22:42:34.382304+08:00",
            "ID": 84,
            "OrderType": "虚拟商品订单",
            "ShippingFee": 170,
            "UpdatedAt": "2025-03-13T22:42:34.382304+08:00"
          }
        ]
--- PASS: TestCreateMapList/CreateMap_IsPlugin:_Uniques,_NoScopes (0.00s)

=== RUN   TestCreateMapList/CreateMap_IsPlugin:_Uniques,_1Scopes(user_id)
SELECT count(*) FROM `tbl_order`
WHERE (1!=1
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((8195,60,390),(4628,14,72),(4381,37,153)))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order` (`amount_discount`,`amount_total`,`created_at`,`order_type`,`shipping_fee`,`updated_at`,`user_id`) VALUES
(60,8195,"2025-03-13 22:42:34.383",1,390,"2025-03-13 22:42:34.383",666),
(14,4628,"2025-03-13 22:42:34.383",2,72,"2025-03-13 22:42:34.383",666),
(37,4381,"2025-03-13 22:42:34.383",4,153,"2025-03-13 22:42:34.383",666) RETURNING `id`
    api_create_test.go:172: [
          {
            "AmountDiscount": 60,
            "AmountTotal": 8195,
            "CreatedAt": "2025-03-13T22:42:34.383062+08:00",
            "ID": 85,
            "OrderType": "普通订单",
            "ShippingFee": 390,
            "UpdatedAt": "2025-03-13T22:42:34.383062+08:00"
          },
          {
            "AmountDiscount": 14,
            "AmountTotal": 4628,
            "CreatedAt": "2025-03-13T22:42:34.383062+08:00",
            "ID": 86,
            "OrderType": "团购订单",
            "ShippingFee": 72,
            "UpdatedAt": "2025-03-13T22:42:34.383062+08:00"
          },
          {
            "AmountDiscount": 37,
            "AmountTotal": 4381,
            "CreatedAt": "2025-03-13T22:42:34.383062+08:00",
            "ID": 87,
            "OrderType": "虚拟商品订单",
            "ShippingFee": 153,
            "UpdatedAt": "2025-03-13T22:42:34.383062+08:00"
          }
        ]
--- PASS: TestCreateMapList/CreateMap_IsPlugin:_Uniques,_1Scopes(user_id) (0.00s)

=== RUN   TestCreateMapList/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id)
SELECT count(*) FROM `tbl_order` WHERE (1!=1
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((7372,58,338),(1385,88,45),(5488,93,83)))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order` (`amount_discount`,`amount_total`,`created_at`,`order_type`,`shipping_fee`,`tenant_id`,`updated_at`,`user_id`) VALUES
(58,7372,"2025-03-13 22:42:34.383",3,338,888,"2025-03-13 22:42:34.383",666),
(88,1385,"2025-03-13 22:42:34.383",2,45,888,"2025-03-13 22:42:34.383",666),
(93,5488,"2025-03-13 22:42:34.383",4,83,888,"2025-03-13 22:42:34.383",666) RETURNING `id`
    api_create_test.go:183: [
          {
            "AmountDiscount": 58,
            "AmountTotal": 7372,
            "CreatedAt": "2025-03-13T22:42:34.383655+08:00",
            "ID": 88,
            "OrderType": "秒杀订单",
            "ShippingFee": 338,
            "TenantID": 888,
            "UpdatedAt": "2025-03-13T22:42:34.383655+08:00",
            "UserID": 666
          },
          {
            "AmountDiscount": 88,
            "AmountTotal": 1385,
            "CreatedAt": "2025-03-13T22:42:34.383655+08:00",
            "ID": 89,
            "OrderType": "团购订单",
            "ShippingFee": 45,
            "TenantID": 888,
            "UpdatedAt": "2025-03-13T22:42:34.383655+08:00",
            "UserID": 666
          },
          {
            "AmountDiscount": 93,
            "AmountTotal": 5488,
            "CreatedAt": "2025-03-13T22:42:34.383655+08:00",
            "ID": 90,
            "OrderType": "虚拟商品订单",
            "ShippingFee": 83,
            "TenantID": 888,
            "UpdatedAt": "2025-03-13T22:42:34.383655+08:00",
            "UserID": 666
          }
        ]
--- PASS: TestCreateMapList/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id) (0.00s)

=== RUN   TestCreateMapList/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id);_Reinforced:_before_call_hooks
SELECT count(*) FROM `tbl_order` WHERE (1!=1
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((6500,50,266),(8715,32,379),(8067,83,153)))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order` (`amount_actual`,`amount_discount`,`amount_total`,`created_at`,`order_status`,`order_type`,`serial`,`shipping_fee`,`tenant_id`,`updated_at`,`user_id`) VALUES
(6716,50,6500,"2025-03-13 22:42:34.384",1,2,"bea3d3d0-753b-4e35-ba67-a10e1097558c",266,888,"2025-03-13 22:42:34.384",666),
(9062,32,8715,"2025-03-13 22:42:34.384",1,4,"0029e0bb-0ecd-4660-b6c7-e0f3b048029e",379,888,"2025-03-13 22:42:34.384",666),
(8137,83,8067,"2025-03-13 22:42:34.384",1,3,"62f028a6-f30d-490b-b789-0ae643c0bba4",153,888,"2025-03-13 22:42:34.384",666) RETURNING `id`
    api_create_test.go:194: [
          {
            "AmountActual": 6716,
            "AmountDiscount": 50,
            "AmountTotal": 6500,
            "CreatedAt": "2025-03-13T22:42:34.384517+08:00",
            "ID": 91,
            "OrderStatus": "待支付",
            "OrderType": "团购订单",
            "Serial": "bea3d3d0-753b-4e35-ba67-a10e1097558c",
            "ShippingFee": 266,
            "TenantID": 888,
            "UpdatedAt": "2025-03-13T22:42:34.384517+08:00",
            "UserID": 666
          },
          {
            "AmountActual": 9062,
            "AmountDiscount": 32,
            "AmountTotal": 8715,
            "CreatedAt": "2025-03-13T22:42:34.384517+08:00",
            "ID": 92,
            "OrderStatus": "待支付",
            "OrderType": "虚拟商品订单",
            "Serial": "0029e0bb-0ecd-4660-b6c7-e0f3b048029e",
            "ShippingFee": 379,
            "TenantID": 888,
            "UpdatedAt": "2025-03-13T22:42:34.384517+08:00",
            "UserID": 666
          },
          {
            "AmountActual": 8137,
            "AmountDiscount": 83,
            "AmountTotal": 8067,
            "CreatedAt": "2025-03-13T22:42:34.384517+08:00",
            "ID": 93,
            "OrderStatus": "待支付",
            "OrderType": "秒杀订单",
            "Serial": "62f028a6-f30d-490b-b789-0ae643c0bba4",
            "ShippingFee": 153,
            "TenantID": 888,
            "UpdatedAt": "2025-03-13T22:42:34.384517+08:00",
            "UserID": 666
          }
        ]
--- PASS: TestCreateMapList/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id);_Reinforced:_before_call_hooks (0.00s)

=== RUN   TestCreateMapList/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id);_Reinforced:_before_and_after_call_hooks
SELECT count(*) FROM `tbl_order` WHERE (1!=1
OR (`amount_total`,`amount_discount`,`shipping_fee`) IN ((8459,20,218),(8241,20,36),(4964,3,171)))
AND `tbl_order`.`deleted_at` IS NULL
AND `tbl_order`.`tenant_id` = 888
AND `tbl_order`.`user_id` = 666

INSERT INTO `tbl_order` (`amount_actual`,`amount_discount`,`amount_total`,`created_at`,`order_status`,`order_type`,`serial`,`shipping_fee`,`tenant_id`,`updated_at`,`user_id`) VALUES
(8657,20,8459,"2025-03-13 22:42:34.385",1,3,"c0044135-7d5c-4540-b6ff-f34a6f4bff8c",218,888,"2025-03-13 22:42:34.385",666),
(8257,20,8241,"2025-03-13 22:42:34.385",1,3,"2c462ac0-c411-4f9a-b75a-071ca358e7c0",36,888,"2025-03-13 22:42:34.385",666),
(5132,3,4964,"2025-03-13 22:42:34.385",1,2,"23126f6c-c46f-43f7-bf40-af84f9ba9da5",171,888,"2025-03-13 22:42:34.385",666) RETURNING `id`
    api_create_test.go:205: [
          {
            "AmountActual": 8657,
            "AmountDiscount": 20,
            "AmountTotal": 8459,
            "CreatedAt": "2025-03-13T22:42:34.385575+08:00",
            "ID": 94,
            "OrderStatus": "待支付",
            "OrderType": "秒杀订单",
            "Serial": "c0044135-7d5c-4540-b6ff-f34a6f4bff8c",
            "ShippingFee": 218,
            "TenantID": 888,
            "UpdatedAt": "2025-03-13T22:42:34.385575+08:00"
          },
          {
            "AmountActual": 8257,
            "AmountDiscount": 20,
            "AmountTotal": 8241,
            "CreatedAt": "2025-03-13T22:42:34.385575+08:00",
            "ID": 95,
            "OrderStatus": "待支付",
            "OrderType": "秒杀订单",
            "Serial": "2c462ac0-c411-4f9a-b75a-071ca358e7c0",
            "ShippingFee": 36,
            "TenantID": 888,
            "UpdatedAt": "2025-03-13T22:42:34.385575+08:00"
          },
          {
            "AmountActual": 5132,
            "AmountDiscount": 3,
            "AmountTotal": 4964,
            "CreatedAt": "2025-03-13T22:42:34.385575+08:00",
            "ID": 96,
            "OrderStatus": "待支付",
            "OrderType": "团购订单",
            "Serial": "23126f6c-c46f-43f7-bf40-af84f9ba9da5",
            "ShippingFee": 171,
            "TenantID": 888,
            "UpdatedAt": "2025-03-13T22:42:34.385575+08:00"
          }
        ]
--- PASS: TestCreateMapList/CreateMap_IsPlugin:_Uniques,_2Scopes(user_id,_tenant_id);_Reinforced:_before_and_after_call_hooks (0.00s)

