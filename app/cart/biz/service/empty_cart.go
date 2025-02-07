package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/cart/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/cart/biz/dal/mysql"
	cart "github.com/douyin-shop/douyin-shop/app/cart/kitex_gen/cart"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {

	err = model.EmptyCart(s.ctx, mysql.DB, int32(req.UserId))
	if err != nil {
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}

	resp = &cart.EmptyCartResp{}

	return
}
