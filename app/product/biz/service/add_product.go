package service

import (
	"context"
	product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
)

type AddProductService struct {
	ctx context.Context
} // NewAddProductService new AddProductService
func NewAddProductService(ctx context.Context) *AddProductService {
	return &AddProductService{ctx: ctx}
}

// Run create note info
func (s *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	
	return
}
