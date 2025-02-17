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

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	remoteSearchProductsReq := &remote_product.SearchProductsReq{}
	err = copier.Copy(remoteSearchProductsReq, req)
	if err != nil {
		return nil, err
	}

	remoteProductResp, err := rpc.ProductClient.SearchProducts(h.Context, remoteSearchProductsReq)
	if err != nil {
		return nil, err
	}

	resp = &product.SearchProductsResp{}
	err = copier.Copy(resp, remoteProductResp)
	if err != nil {
		return nil, err
	}
	return
}
