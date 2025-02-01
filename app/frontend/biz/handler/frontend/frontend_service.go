package frontend

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/service"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/utils"
	frontend "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Login .
// @router /login [GET]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req frontend.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &frontend.LoginResp{}
	resp, err = service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetCart .
// @router cart/get_cart [POST]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req frontend.GetCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetCartService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// Register .
// @router /register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req frontend.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewRegisterService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// Logout .
// @router /logout [GET]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req emptypb.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewLogoutService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
