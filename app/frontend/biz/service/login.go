package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
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

	resolver := nacos.GetNacosResolver()

	newClient := userservice.MustNewClient("user", client.WithResolver(resolver))

	res, err := newClient.Login(h.Context, &user.LoginReq{
		Email:    "351244716@qq.com",
		Password: "88888888",
	})

	hlog.Debug("res: ", res.UserId)
	return
}
