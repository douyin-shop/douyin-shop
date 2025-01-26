// Package middleware @Author Adrian.Wang 2025/1/26 20:26:00
package middleware

import (
	"context"
	"fmt"
	casbin_sdk "github.com/casbin/casbin/v2"
	casbin_model "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/frontend/conf"
	"github.com/hertz-contrib/casbin"
	"log"
)

// CasbinMiddleware casbin中间件
func CasbinMiddleware() *casbin.Middleware {
	casbinAdapter, err := gormadapter.NewAdapterByDB(mysql.DB)
	if err != nil {
		hlog.Fatal(err)
		return nil
	}
	model, err := casbin_model.NewModelFromFile(fmt.Sprintf("./conf/%s/model.conf", conf.GetConf().Env))
	enforcer, err := casbin_sdk.NewCachedEnforcer(model, casbinAdapter)
	if err != nil {
		hlog.Fatal(err)
		return nil
	}

	auth, err := casbin.NewCasbinMiddlewareFromEnforcer(enforcer,
		func(ctx context.Context, c *app.RequestContext) string {
			// 从ctx中获取user_id
			userID := ctx.Value("user_id")

			// 如果user_id不是string类型
			if _, ok := userID.(string); !ok {
				return ""
			}
			hlog.Debug("user_id: ", userID.(string))
			fmt.Println("user_id", userID.(string))

			return userID.(string)
		},
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return auth

}
