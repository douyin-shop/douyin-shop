package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	product "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/product"
)

type AddProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddProductService(Context context.Context, RequestContext *app.RequestContext) *AddProductService {
	return &AddProductService{RequestContext: RequestContext, Context: Context}
}

func (h *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
