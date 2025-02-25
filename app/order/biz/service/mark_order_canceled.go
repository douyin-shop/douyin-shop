package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/mysql"
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

	orderId := req.OrderId

	err = model.MarkOrderCanceled(mysql.DB, orderId)
	if err != nil {
		klog.Error("mark order canceled failed", err)
		return nil, err
	}

	resp = &order.MarkOrderCanceledResp{}
	return
}
