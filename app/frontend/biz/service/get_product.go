package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"
	remote_product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
	"github.com/jinzhu/copier"

	"github.com/cloudwego/hertz/pkg/app"
	product "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/product"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	remoteGetProductReq := &remote_product.GetProductReq{}
	err = copier.Copy(remoteGetProductReq, req)
	if err != nil {
		return nil, err
	}

	remoteProductResp, err := rpc.ProductClient.GetProduct(h.Context, remoteGetProductReq)
	if err != nil {
		return nil, err
	}

	resp = &product.GetProductResp{}
	err = copier.Copy(resp, remoteProductResp)
	if err != nil {
		return nil, err
	}
	return
}
