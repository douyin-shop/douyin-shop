package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth/authservice"
	"github.com/douyin-shop/douyin-shop/common/nacos"
)

// VerifyTokenMiddleware 是一个中间件，用于验证token
func VerifyTokenMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 通过微服务调用user服务
		resolver := nacos.GetNacosResolver()

		// 通过微服务调用auth服务
		authService := authservice.MustNewClient("auth", client.WithResolver(resolver))
		authRes, err := authService.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{
			Token: "",
		})

		if err != nil {
			hlog.Error("authService.DeliverTokenByRPC err: ", err)
			return
		}

		if !authRes.Res {
			hlog.Error("token验证失败")
			return
		}

		c.Next(ctx)
	}
}