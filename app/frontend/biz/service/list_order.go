package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	order "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/order"
)

type ListOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListOrderService(Context context.Context, RequestContext *app.RequestContext) *ListOrderService {
	return &ListOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
