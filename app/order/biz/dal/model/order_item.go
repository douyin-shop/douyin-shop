package model

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderId     string  `gorm:"type:varchar(36);uniqueIndex;not null" json:"order_id" binding:"required" label:"订单ID"`
	ProductId   uint32  `gorm:"type:varchar(36);not null" json:"product_id" binding:"required" label:"商品ID"`
	Quantity    int32   `gorm:"int;not null;" json:"quantity" binding:"required" label:"购买数量"`
	Price       float32 `gorm:"float;not null;" json:"price" binding:"required" label:"商品单价"`
	TotalAmount float32 `gorm:"not null" json:"total_amount" binding:"required" label:"订单总金额"`
	ProductName string  `gorm:"type:varchar(255);not null" json:"product_name" binding:"required" label:"商品名称"`
}
