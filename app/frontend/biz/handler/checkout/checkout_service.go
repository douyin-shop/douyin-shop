package checkout

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/service"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/utils"
	checkout "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/checkout"
)

// Checkout .
// @router checkout [POST]
func Checkout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &checkout.CheckoutResp{}
	resp, err = service.NewCheckoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
