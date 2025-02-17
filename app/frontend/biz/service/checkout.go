package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"
	"github.com/jinzhu/copier"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	remote_checkout "github.com/douyin-shop/douyin-shop/app/checkout/kitex_gen/checkout"
	checkout "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/checkout"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	userId := h.Context.Value("user_id")

	hlog.Info("请求用户id:", userId)

	// 将用户id转换为int
	userIdInt, err := strconv.Atoi(userId.(string))
	if err != nil {
		return nil, err
	}

	remoteCheckReq := &remote_checkout.CheckoutReq{}
	err = copier.Copy(remoteCheckReq, req)
	remoteCheckReq.UserId = uint32(userIdInt)
	if err != nil {
		return nil, err
	}
	hlog.Debug("remoteCheckReq:", remoteCheckReq)
	checkoutResp, err := rpc.CheckoutClient.Checkout(h.Context, remoteCheckReq)
	if err != nil {
		return nil, err
	}

	resp = &checkout.CheckoutResp{}
	err = copier.Copy(resp, checkoutResp)
	if err != nil {
		return nil, err
	}

	return
}
