package service

import (
	"context"
	order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
	"github.com/google/uuid"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {

	// 调用uuid生成订单号
	// 将订单存储到数据库
	u := uuid.New()

	resp = &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: u.String(),
		},
	}

	return
}
