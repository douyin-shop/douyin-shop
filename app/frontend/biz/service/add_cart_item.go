package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	cartService "github.com/douyin-shop/douyin-shop/app/cart/kitex_gen/cart"
	cart "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/cart"
	common "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/common"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	// 获取用户id
	userId := h.Context.Value("user_id")

	hlog.Info("请求用户id:", userId)

	userIdInt, err := strconv.Atoi(userId.(string))
	if err != nil {
		return nil, err
	}

	_, err = rpc.CartClient.AddItem(h.Context, &cartService.AddItemReq{
		UserId: uint32(userIdInt),
		Item: &cartService.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	if err != nil {
		return nil, err
	}

	return
}
