package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/mysql"
	order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
	"github.com/douyin-shop/douyin-shop/app/order/rpc"
	"github.com/douyin-shop/douyin-shop/app/order/utils/code"
	"github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
	"github.com/jinzhu/copier"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {

	orders, err := model.ListOrders(mysql.DB, req.UserId)
	if err != nil {
		return nil, code.GetError(code.ListOrderError)
	}

	var orderResult []*order.Order

	for _, orderModel := range orders {

		orderAddress := &order.Address{}
		var orderStatus order.OrderStatus

		switch orderModel.Status {
		case model.OrderStatusPending:
			orderStatus = order.OrderStatus_ORDER_STATUS_PENDING
		case model.OrderStatusPaid:
			orderStatus = order.OrderStatus_ORDER_STATUS_PAID
		case model.OrderStatusCanceled:
			orderStatus = order.OrderStatus_ORDER_STATUS_CANCELED
		default:
			orderStatus = order.OrderStatus_ORDER_STATUS_UNSPECIFIED
		}

		if err = copier.Copy(orderAddress, orderModel.Address); err != nil {
			klog.Error("copier.Copy failed", err)
			return nil, code.GetError(code.StructConvertError)
		}

		var orderItems []*order.OrderItem

		for _, orderModelItem := range orderModel.OrderItems {

			var orderName string

			getProductResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: orderModelItem.ProductID})
			if err != nil {
				klog.Error("rpc.ProductClient.GetProduct failed", err)
				orderName = "unknown"
			} else {
				orderName = getProductResp.Product.Name
			}
			orderItem := &order.OrderItem{
				Item: &order.CartItem{
					ProductId: orderModelItem.ProductID,
					Quantity:  orderModelItem.Quantity,
				},
				Name: orderName,
				Cost: float32(orderModelItem.Cost),
			}

			orderItems = append(orderItems, orderItem)
		}

		createdTime := int32(orderModel.CreatedAt.Second())
		var canceledTime int32
		if orderModel.CanceledAt != nil {
			canceledTime = int32(orderModel.CanceledAt.Second())
		}

		singleOrderResult := &order.Order{
			OrderItems:   orderItems,
			OrderId:      orderModel.OrderID,
			UserId:       orderModel.UserID,
			UserCurrency: orderModel.UserCurrency,
			Address:      orderAddress,
			Email:        orderModel.Email,
			CreatedAt:    createdTime,
			Status:       orderStatus,
			CanceledAt:   canceledTime,
		}

		orderResult = append(orderResult, singleOrderResult)

	}

	resp = &order.ListOrderResp{Orders: orderResult}

	return
}
