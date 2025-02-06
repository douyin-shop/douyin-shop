package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/common"
	user "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/user"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *common.Empty) (resp *user.LogoutResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	// 从请求的上下文中获取token
	token := string(h.RequestContext.GetHeader("Authorization"))

	// 调用auth服务来销毁token
	logoutRes, err := rpc.AuthClient.Logout(h.Context, &auth.LogoutReq{Token: token})

	if err != nil {
		hlog.Error("authClient.Logout err: ", err)
		return
	}

	// 返回结果
	resp = &user.LogoutResp{
		Success: logoutRes.Success,
	}

	return
}
