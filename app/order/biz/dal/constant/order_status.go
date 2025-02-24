package constant

// OrderStatus 定义订单状态的常量

const (
	// Created 订单已创建
	Order_Created = iota
	// Paid 订单已支付
	Order_Paid
	// Shipped 订单已发货
	Order_Shipped
	// Finished 订单已完成
	Order_Finished
	// Cancelled 订单已取消
	Order_Cancelled
)
