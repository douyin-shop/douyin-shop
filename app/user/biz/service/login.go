package service

import (
	"context"

	"github.com/douyin-shop/douyin-shop/app/user/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/user/biz/dal/mysql"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/user/biz/utils/code"

	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	var u *model.User
	userCode, u := model.CheckUserExist(mysql.DB, req.Email)
	if userCode == code.UserExist {
		Passworderr := bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(req.Password))
		if Passworderr != nil {
			return nil, kerrors.NewGRPCBizStatusError(code.PassWordError, code.GetMsg(code.PassWordError))
		} else {
			return &user.LoginResp{
				UserId: int32(u.ID),
			}, nil
		}
	}
	return nil, kerrors.NewGRPCBizStatusError(code.UserNotExist, code.GetMsg(userCode))
}
