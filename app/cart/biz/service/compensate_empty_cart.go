package service

import (
	"context"
	cart "github.com/douyin-shop/douyin-shop/app/cart/kitex_gen/cart"
)

type CompensateEmptyCartService struct {
	ctx context.Context
} // NewCompensateEmptyCartService new CompensateEmptyCartService
func NewCompensateEmptyCartService(ctx context.Context) *CompensateEmptyCartService {
	return &CompensateEmptyCartService{ctx: ctx}
}

// Run create note info
func (s *CompensateEmptyCartService) Run(req *cart.RestoreCartItemsReq) (resp *cart.RestoreCartItemsResp, err error) {
	// Finish your business logic.

	return
}
