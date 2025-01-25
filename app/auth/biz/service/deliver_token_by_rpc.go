package service

import (
	"context"
	"fmt"
	"time"

	"github.com/douyin-shop/douyin-shop/app/auth/conf"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/golang-jwt/jwt"
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

	// 创建 token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString([]byte(conf.GetConf().Jwt.Secret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %w", err)
	}

	resp = &auth.DeliveryResp{
		Token: tokenString,
	}

	return
}
