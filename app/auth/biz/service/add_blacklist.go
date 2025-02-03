package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/redis"
	"github.com/douyin-shop/douyin-shop/app/auth/conf"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	redis_core "github.com/redis/go-redis/v9"
)

type AddBlacklistService struct {
	ctx context.Context
} // NewAddBlacklistService new AddBlacklistService
func NewAddBlacklistService(ctx context.Context) *AddBlacklistService {
	return &AddBlacklistService{ctx: ctx}
}

// Run create note info
func (s *AddBlacklistService) Run(req *auth.AddBlackListReq) (resp *auth.AddBlackListResp, err error) {
	resp = new(auth.AddBlackListResp)

	userStatusKey := fmt.Sprintf("user:%d:status", req.Blacklist.UserId)

	resp.Res = true

	// 只有黑名单时间过长才会写入数据库
	if req.Blacklist.Exp >= conf.GetConf().BlackList.MaxExpireTime {

		// 先更新数据库
		if err := model.AddOrUpdateBlackList(
			mysql.DB,
			context.Background(),
			req.Blacklist.UserId,
			req.Blacklist.Exp); err != nil {

			klog.Error("更新MySQL失败:", err)
			resp.Res = false
			err = kerrors.NewBizStatusError(502, err.Error())
			return resp, err
		}
	}

	// 删除缓存
	if err := redis.RedisClient.Del(context.Background(), userStatusKey).Err(); err != nil && !errors.Is(err, redis_core.Nil) {
		klog.Error("删除Redis缓存失败:", err)
		resp.Res = false
		err = kerrors.NewBizStatusError(502, err.Error())

		return resp, err
	}

	// 写入缓存，如果黑名单时间过长，则只在缓存中添加过期时间为最大值
	expireTime := req.Blacklist.Exp
	if req.Blacklist.Exp >= conf.GetConf().BlackList.MaxExpireTime {
		expireTime = conf.GetConf().BlackList.MaxExpireTime
	}

	// 写入缓存
	if err := redis.RedisClient.Set(context.Background(), userStatusKey, string(model.Ban), time.Duration(expireTime)*time.Second).Err(); err != nil {
		klog.Error("写入Redis缓存失败:", err)
		resp.Res = false
		err = kerrors.NewBizStatusError(502, err.Error())
		return resp, err
	}

	return

}
