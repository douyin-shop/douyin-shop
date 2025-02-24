package model

import (
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	productMysql "github.com/douyin-shop/douyin-shop/app/product/biz/dal/mysql"
	"gorm.io/gorm"
)

type ProductInfo struct {
	Price       float32 `gorm:"float;not null;" json:"price" binding:"required" label:"商品单价"`
	TotalAmount float32 `gorm:"not null" json:"total_amount" binding:"required" label:"订单总金额"`
	ProductName string  `gorm:"type:varchar(255);not null" json:"product_name" binding:"required" label:"商品名称"`
}

type OrderItem struct {
	gorm.Model
	OrderId     string  `gorm:"type:varchar(36);uniqueIndex;not null" json:"order_id" binding:"required" label:"订单ID"`
	ProductId   uint32  `gorm:"type:varchar(36);not null" json:"product_id" binding:"required" label:"商品ID"`
	Quantity    int32   `gorm:"int;not null;" json:"quantity" binding:"required" label:"购买数量"`
	Price       float32 `gorm:"float;not null;" json:"price" binding:"required" label:"商品单价"`
	TotalAmount float32 `gorm:"not null" json:"total_amount" binding:"required" label:"订单总金额"`
	ProductName string  `gorm:"type:varchar(255);not null" json:"product_name" binding:"required" label:"商品名称"`
}

func GetOrderItemsByOrderItemIds(ids []uint32) (OrderItems []OrderItem, err error) {
	db := mysql.DB
	err = db.Where("id in (?)", ids).Find(&OrderItems).Error
	return
}

func GetProductInfo(productId uint32, quantity int32) (ProductInfo, error) {
	var product model.Product
	db := productMysql.Db
	err := db.Where("id = ?", productId).First(&product).Error
	if err != nil {
		return ProductInfo{}, err
	}
	price := float32(product.Price)
	return ProductInfo{
		Price:       price,
		TotalAmount: price * float32(quantity),
		ProductName: product.Name,
	}, nil
}
