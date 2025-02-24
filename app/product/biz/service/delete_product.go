package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/mysql"
	product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	c := model.DeleteProduct(int(req.Id), mysql.Db)
	if c == code.Error {
		return nil, kerrors.NewGRPCBizStatusError(int32(c), code.GetMessage(c))
	}
	return &product.DeleteProductResp{
		Success: true,
	}, nil
}
