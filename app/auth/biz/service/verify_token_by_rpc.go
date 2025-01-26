package service

import (
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/redis"
	"github.com/douyin-shop/douyin-shop/app/auth/conf"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/golang-jwt/jwt"
	"strconv"
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

	// TODO  根据userId获取用户状态，如果用户已经进入黑名单，则直接返回false

	// 使用Redis Set检查用户是否在黑名单中
	blackListKey := "user:blacklist" // 存储所有黑名单用户ID的Set key
	isMember, err := redis.RedisClient.SIsMember(context.Background(), blackListKey, userId).Result()

	// Redis存在错误，返回报错
	if err != nil {
		klog.Error("redis SISMEMBER error: ", err)
		return nil, err
	}

	// isMember为true说明用户在黑名单中，返回false，表示用户被拒绝
	if isMember {
		return &auth.VerifyResp{
			Res: false,
		}, nil
	}

	// 将userId从int32 转化成string

	// 在metadata中设置用户id
	ok := metainfo.SendBackwardValue(s.ctx, "user_id", strconv.Itoa(int(userId)))

	if !ok {
		klog.Error("set user_id in metadata error")
	}

	if claims.Valid {
		resp = &auth.VerifyResp{
			Res: true,
		}
	} else {
		resp = &auth.VerifyResp{
			Res: false,
		}
	}

	return
}
