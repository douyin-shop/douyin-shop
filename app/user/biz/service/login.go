package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.

	klog.Debug("req: ", req)

	resp = new(user.LoginResp)

	resp.UserId = 1

	return
}
