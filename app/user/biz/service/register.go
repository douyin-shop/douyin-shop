package service

import (
	"context"

<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/user/biz/model"
	"github.com/douyin-shop/douyin-shop/app/user/biz/utils/code"
=======
<<<<<<< Updated upstream
	"github.com/douyin-shop/douyin-shop/app/user/biz/model"
	"github.com/douyin-shop/douyin-shop/app/user/code"
=======
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/user/biz/model"
	"github.com/douyin-shop/douyin-shop/app/user/biz/utils/code"
>>>>>>> Stashed changes
>>>>>>> ae6c4a5 (测试)
=======
	"github.com/douyin-shop/douyin-shop/app/user/biz/model"
	"github.com/douyin-shop/douyin-shop/app/user/code"
>>>>>>> 72916bb (测试)
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type RegisterService struct {
	ctx context.Context

} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	userCode,_:=model.CheckUserExist(req.Email)
	if userCode==code.UserExist{
		return nil,kerrors.NewGRPCBizStatusError(code.UserExist,code.GetMsg(code.UserExist))
	}
	newUser:=&model.User{
		Email:req.Email,
		PassWord: req.Password,
	}
	Id,err:=model.CreateUser(newUser)
	if err!=nil{
		return nil,kerrors.NewGRPCBizStatusError(code.UserCreateFailed,code.GetMsg(code.UserCreateFailed))
	}
	return  &user.RegisterResp{
		UserId:int32(Id) ,
	},nil
}
