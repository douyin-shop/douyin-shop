package service

import (
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/redis"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/utils"
	"github.com/douyin-shop/douyin-shop/app/auth/conf"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/golang-jwt/jwt"
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
		return nil, err
	}

	klog.Infof("claims: %v", claims)

	userId := int32(claims.Claims.(jwt.MapClaims)["user_id"].(float64))

	// 从Redis中取出token，与当前token比较，如果不同就直接返回
	tokenId := utils.GenerateTokenKey(userId)
	result, err := redis.RedisClient.Get(context.Background(), tokenId).Result()

	if err != nil {
		klog.Error("redis get error: ", err)
		return nil, err
	}

	if result == "" || result != token {
		klog.Info("token检验不通过，可能已经过期！！！", result, token, "用户id：", userId)
		resp = &auth.VerifyResp{
			Res: false,
		}
		return
	}

	klog.Debug("userID:", userId)

	// 在metadata中设置用户id
	ok := metainfo.SendBackwardValue(s.ctx, "user_id", string(userId))

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
