package main

import (
	"context"
	"github.com/douyin-shop/douyin-shop/app/order/biz/service"
	order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = service.NewMarkOrderPaidService(ctx).Run(req)

	return resp, err
}

// MarkOrderCanceled implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderCanceled(ctx context.Context, req *order.MarkOrderCanceledReq) (resp *order.MarkOrderCanceledResp, err error) {
	resp, err = service.NewMarkOrderCanceledService(ctx).Run(req)

	return resp, err
}
