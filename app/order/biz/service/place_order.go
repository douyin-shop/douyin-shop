package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/mysql"
	order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
	"github.com/douyin-shop/douyin-shop/app/order/utils/code"
	"github.com/jinzhu/copier"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {

	address := model.Address{}

	err = copier.Copy(&address, req.Address)
	if err != nil {
		klog.Error("copier.Copy failed", err)
		return nil, code.GetError(code.StructConvertError)
	}

	var orderItems []model.OrderItem

	for _, item := range req.OrderItems {
		orderItem := model.OrderItem{
			OrderID:   "",
			ProductID: item.Item.ProductId,
			Quantity:  item.Item.Quantity,
			Cost:      float64(item.Cost),
		}

		orderItems = append(orderItems, orderItem)
	}

	orderDetail := &model.Order{
		UserID:       req.UserId,
		UserCurrency: req.UserCurrency,
		Address:      address,
		Email:        req.Email,
		Status:       model.OrderStatusPending,
		OrderItems:   orderItems,
	}

	var orderId string
	if orderId, err = model.CreateOrderWithItems(mysql.DB, orderDetail); err != nil {
		klog.Error("CreateOrderWithItems failed", err)
		return nil, err
	}

	resp = &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: orderId,
		},
	}

	return
}
