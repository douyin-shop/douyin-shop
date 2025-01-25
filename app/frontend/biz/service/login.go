package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth/authservice"
	frontend "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend"
	"github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
	"github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user/userservice"
	"github.com/douyin-shop/douyin-shop/common/nacos"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *frontend.LoginReq) (resp *frontend.LoginResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	// 通过微服务调用user服务
	resolver := nacos.GetNacosResolver()
	userClient := userservice.MustNewClient("user", client.WithResolver(resolver))
	checkUserRes, err := userClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		hlog.Error("userClient.Login err: ", err)
		return
	}

	hlog.Debug("账号密码校验成功，正在准备请求Auth分发token： ", checkUserRes.UserId)

	// 通过微服务调用auth服务
	authService := authservice.MustNewClient("auth", client.WithResolver(resolver))
	authRes, err := authService.DeliverTokenByRPC(h.Context, &auth.DeliverTokenReq{UserId: checkUserRes.UserId})

	if err != nil {
		hlog.Error("authService.DeliverTokenByRPC err: ", err)
		return
	}

	hlog.Debug("token分发成功，正在返回结果： ", authRes.Token)

	// 返回结果
	resp = &frontend.LoginResp{
		UserId: checkUserRes.UserId,
		Token:  authRes.Token,
	}

	return
}
