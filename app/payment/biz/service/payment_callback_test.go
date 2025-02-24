package service

import (
	"context"
	payment "github.com/douyin-shop/douyin-shop/app/payment/kitex_gen/payment"
	"testing"
)

func TestPaymentCallback_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPaymentCallbackService(ctx)
	// init req and assert value

	req := &payment.PaymentCallbackReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
