package service

import (
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/klog"
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
		return nil, err
	}

	klog.Infof("claims: %v", claims)

	userId := int(claims.Claims.(jwt.MapClaims)["user_id"].(float64))

	klog.Debug("userID:", userId)

	// 在metadata中设置用户id
	ok := metainfo.SendBackwardValue(s.ctx, "user_id", strconv.Itoa(userId))

	if claims.Valid && ok {
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
