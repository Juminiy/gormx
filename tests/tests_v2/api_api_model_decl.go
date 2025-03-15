package gormx_testv2

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Serial          string         `gorm:"primaryKey" x:"unique" json:",omitempty"`
	UserID          uint           `gorm:"index" x:"user" json:",omitempty"`
	TenantID        uint           `gorm:"index" x:"tenant" json:",omitempty"`
	AmountTotal     int64          `gorm:"not null" x:"unique:fee" json:",omitempty"`
	AmountDiscount  int64          `gorm:"not null" x:"unique:fee" json:",omitempty"`
	ShippingFee     int64          `gorm:"not null" x:"unique:fee,logistics" json:",omitempty"`
	AmountActual    int64          `x:"unique:-" json:",omitempty"`
	OrderType       OrderType      `gorm:"not null" json:",omitempty"`
	OrderStatus     OrderStatus    `json:",omitempty"`
	PayTime         sql.NullTime   `json:",omitempty"`
	PayMethod       PayMethod      `json:",omitempty"`
	ReceiverName    string         `x:"unique:receiver" json:",omitempty"`
	ReceiverPhone   string         `x:"unique:receiver" json:",omitempty"`
	ReceiverAddress string         `x:"unique:receiver" json:",omitempty"`
	ShippedTime     sql.NullTime   `json:",omitempty"`
	FinishedTime    sql.NullTime   `json:",omitempty"`
	LogisticsID     uint           `x:"unique:logistics" json:",omitempty"`
	LogisticsName   string         `x:"unique:logistics" json:",omitempty"`
	ExtrasInfo      sql.NullString `json:",omitempty"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.Serial = uuid.NewString()
	o.AmountActual = o.AmountTotal - o.AmountDiscount + o.ShippingFee
	o.OrderStatus = StatusWaiting
	return nil
}

func (o *Order) AfterCreate(tx *gorm.DB) error {
	o.UserID = 0
	return nil
}

func (o *Order) JSONString() string {
	b, _ := json.MarshalIndent(o, "", "  ")
	return string(b)
}

func (o *Order) SetPayInfo() *Order {
	o.PayTime = sql.NullTime{Time: time.Now().Add(util.DurationDay), Valid: true}
	o.PayMethod = PayMethod(gofakeit.IntRange(1, 4))
	o.OrderStatus = StatusPaid
	return o
}

func (o *Order) SetShipInfo() *Order {
	o.ShippedTime = sql.NullTime{Time: time.Now().Add(util.DurationDay * 7), Valid: true}
	o.LogisticsID = gofakeit.UintRange(1<<10, 1<<16)
	o.LogisticsName = gofakeit.Company()
	return o
}

func RandomOrder() *Order {
	return &Order{
		AmountTotal:     int64(gofakeit.IntRange(1000, 10000)),
		AmountDiscount:  int64(gofakeit.IntRange(1, 100)),
		ShippingFee:     int64(gofakeit.IntRange(1, 500)),
		OrderType:       OrderType(gofakeit.IntRange(1, 4)),
		ReceiverName:    gofakeit.Name(),
		ReceiverPhone:   gofakeit.Phone(),
		ReceiverAddress: gofakeit.Address().Address,
	}
}

func RandomOrderMap() map[string]any {
	return map[string]any{
		"AmountTotal":    int64(gofakeit.IntRange(1000, 10000)),
		"AmountDiscount": int64(gofakeit.IntRange(1, 100)),
		"ShippingFee":    int64(gofakeit.IntRange(1, 500)),
		"OrderType":      OrderType(gofakeit.IntRange(1, 4)),
	}
}

func RandomShipMap(shipFee int64) map[string]any {
	return map[string]any{
		"shipped_time":   time.Now().Add(util.DurationDay * 7),
		"shipping_fee":   shipFee,
		"logistics_id":   gofakeit.UintRange(1<<10, 1<<16),
		"logistics_name": gofakeit.Company(),
	}
}

type OrderType int

const (
	TypeInvalid OrderType = 0
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

func (t OrderType) Legal() bool { return t.Valid() || t == TypeInvalid }

func (t *OrderType) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &t); err != nil || !t.Legal() {
		return ErrOrderType
	}
	return nil
}

func (t OrderType) MarshalJSON() ([]byte, error) {
	if t.Legal() {
		return json.Marshal(t.String())
	}
	return nil, ErrOrderType
}

type OrderStatus int

const (
	StatusInvalid   OrderStatus = 0
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

func (t OrderStatus) Legal() bool {
	return t.Valid() || t == StatusInvalid
}

func (t *OrderStatus) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &t); err != nil || !t.Legal() {
		return ErrOrderStatus
	}
	return nil
}

func (t OrderStatus) MarshalJSON() ([]byte, error) {
	if t.Legal() {
		return json.Marshal(t.String())
	}
	return nil, ErrOrderStatus
}

type PayMethod int

const (
	MethodInvalid      PayMethod = 0
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

func (t PayMethod) Legal() bool {
	return t.Valid() || t == MethodInvalid
}

func (t *PayMethod) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &t); err != nil || !t.Legal() {
		return ErrPayMethod
	}
	return nil
}

func (t PayMethod) MarshalJSON() ([]byte, error) {
	if t.Legal() {
		return json.Marshal(t.String())
	}
	return nil, ErrPayMethod
}
