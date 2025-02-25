package service

import (
	"context"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/mysql"
	order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	orderId := req.OrderId

	err = model.MarkOrderPaid(mysql.DB, orderId)
	if err != nil {
		return nil, err
	}

	return &order.MarkOrderPaidResp{}, nil
}
