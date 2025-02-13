package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/utils/code"
	payment "github.com/douyin-shop/douyin-shop/app/payment/kitex_gen/payment"
	"github.com/goccy/go-json"
	"net/http"
)

type Pay struct {
	Code int
	Msg  string
}

func (s *ChargeService) SimulatePaymentSuccess(req *payment.ChargeReq, c *app.RequestContext, transactionID string) string {

	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    code.FailedPayment,
			"message": code.GetMsg(code.FailedPayment),
		})
		return code.GetMsg(code.FailedPayment)
	}

	// 2. 执行支付成功处理
	err := s.HandlePaymentSuccess(transactionID)
	if err != code.Success {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    code.FailedPayment,
			"message": code.GetMsg(code.FailedPayment),
		})
		return code.GetMsg(code.FailedPayment)
	}

	// 3. 返回成功响应
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    code.Success,
		"message": code.GetMsg(code.Success),
		"data": map[string]string{
			"transaction_id": transactionID,
			"status":         "success",
		},
	})
	return code.GetMsg(code.Success)
}

// 支付成功处理
func (s *ChargeService) HandlePaymentSuccess(transactionID string) int {
	paymentKey := fmt.Sprintf("payment:%s", transactionID)
	if err := s.redis.Del(s.ctx, paymentKey).Err(); err != nil {
		return code.FailedPayment
	}

	orderID, err := model.GetOrderIDByTransaction(mysql.DB, s.ctx, transactionID)
	if err != nil {
		return code.FailedPayment
	}
	if err := s.sendToPaymentQueue(orderID, transactionID); err != code.Success {
		return code.Queueerr
	}

	return code.Overtime
}

// 启动Redis过期监听
func (s *ChargeService) handleOvertime(transactionID string) int {

	paymentKey := fmt.Sprintf("payment:%s", transactionID)
	if err := s.redis.Del(s.ctx, paymentKey).Err(); err != nil {
		return code.FailedPayment
	}

	return code.Overtime
}

type PaymentMessage struct {
	OrderID       int
	TransactionID string
}

func (s *ChargeService) sendToPaymentQueue(orderID int, transactionID string) int {
	msg := PaymentMessage{
		OrderID:       orderID,
		TransactionID: transactionID,
	}

	_, err := json.Marshal(msg)
	if err != nil {
		return code.Queueerr
	}

	// 使用Redis Stream确保消息持久化
	return code.Success
}

// HandlePaymentSuccess 处理支付成功
