package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/order/biz/utils/code"
	"github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	//查询一个用户的所有订单
	//首先检查请求是否合法
	if req.UserId <= 0 {
		return nil, kerrors.NewGRPCBizStatusError(code.InvalidRequest, code.GetMsg(code.InvalidRequest))
	}

	db := mysql.DB
	//查询订单
	orders, err := model.GetOrdersByUserId(db, req.UserId)
	if err != nil || orders == nil {
		return nil, kerrors.NewGRPCBizStatusError(code.InternalError, code.GetMsg(code.InternalError))
	}

	//构造返回数据
	resp, err = convertOrders(orders)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(code.InternalError, code.GetMsg(code.InternalError))
	}
	return resp, nil
}

func convertOrders(orders []model.Order) (*order.ListOrderResp, error) {
	var orderlists []*order.Order
	for _, instance := range orders {
		orderItemIds := instance.OrderItemIdList
		//根据订单商品ID查询订单商品
		orderItems, err := model.GetOrderItemsByOrderItemIds(orderItemIds)
		if err != nil || orderItems == nil {
			return nil, kerrors.NewGRPCBizStatusError(code.InternalError, code.GetMsg(code.InternalError))
		}
		orderItemResps := getOrderItemsResp(orderItems)
		address := order.Address{
			StreetAddress: instance.Address.StreetAddress,
			City:          instance.Address.City,
			State:         instance.Address.State,
			Country:       instance.Address.Country,
			ZipCode:       instance.Address.ZipCode,
		}
		orderResp := order.Order{
			OrderItems:   orderItemResps,
			OrderId:      instance.OrderId,
			UserId:       instance.UserId,
			UserCurrency: instance.UserCurrency,
			Address:      &address,
			Email:        instance.Email,
			CreatedAt:    int32(instance.PlaceOrderTime.Unix()),
		}
		orderlists = append(orderlists, &orderResp)
	}
	return &order.ListOrderResp{
		Orders: orderlists,
	}, nil
}

func getOrderItemsResp(orderItems []model.OrderItem) []*order.OrderItem {
	orderItemResps := make([]*order.OrderItem, 0)
	for _, orderItem := range orderItems {
		orderItemResp := &order.OrderItem{
			Item: &order.CartItem{
				ProductId: orderItem.ProductId,
				Quantity:  orderItem.Quantity,
			},
			Cost: orderItem.TotalAmount,
		}
		orderItemResps = append(orderItemResps, orderItemResp)
	}
	return orderItemResps
}
