package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	product "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/product"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"
	remote_product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
	"github.com/jinzhu/copier"
)

type AddProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddProductService(Context context.Context, RequestContext *app.RequestContext) *AddProductService {
	return &AddProductService{RequestContext: RequestContext, Context: Context}
}

// Run 添加商品功能
// 调用链路：frontend -> product
func (h *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	remoteAddProductReq := &remote_product.AddProductReq{}
	err = copier.Copy(remoteAddProductReq, req)
	if err != nil {
		return nil, err
	}

	remoteProductResp, err := rpc.ProductClient.AddProduct(h.Context, remoteAddProductReq)
	if err != nil {
		return nil, err
	}

	resp = &product.AddProductResp{}
	err = copier.Copy(resp, remoteProductResp)
	if err != nil {
		return nil, err
	}
	return

}
