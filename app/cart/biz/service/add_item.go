package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/cart/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/cart/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/cart/biz/utils/code"
	cart "github.com/douyin-shop/douyin-shop/app/cart/kitex_gen/cart"
	"github.com/douyin-shop/douyin-shop/app/cart/rpc"
	"github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// 调用Product服务校验商品是否存在
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}

	if productResp == nil || productResp.Product == nil || productResp.Product.Id == 0 {
		return nil, code.GetError(code.NotFoundProduct)
	}

	cartItem := &model.Cart{
		UserId:    int32(req.UserId),
		ProductId: int32(req.Item.ProductId),
		Quantity:  req.Item.Quantity,
	}

	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
