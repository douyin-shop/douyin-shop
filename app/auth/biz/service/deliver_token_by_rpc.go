package service

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/redis"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/utils"
	"time"

	"github.com/douyin-shop/douyin-shop/app/auth/conf"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/golang-jwt/jwt"
	redis_core "github.com/redis/go-redis/v9"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {

	userId := req.UserId

	result, err := redis.RedisClient.Get(context.Background(), utils.GenerateTokenKey(userId)).Result()

	// 如果redis中有token，说明用户已经登录过，直接返回
	if err == nil && result != "" {
		err = kerrors.NewBizStatusError(400, "用户已登陆")
		return nil, err
	}

	// 如果redis中有错误，直接返回
	if err != nil && !errors.Is(err, redis_core.Nil) {
		klog.Error("redis get error: ", err)
		err = kerrors.NewBizStatusError(502, err.Error())
		return nil, err
	}

	// 创建 token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString([]byte(conf.GetConf().Jwt.Secret))
	if err != nil {
		err = kerrors.NewBizStatusError(502, err.Error())

		return nil, err
	}

	// 将token存入redis
	err = redis.RedisClient.Set(context.Background(), utils.GenerateTokenKey(userId), tokenString, time.Hour*3).Err()

	if err != nil {
		klog.Error("redis set error: ", err)
		err = kerrors.NewBizStatusError(502, err.Error())

		return nil, err
	}

	resp = &auth.DeliveryResp{
		Token: tokenString,
	}

	return
}
