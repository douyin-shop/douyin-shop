package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/cart/kitex_gen/cart"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/mysql"
	order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
	"github.com/douyin-shop/douyin-shop/app/order/rpc"
	"github.com/douyin-shop/douyin-shop/app/order/utils/code"
)

type MarkOrderCanceledService struct {
	ctx context.Context
} // NewMarkOrderCanceledService new MarkOrderCanceledService
func NewMarkOrderCanceledService(ctx context.Context) *MarkOrderCanceledService {
	return &MarkOrderCanceledService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderCanceledService) Run(req *order.MarkOrderCanceledReq) (resp *order.MarkOrderCanceledResp, err error) {

	orderId := req.OrderId

	err = model.MarkOrderCanceled(mysql.DB, orderId)
	if err != nil {
		klog.Error("mark order canceled failed", err)
		return nil, err
	}

	// 获取订单对应的购物车商品
	orderDetail, err := model.GetOrderByID(mysql.DB, orderId)
	if err != nil {
		klog.Error("get order by id failed", err)
		return nil, code.GetError(code.GetOrderError)
	}

	// 通知购物车服务加回购物车商品
	for _, orderItem := range orderDetail.OrderItems {
		_, err := rpc.CartClient.AddItem(s.ctx, &cart.AddItemReq{
			UserId: orderDetail.UserID,
			Item: &cart.CartItem{
				ProductId: orderItem.ProductID,
				Quantity:  orderItem.Quantity,
			},
		})
		if err != nil {
			klog.Error("add item to cart failed", err)
			return nil, code.GetError(code.RecoversError)
		}
	}

	resp = &order.MarkOrderCanceledResp{}
	return
}
