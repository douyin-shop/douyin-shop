package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/user/biz/model"
	"github.com/douyin-shop/douyin-shop/app/user/code"
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
)

type GetService struct {
	ctx context.Context
} // NewGetService new GetService
func NewGetService(ctx context.Context) *GetService {
	return &GetService{ctx: ctx}
}

// Run create note info
func (s *GetService) Run(req *user.GetReq) (resp *user.GetResp, err error) {
	var u *model.User
	if req.UserId == 0 {
		return nil, kerrors.NewGRPCBizStatusError(code.UserNotExist, "UserId不存在")
	}
	userCode, u := model.CheckUserIDExist(req.UserId)
	if userCode == code.UserExist {

		return &user.GetResp{

			Email:    u.Email,
			Role:     int32(u.Role),
			Password: u.PassWord,
		}, nil
	}
	return nil, kerrors.NewGRPCBizStatusError(code.UserNotExist, code.GetMsg(userCode))
}
