package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/utils"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	user "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/user"
	userService "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	// 通过微服务调用user服务
	registerRes, err := rpc.UserClient.Register(h.Context, &userService.RegisterReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		hlog.Error("userClient.Register err: ", err)
		return
	}

	hlog.Debug("账号注册成功 ", registerRes.UserId)

	// 为用户添加角色
	utils.AddRoleToUser(mysql.DB, registerRes.UserId, "user")

	// 返回结果
	resp = &user.RegisterResp{
		UserId: registerRes.UserId,
	}
	return
}
