package service

import (
	"context"

	"github.com/douyin-shop/douyin-shop/app/user/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/user/biz/dal/mysql"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/user/biz/utils/code"
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
	"gorm.io/gorm"
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
	//构建一个user对象字段的map
	userMap := make(map[string]interface{})
	userMap["phone"] = req.Phone
	userMap["email"] = req.Email
	userMap["address"] = req.Address
	userMap["nickname"] = req.Nickname
	//根据id查询用户,判断用户是否存在
	userModel := &model.User{}
	err = mysql.DB.WithContext(s.ctx).Where("id = ?", req.UserId).First(userModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, kerrors.NewBizStatusError(code.UserNotExist, code.GetMsg(code.UserNotExist))
		}
		return nil, err
	}
	//更新用户信息
	err = userModel.UpdateUserInfo(mysql.DB, userMap)
	if err != nil {
		return nil, err
	}
	return &user.UpdateResp{}, nil
}
