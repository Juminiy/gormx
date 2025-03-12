package gormx_testv2

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Serial          string      `gorm:"primaryKey" x:"unique"`
	UserID          uint        `gorm:"index" x:"user"`
	TenantID        uint        `gorm:"index" x:"tenant"`
	AmountTotal     int64       `gorm:"not null" x:"unique:fee"`
	AmountDiscount  int64       `gorm:"not null" x:"unique:fee"`
	ShippingFee     int64       `gorm:"not null" x:"unique:fee,logistics"`
	AmountActual    int64       `gorm:"not null" x:"unique:fee"`
	OrderType       OrderType   `gorm:"not null"`
	OrderStatus     OrderStatus `gorm:"not null"`
	PayTime         sql.NullTime
	PayMethod       PayMethod `gorm:"not null"`
	ReceiverName    string    `gorm:"not null" x:"unique:receiver"`
	ReceiverPhone   string    `gorm:"not null" x:"unique:receiver"`
	ReceiverAddress string    `gorm:"not null" x:"unique:receiver"`
	ShippedTime     sql.NullTime
	FinishedTime    sql.NullTime
	LogisticsID     uint   `gorm:"not null" x:"unique:logistics"`
	LogisticsName   string `gorm:"not null" x:"unique:logistics"`
	ExtrasInfo      sql.NullString
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.Serial = uuid.NewString()
	o.AmountActual = o.AmountTotal - o.AmountDiscount + o.ShippingFee
	o.OrderStatus = StatusWaiting
	return nil
}

func RandomOrder() *Order {
	return &Order{
		AmountTotal:    int64(gofakeit.IntRange(1000, 10000)),
		AmountDiscount: int64(gofakeit.IntRange(1, 100)),
		ShippingFee:    int64(gofakeit.IntRange(1, 500)),
		OrderType:      OrderType(gofakeit.IntRange(1, 4)),
	}
}

type OrderType int

const (
	TypeRegular OrderType = 1
	TypeGroup   OrderType = 2
	TypeFlash   OrderType = 3
	TypeVirtual OrderType = 4
)

var ErrOrderType = errors.New("order type value is error")

func (t OrderType) String() string {
	switch t {
	case TypeRegular:
		return "普通订单"
	case TypeGroup:
		return "团购订单"
	case TypeFlash:
		return "秒杀订单"
	case TypeVirtual:
		return "虚拟商品订单"
	default:
		return "无效订单类型"
	}
}

func (t OrderType) Valid() bool {
	return util.ElemIn(t, TypeRegular, TypeGroup, TypeFlash, TypeVirtual)
}

func (t *OrderType) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &t); err != nil || !t.Valid() {
		return ErrOrderType
	}
	return nil
}

func (t OrderType) MarshalJSON() ([]byte, error) {
	if t.Valid() {
		return json.Marshal(t.String())
	}
	return nil, ErrOrderType
}

type OrderStatus int

const (
	StatusWaiting   OrderStatus = 1
	StatusPaid      OrderStatus = 2
	StatusShipping  OrderStatus = 3
	StatusFinished  OrderStatus = 4
	StatusCancel    OrderStatus = 5
	StatusRefunding OrderStatus = 6
)

var ErrOrderStatus = errors.New("order status value is error")

func (t OrderStatus) String() string {
	switch t {
	case StatusWaiting:
		return "待支付"
	case StatusPaid:
		return "已支付"
	case StatusShipping:
		return "运输中"
	case StatusFinished:
		return "已完成"
	case StatusCancel:
		return "已取消"
	case StatusRefunding:
		return "退款中"
	default:
		return "无效订单状态"
	}
}

func (t OrderStatus) Valid() bool {
	return util.ElemIn(t, StatusWaiting, StatusPaid, StatusShipping, StatusFinished, StatusCancel, StatusRefunding)
}

func (t *OrderStatus) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &t); err != nil || !t.Valid() {
		return ErrOrderStatus
	}
	return nil
}

func (t OrderStatus) MarshalJSON() ([]byte, error) {
	if t.Valid() {
		return json.Marshal(t.String())
	}
	return nil, ErrOrderStatus
}

type PayMethod int

const (
	MethodAlipay       PayMethod = 1
	MethodWechatPay    PayMethod = 2
	MethodUnionPay     PayMethod = 3
	MethodDeliveryCash PayMethod = 4
)

var ErrPayMethod = errors.New("order pay method value is error")

func (t PayMethod) String() string {
	switch t {
	case MethodAlipay:
		return "支付宝"
	case MethodWechatPay:
		return "微信支付"
	case MethodUnionPay:
		return "银联支付"
	case MethodDeliveryCash:
		return "货到付款"
	default:
		return "无效支付类型"
	}
}

func (t PayMethod) Valid() bool {
	return util.ElemIn(t, MethodAlipay, MethodWechatPay, MethodUnionPay, MethodDeliveryCash)
}

func (t *PayMethod) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &t); err != nil || !t.Valid() {
		return ErrPayMethod
	}
	return nil
}

func (t PayMethod) MarshalJSON() ([]byte, error) {
	if t.Valid() {
		return json.Marshal(t.String())
	}
	return nil, ErrPayMethod
}
