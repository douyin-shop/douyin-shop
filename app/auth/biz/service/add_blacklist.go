package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/redis"
	"github.com/douyin-shop/douyin-shop/app/auth/conf"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"gorm.io/gorm"
	"time"
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
	newExpireTime := req.Blacklist.Exp

	// 1. 获取Redis中的状态和过期时间
	userStatus, err := redis.RedisClient.Get(context.Background(), userStatusKey).Result()
	redisTTL, err := redis.RedisClient.TTL(context.Background(), userStatusKey).Result()

	klog.Debug("userStatus: ", userStatus)
	klog.Debug("redisTTL: ", redisTTL)

	// 2. 获取MySQL中的记录
	_, mysqlExpire, err := model.GetUserStatusFromDB(mysql.DB, context.Background(), req.Blacklist.UserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		klog.Error("查询MySQL记录失败:", err)
		resp.Res = false
		return resp, err
	}

	// 3. 判断是否需要更新Redis
	needUpdateRedis := false
	if err == nil && model.UserStatus(userStatus) == model.Ban {
		// Redis中存在记录，比较过期时间
		if newExpireTime > int64(redisTTL.Seconds()) {
			needUpdateRedis = true
		}
	} else {
		needUpdateRedis = true
	}

	klog.Debug("needUpdateRedis: ", needUpdateRedis)

	// 4. 判断是否需要更新MySQL
	needUpdateMySQL := false
	if newExpireTime > conf.GetConf().BlackList.MaxExpireTime {
		if mysqlExpire == 0 || newExpireTime > mysqlExpire {
			needUpdateMySQL = true
		}
	}

	klog.Debug("needUpdateMySQL: ", needUpdateMySQL)

	// 5. 更新Redis
	if needUpdateRedis {
		expireTime := time.Duration(newExpireTime) * time.Second
		if expireTime > time.Duration(conf.GetConf().BlackList.MaxExpireTime)*time.Second {
			expireTime = time.Duration(conf.GetConf().BlackList.MaxExpireTime) * time.Second
		}

		_, err = redis.RedisClient.Set(
			context.Background(),
			userStatusKey,
			string(model.Ban),
			expireTime,
		).Result()
		if err != nil {
			klog.Error("更新Redis失败:", err)
		}
	}

	// 6. 更新MySQL
	if needUpdateMySQL {
		if err := model.AddUserToBlackList(mysql.DB, context.Background(), req.Blacklist.UserId, newExpireTime); err != nil {
			klog.Error("更新MySQL失败:", err)
			resp.Res = false
			return resp, err
		}
	}

	resp.Res = true
	return resp, nil
}
