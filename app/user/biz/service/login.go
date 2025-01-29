package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/user/biz/model"
<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/douyin-shop/douyin-shop/app/user/biz/utils/code"

=======
<<<<<<< Updated upstream
	"github.com/douyin-shop/douyin-shop/app/user/code"
=======
	"github.com/douyin-shop/douyin-shop/app/user/biz/utils/code"

>>>>>>> Stashed changes
>>>>>>> ae6c4a5 (测试)
=======
	"github.com/douyin-shop/douyin-shop/app/user/code"
>>>>>>> 72916bb (测试)
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
	userCode,u:=model.CheckUserExist(req.Email)
	if userCode==code.UserExist{
	    Passworderr:=bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(u.PassWord))
	    if Passworderr!=nil{
			return nil,kerrors.NewGRPCBizStatusError(code.PassWordError,code.GetMsg(code.PassWordError))
	    }else{
			return &user.LoginResp{
				UserId: int32(u.ID),
			}, nil
		}
	}
	return nil,kerrors.NewGRPCBizStatusError(code.UserNotExist,code.GetMsg(userCode))
}
