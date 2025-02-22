package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"
	"github.com/jinzhu/copier"

	"github.com/cloudwego/hertz/pkg/app"
	order "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/order"
	remote_order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
)

type ListOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListOrderService(Context context.Context, RequestContext *app.RequestContext) *ListOrderService {
	return &ListOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	remoteListOrderReq := &remote_order.ListOrderReq{}
	err = copier.Copy(remoteListOrderReq, req)
	if err != nil {
		return nil, err
	}

	remoteOrderResp, err := rpc.OrderClient.ListOrder(h.Context, remoteListOrderReq)

	if err != nil {
		return nil, err
	}

	resp = &order.ListOrderResp{}
	err = copier.Copy(resp, remoteOrderResp)
	if err != nil {
		return nil, err
	}
	return
}
