package main

import (
	"context"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/service"
	payment "github.com/douyin-shop/douyin-shop/app/payment/kitex_gen/payment"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}

// PaymentCallback implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) PaymentCallback(ctx context.Context, req *payment.PaymentCallbackReq) (resp *payment.PaymentCallbackResp, err error) {
	resp, err = service.NewPaymentCallbackService(ctx).Run(req)

	return resp, err
}
