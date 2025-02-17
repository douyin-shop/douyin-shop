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

type ListProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListProductsService(Context context.Context, RequestContext *app.RequestContext) *ListProductsService {
	return &ListProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	remoteListProductsReq := &remote_product.ListProductsReq{}
	err = copier.Copy(remoteListProductsReq, req)
	if err != nil {
		return nil, err
	}

	remoteProductResp, err := rpc.ProductClient.ListProducts(h.Context, remoteListProductsReq)
	if err != nil {
		return nil, err
	}

	resp = &product.ListProductsResp{}
	err = copier.Copy(resp, remoteProductResp)
	if err != nil {
		return nil, err
	}
	return
}
