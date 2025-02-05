package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/cart/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/cart/biz/dal/mysql"
	cart "github.com/douyin-shop/douyin-shop/app/cart/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {

	list, err := model.GetCartByUserId(s.ctx, mysql.DB, int32(req.UserId))
	if err != nil {
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}
	var items []*cart.CartItem

	for _, item := range list {
		items = append(items, &cart.CartItem{
			ProductId: uint32(item.ProductId),
			Quantity:  item.Quantity,
		})
	}

	return &cart.GetCartResp{Cart: &cart.Cart{Items: items}}, nil
}
