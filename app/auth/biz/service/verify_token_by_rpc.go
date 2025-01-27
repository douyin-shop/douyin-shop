package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/redis"
	"github.com/douyin-shop/douyin-shop/app/auth/conf"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/dal/mysql"
	"github.com/golang-jwt/jwt"
	redis_core "github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	token := req.Token

	// 从token中解析出用户id
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.GetConf().Jwt.Secret), nil
	})

	if err != nil {
		klog.Error("parse token error: ", err)
		return nil, err
	}

	klog.Infof("claims: %v", claims)

	userId := int32(claims.Claims.(jwt.MapClaims)["user_id"].(float64))

	klog.Debug("token验证通过，UserId:", userId)

	// TODO 根据userId获取用户状态，如果用户已经进入黑名单，则直接返回false
	// TODO 先从Redis中获取当前用户状态，如果状态为黑名单，则直接返回false
	// TODO 如果Redis中没有当前用户状态，则从数据库中获取用户状态，并且存入Redis中
	// TODO 如果用户状态为黑名单，则直接返回false
	// TODO 如果用户状态不为黑名单，则返回true
	// TODO 在将用户加入黑名单的时候，需要将用户状态存入Redis中，修改其状态从正常到黑名单

	// 使用Redis String存储用户状态
	userStatusKey := fmt.Sprintf("user:%d:status", userId)
	// 从Redis中获取用户状态
	userStatus, err := redis.RedisClient.Get(context.Background(), userStatusKey).Result()

	// Redis存在错误，返回报错
	if err != nil && !errors.Is(err, redis_core.Nil) {
		klog.Error("redis Get error: ", err)
		return nil, err
	}

	// 如果Redis中无用户状态，从数据库中获取用户状态
	if err != nil && errors.Is(err, redis_core.Nil) {
		// 从数据库中获取用户状态
		userStatus, expire, err := model.GetUserStatusFromDB(mysql.DB, context.Background(), userId)
		if err != nil {
			klog.Error("GetUserStatusFromDB error: ", err)
			return nil, err
		}

		// expire过期时间最多为24h，以减少Redis中的数据量
		if expire > 24*3600 {
			expire = 24 * 3600
		}

		if userStatus == model.Ban {
			klog.Infof("用户id: %d 在黑名单中,阻止请求！", userId)
			// 最多
			redis.RedisClient.Set(context.Background(), userStatusKey, model.Ban, time.Duration(expire)*time.Second)
			return &auth.VerifyResp{
				Res: false,
			}, nil

		}

		// 存储用户状态到Redis中,过期时间为24h
		_, err = redis.RedisClient.Set(context.Background(), userStatusKey, model.Normal, time.Hour*24).Result()
		if err != nil {
			klog.Error("redis Set error: ", err)
			return nil, err
		}
	}

	// 如果Redis里面有用户状态，直接判断用户状态
	if model.UserStatus(userStatus) == model.Ban {
		klog.Infof("用户id: %d 在黑名单中,阻止请求！", userId)
		return &auth.VerifyResp{
			Res: false,
		}, nil
	}

	// 在metadata中设置用户id,以供调用链使用
	ok := metainfo.SendBackwardValue(s.ctx, "user_id", strconv.Itoa(int(userId)))
	if !ok {
		klog.Error("set user_id in metadata error")
	}

	// 返回true，表示用户通过验证
	if claims.Valid {
		resp = &auth.VerifyResp{
			Res: true,
		}
		return
	}

	resp = &auth.VerifyResp{
		Res: false,
	}
	return

}
