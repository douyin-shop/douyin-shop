package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	user "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/user"
	userService "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	// 通过微服务调用user服务
	checkUserRes, err := rpc.UserClient.Login(h.Context, &userService.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		hlog.Error("userClient.Login err: ", err)
		return
	}

	hlog.Debug("账号密码校验成功，正在准备请求Auth分发token： ", checkUserRes.UserId)

	// 通过微服务调用auth服务
	authRes, err := rpc.AuthClient.DeliverTokenByRPC(h.Context, &auth.DeliverTokenReq{UserId: checkUserRes.UserId})

	if err != nil {
		hlog.Error("authService.DeliverTokenByRPC err: ", err)
		return
	}

	hlog.Debug("token分发成功，正在返回结果： ", authRes.Token)

	// 返回结果
	resp = &user.LoginResp{
		UserId: checkUserRes.UserId,
		Token:  authRes.Token,
	}

	return
}
