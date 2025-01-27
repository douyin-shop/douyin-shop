package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/redis"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
)

type DeleteBlacklistService struct {
	ctx context.Context
} // NewDeleteBlacklistService new DeleteBlacklistService
func NewDeleteBlacklistService(ctx context.Context) *DeleteBlacklistService {
	return &DeleteBlacklistService{ctx: ctx}
}

// Run create note info
func (s *DeleteBlacklistService) Run(req *auth.DeleteBlackListReq) (resp *auth.DeleteBlackListResp, err error) {
	resp = new(auth.DeleteBlackListResp)
	userId := req.UserId

	// 1. 删除Redis中的用户状态
	userStatusKey := fmt.Sprintf("user:%d:status", userId)
	result, err := redis.RedisClient.Del(context.Background(), userStatusKey).Result()
	if err != nil {
		klog.Errorf("删除Redis缓存失败: %v", err)
		// Redis删除失败不影响主流程
	}
	klog.Debug("删除Redis缓存成功: ", result)

	// 2. 删除MySQL中的黑名单记录
	if err := model.DeleteFromBlackList(mysql.DB, context.Background(), userId); err != nil {
		klog.Errorf("从黑名单删除用户失败: %v", err)
		resp.Res = false
		return resp, err
	}

	resp.Res = true
	return resp, nil

	return
}
