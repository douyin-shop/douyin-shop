package service

import (
	"context"
	order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
)

type MarkOrderCanceledService struct {
	ctx context.Context
} // NewMarkOrderCanceledService new MarkOrderCanceledService
func NewMarkOrderCanceledService(ctx context.Context) *MarkOrderCanceledService {
	return &MarkOrderCanceledService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderCanceledService) Run(req *order.MarkOrderCanceledReq) (resp *order.MarkOrderCanceledResp, err error) {
	// Finish your business logic.

	return
}
