package service

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/redis"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/utils"
	"github.com/douyin-shop/douyin-shop/app/auth/conf"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/golang-jwt/jwt"
)

type LogoutService struct {
	ctx context.Context
} // NewLogoutService new LogoutService
func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

// Run create note info
func (s *LogoutService) Run(req *auth.LogoutReq) (resp *auth.LogoutResp, err error) {

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

	// 使用Redis String存储用户状态
	userStatusKey := utils.GenerateTokenKey(userId)

	// 从redis中删除token
	n, err := redis.RedisClient.Del(context.Background(), userStatusKey).Result()

	// Redis删除了0个key，说明token不存在
	if n == 0 {
		klog.Error("用户未登录")
		return nil, errors.New("用户未登录")
	}

	if err != nil {
		return nil, err
	}

	resp = &auth.LogoutResp{
		Success: true,
	}

	return
}
