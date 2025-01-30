package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/user/biz/model"
	"github.com/douyin-shop/douyin-shop/app/user/code"
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
)

type UpdateService struct {
	ctx context.Context
} // NewUpdateService new UpdateService
func NewUpdateService(ctx context.Context) *UpdateService {
	return &UpdateService{ctx: ctx}
}

// Run create note info
func (s *UpdateService) Run(req *user.UpdateReq) (resp *user.UpdateResp, err error) {
	// Finish your business logic.
	var u *model.User
	userCode, u := model.CheckUserExist(req.Email)
	if userCode == code.UserExist {
		err := model.UpdateUser(u)
		if err != nil {
			return nil, kerrors.NewGRPCBizStatusError(code.UpdateError, code.GetMsg(code.UpdateError))
		}
		return &user.UpdateResp{
			Success: true,
		}, nil
	}
	return nil, kerrors.NewGRPCBizStatusError(code.UserNotExist, code.GetMsg(userCode))

}
