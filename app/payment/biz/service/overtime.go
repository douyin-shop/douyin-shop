package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/douyin-shop/douyin-shop/app/payment/biz/utils/code"
	payment "github.com/douyin-shop/douyin-shop/app/payment/kitex_gen/payment"

	"net/http"
)

func (s *ChargeService) SimulatePaymentOvertime(req *payment.ChargeReq, c *app.RequestContext, transactionID string) {

	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    code.FailedPayment,
			"message": code.GetMsg(code.FailedPayment),
		})
		return
	}

	if err := s.handleOvertime(transactionID); err != code.Success {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    code.FailedPayment,
			"message": code.GetMsg(code.FailedPayment),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    code.Success,
		"message": code.GetMsg(code.Success),
		"data": map[string]string{
			"transaction_id": transactionID,
			"status":         "timeout",
		},
	})
}

func (s *ChargeService) handleOvertime(transactionID string) int {

	paymentKey := fmt.Sprintf("payment:%s", transactionID)
	if err := s.redis.Del(s.ctx, paymentKey).Err(); err != nil {
		return code.FailedPayment
	}

	return code.Overtime
}
