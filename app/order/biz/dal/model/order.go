package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	OrderId         string    `gorm:"type:varchar(36);uniqueIndex;not null" json:"order_id" binding:"required" label:"订单ID"`
	OrderItemIdList []uint32  `gorm:"type:varchar(36);not null" json:"order_item_id_list" binding:"required" label:"订单商品ID列表"`
	PaymentId       string    `gorm:"type:varchar(36);uniqueIndex" json:"payment_id" binding:"required" label:"支付ID"`
	TotalAmount     float32   `gorm:"not null" json:"total_amount" binding:"required" label:"订单总金额"`
	OrderStatus     int       `gorm:"tinyint;not null;" json:"order_status" binding:"required" label:"订单状态"`
	UserId          uint32    `gorm:"int;not null;" json:"user_id" binding:"required" label:"用户id"`
	Phone           string    `gorm:"varchar(20);not null;" json:"phone" binding:"required" label:"收货电话"`
	Email           string    `gorm:"varchar(30);not null;" json:"email" binding:"required,email" label:"用户邮箱"`
	Address         Address   `gorm:"type:varchar(100);not null" json:"address" binding:"required" label:"收货地址"`
	PaymentTime     time.Time `gorm:"default:null" json:"payment_time" label:"支付时间"`
	ShippingTime    time.Time `gorm:"default:null" json:"shipping_time" label:"发货时间"`
	RefundTime      time.Time `gorm:"default:null" json:"refund_time" label:"退款时间"`
	PlaceOrderTime  time.Time `gorm:"default:null" json:"place_order_time" label:"下单时间"`
	ShippingStatus  string    `gorm:"type:varchar(20);not null" json:"shipping_status" binding:"required" label:"发货状态"`
	RefundStatus    string    `gorm:"type:varchar(20);not null" json:"refund_status" binding:"required" label:"退款状态"`
}

type Address struct {
	StreetAddress string `protobuf:"bytes,1,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode       int32  `protobuf:"varint,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
}
