package service

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/bxcodec/faker/v3"
	order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
	"github.com/google/uuid"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {

	orders := make([]*order.Order, 10)
	orderItems := make([]*order.OrderItem, 10)

	// 生成10个假数据
	for i := 0; i < 10; i++ {

		randomPrice := gofakeit.Price(100, 10000)

		orderItems[i] = &order.OrderItem{
			Item: &order.CartItem{
				ProductId: uint32(i + 1),
				Quantity:  1,
			},
			Cost: float32(randomPrice),
		}

		u := uuid.New()

		orders[i] = &order.Order{
			OrderItems:   orderItems,
			OrderId:      u.String(),
			UserId:       1,
			UserCurrency: faker.Currency(),
			Address: &order.Address{
				StreetAddress: gofakeit.Address().Street,
				City:          gofakeit.City(),
				State:         gofakeit.State(),
				Country:       gofakeit.Country(),
				ZipCode:       gofakeit.Int32(),
			},
			Email:     gofakeit.Email(),
			CreatedAt: int32(gofakeit.Date().Unix()),
		}
	}

	resp = &order.ListOrderResp{Orders: orders}

	return
}
