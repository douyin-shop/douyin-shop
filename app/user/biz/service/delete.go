package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/user/biz/model"
	"github.com/douyin-shop/douyin-shop/app/user/code"
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
)

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	var u *model.User
	userCode, u := model.CheckUserIDExist(req.UserId)
	if userCode == code.UserExist {
		err := model.DeleteUserByID(u, req.UserId)
		if err != nil {
			return nil, kerrors.NewGRPCBizStatusError(code.DeleteError, code.GetMsg(userCode))
		}

		return &user.DeleteResp{
			Success: true,
		}, nil
	}
	return nil, kerrors.NewGRPCBizStatusError(code.UserNotExist, code.GetMsg(userCode))
}
