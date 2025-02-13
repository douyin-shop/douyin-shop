// @Author Adrian.Wang 2025/2/7 15:26:00
package model

import (
	"context"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/utils/code"
	"gorm.io/gorm"
	"time"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32    `json:"user_id" gorm:"comment:用户id"`
	OrderId       string    `json:"order_id" gorm:"comment:订单id"`
	Amount        float32   `json:"amount" gorm:"comment:支付金额"`
	TransactionId string    `json:"transaction_id" gorm:"comment:交易id"`
	PayTime       time.Time `json:"pay_time" gorm:"comment:支付时间"`
}

// TableName 设置表名
func (PaymentLog) TableName() string {
	return "payment_log"
}

// CreatePaymentLog 创建支付日志
func CreatePaymentLog(db *gorm.DB, ctx context.Context, paymentLog *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(paymentLog).Error
}

func GetOrderIDByTransaction(db *gorm.DB, ctx context.Context, transactionID string) (int, error) {
	var paymentLog PaymentLog
	result := db.WithContext(ctx).
		Select("order_id").
		Where("transaction_id = ?", transactionID).
		First(&paymentLog)
	err := result.Error
	if err != nil {
		return code.FailedPayment, err
	}
	return code.Success, nil
}
