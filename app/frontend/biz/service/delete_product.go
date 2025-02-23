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

type DeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *DeleteProductService {
	return &DeleteProductService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	remoteDeleteProductsReq := &remote_product.DeleteProductReq{}
	err = copier.Copy(remoteDeleteProductsReq, req)
	if err != nil {
		return nil, err
	}

	remoteProductResp, err := rpc.ProductClient.DeleteProduct(h.Context, remoteDeleteProductsReq)
	if err != nil {
		return nil, err
	}

	resp = &product.DeleteProductResp{}
	err = copier.Copy(resp, remoteProductResp)
	if err != nil {
		return nil, err
	}
	return
}
