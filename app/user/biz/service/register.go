package service

import (
	"context"

	"github.com/douyin-shop/douyin-shop/app/user/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/user/biz/dal/mysql"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/user/biz/utils/code"
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	userCode, _ := model.CheckUserExist(mysql.DB, req.Email)
	if userCode == code.UserExist {
		return nil, kerrors.NewGRPCBizStatusError(code.UserExist, code.GetMsg(code.UserExist))
	}
	newUser := &model.User{
		Email:    req.Email,
		PassWord: req.Password,
	}
	Id, err := model.CreateUser(mysql.DB, newUser)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(code.UserCreateFailed, code.GetMsg(code.UserCreateFailed))
	}
	return &user.RegisterResp{
		UserId: int32(Id),
	}, nil
}
