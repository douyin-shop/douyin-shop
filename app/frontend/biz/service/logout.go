package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	frontend "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *emptypb.Empty) (resp *frontend.LogoutResp, err error) {
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
	resp = &frontend.LogoutResp{
		Success: logoutRes.Success,
	}

	return
}
