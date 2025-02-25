// Package model @Author Adrian.Wang 2025/2/25 12:35:00
package model

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// OrderStatus 订单状态枚举（与proto枚举对应）
type OrderStatus string

const (
	OrderStatusUnspecified OrderStatus = "unspecified" // 未指定
	OrderStatusPending     OrderStatus = "pending"     // 待支付
	OrderStatusPaid        OrderStatus = "paid"        // 已支付
	OrderStatusCanceled    OrderStatus = "canceled"    // 已取消
)

// Order 主订单模型（包含gorm.Model基础字段）
type Order struct {
	gorm.Model
	OrderID      string      `gorm:"uniqueIndex;type:varchar(36);column:order_id"` // 业务用订单号
	UserID       uint32      `gorm:"index;not null"`
	UserCurrency string      `gorm:"type:varchar(10);not null"`
	Address      Address     `gorm:"embedded;embeddedPrefix:address_"`
	Email        string      `gorm:"type:varchar(255);not null"`
	Status       OrderStatus `gorm:"type:varchar(20);index;default:pending"`
	CanceledAt   *time.Time  `gorm:"index"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnDelete:CASCADE;"`
}

// OrderItem 订单商品项（包含gorm.Model基础字段）
type OrderItem struct {
	gorm.Model
	OrderID   string  `gorm:"type:varchar(36);index;not null"` // 关联Order.OrderID
	ProductID uint32  `gorm:"index;not null"`
	Quantity  int32   `gorm:"not null"`
	Cost      float64 `gorm:"type:decimal(10,2);not null"` // 使用decimal精确存储金额
}

// Address 地址信息（嵌入结构体）
type Address struct {
	StreetAddress string `gorm:"type:varchar(255);not null"`
	City          string `gorm:"type:varchar(100);not null"`
	State         string `gorm:"type:varchar(100);not null"`
	Country       string `gorm:"type:varchar(100);not null"`
	ZipCode       int32  `gorm:"not null"`
}

// BeforeCreate 钩子函数 - 自动生成业务订单号
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.OrderID = uuid.New().String()
	return nil
}

// TableName 设置表名
func (*Order) TableName() string {
	return "orders"
}

// TableName 设置表名
func (*OrderItem) TableName() string {
	return "order_items"
}

// TableName 设置表名
func (*Address) TableName() string {
	return "order_addresses"
}

// CreateOrderWithItems 创建订单及关联商品项（包含事务处理）
func CreateOrderWithItems(db *gorm.DB, order *Order) (string, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1. 创建主订单记录（BeforeCreate钩子会自动生成OrderID）
		if err := tx.Create(order).Error; err != nil {
			klog.Error("CreateOrderWithItems failed", err)
			return err
		}

		return nil
	})

	// 返回生成的订单号
	if err != nil {
		return "", err
	}
	return order.OrderID, nil
}

// GetOrderByID 根据订单号查询订单及关联商品项
func GetOrderByID(db *gorm.DB, orderID string) (*Order, error) {
	var order Order
	err := db.Preload("OrderItems").
		Where("order_id = ?", orderID).
		First(&order).Error
	return &order, err
}

// ListOrders 根据用户ID查询订单列表（包含预加载关联商品项）
func ListOrders(db *gorm.DB, userID uint32) ([]*Order, error) {
	var orders []*Order

	// 执行查询（自动包含软删除过滤）
	err := db.Preload("OrderItems"). // 预加载关联商品项
						Where("user_id = ?", userID). // 用户ID过滤
						Order("created_at DESC").     // 按创建时间倒序排列
						Find(&orders).Error           // 查询结果注入切片

	if err != nil {
		klog.Errorf("ListOrders failed for user %d: %v", userID, err)
		return nil, err
	}

	return orders, nil
}

// MarkOrderPaid 修改订单状态为已支付
func MarkOrderPaid(db *gorm.DB, orderID string) error {
	// 判断订单状态是否为待支付
	order := Order{}
	err := db.Where("order_id = ? AND status = ?", orderID, OrderStatusPending).
		First(&order).Error
	if err != nil {
		klog.Error("MarkOrderPaid failed", err)
		return err
	}

	// 更新订单状态为已支付
	err = db.Model(&order).
		Updates(map[string]interface{}{"status": OrderStatusPaid}).Error
	if err != nil {
		klog.Error("MarkOrderPaid failed", err)
		return err
	}

	return nil

}
