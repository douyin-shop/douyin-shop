// @Author Adrian.Wang 2025/2/7 15:26:00
package model

import (
	"context"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/redis"
	"gorm.io/gorm"
	"log"
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

func (pl PaymentLog) BeforeCreate(tx *gorm.DB) (err error) {

	redisKey := "payment:" + pl.TransactionId
	// 存储支付状态，初始为 "PENDING"
	err = redis.RedisClient.Set(context.Background(), redisKey, "PENDING", 30*time.Minute).Err()
	if err != nil {
		log.Printf("Failed to store payment in Redis: %v", err)
		return err
	}
	log.Printf("Payment transaction %s stored in Redis with status PENDING", pl.TransactionId)
	return nil

}

// CreatePaymentLog 创建支付日志
func CreatePaymentLog(db *gorm.DB, ctx context.Context, paymentLog *PaymentLog) error {
	// 插入数据库
	err := db.WithContext(ctx).Model(&PaymentLog{}).Create(paymentLog).Error
	if err != nil {
		return err
	}

	return nil
}
