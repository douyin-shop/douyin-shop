package service

import (
	"context"
	product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.

	return &product.GetProductResp{
		Product: &product.Product{
			Id:          1,
			Name:        "test",
			Description: "test",
			Picture:     "test",
			Price:       100.11,
			Categories:  []string{"111", "222"},
		},
	}, nil
}
