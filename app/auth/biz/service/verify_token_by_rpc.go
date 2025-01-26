package service

import (
	"context"
	"github.com/douyin-shop/douyin-shop/app/auth/conf"

	"github.com/cloudwego/kitex/pkg/klog"
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

	klog.Debug("userID:", claims.Claims.(jwt.MapClaims)["user_id"])

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
