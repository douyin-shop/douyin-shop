package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/checkout/rpc"
	remote_payment "github.com/douyin-shop/douyin-shop/app/payment/kitex_gen/payment"
	"github.com/jinzhu/copier"

	"github.com/cloudwego/hertz/pkg/app"
	payment "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/payment"
)

type PaymentCallbackService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPaymentCallbackService(Context context.Context, RequestContext *app.RequestContext) *PaymentCallbackService {
	return &PaymentCallbackService{RequestContext: RequestContext, Context: Context}
}

func (h *PaymentCallbackService) Run(req *payment.PaymentCallbackReq) (resp *payment.PaymentCallbackResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	remoteSearchProductsReq := &remote_payment.PaymentCallbackReq{}
	err = copier.Copy(remoteSearchProductsReq, req)
	if err != nil {
		return nil, err
	}

	paymentRes, err := rpc.PaymentClient.PaymentCallback(h.Context, remoteSearchProductsReq)
	if err != nil {
		return nil, err
	}

	resp = &payment.PaymentCallbackResp{}
	err = copier.Copy(resp, paymentRes)
	if err != nil {
		return nil, err
	}

	return
}
