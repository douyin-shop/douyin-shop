package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/utils/code"
	product "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/product"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"
	remote_product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
	"github.com/jinzhu/copier"
	"io"
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

	imageFile, err := h.RequestContext.FormFile("image_file")
	if err != nil {
		return nil, code.GetErr(code.ImageRequired)
	}

	imageFileHandler, err := imageFile.Open()
	if err != nil {
		return nil, code.GetErr(code.ImageOpenFailed)
	}
	defer imageFileHandler.Close()

	// 读取图片文件
	imageData, err := io.ReadAll(imageFileHandler)
	if err != nil {
		return nil, code.GetErr(code.ImageOpenFailed)
	}

	// 获取图片文件,转化成二进制
	imageName := imageFile.Filename

	err = copier.Copy(remoteAddProductReq, req)
	if err != nil {
		return nil, err
	}

	remoteAddProductReq.Product.ImageName = imageName
	remoteAddProductReq.Product.ImageUrl = ""
	remoteAddProductReq.Product.Image = imageData

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
