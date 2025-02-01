package middleware

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"
)

// VerifyTokenMiddleware 是一个中间件，用于验证token
func VerifyTokenMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		// 从请求的上下文中获取token
		token := string(c.GetHeader("Authorization"))

		hlog.Debug("token: ", token)

		// 通过微服务调用auth服务来验证jwt，并且会将user_id解析出来放到ctx中
		ctx = metainfo.WithBackwardValues(ctx)
		authRes, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{
			Token: token,
		})

		if err != nil {
			hlog.Error("authService.DeliverTokenByRPC err: ", err)
			return
		}

		// 获取服务端返回的user_id
		userId, ok := metainfo.RecvBackwardValue(ctx, "user_id") // 获取服务端传回的元数据
		if !ok || !authRes.Res {
			hlog.Error("token验证失败")
			return
		}

		ctx = context.WithValue(ctx, "user_id", userId)

		c.Next(ctx)
	}
}
