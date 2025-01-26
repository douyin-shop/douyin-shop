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
	"regexp"
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
	enforcer.AddFunction("RegexRouterMatcher", RegexRouterMatcherFunc)
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

			return userID.(string)
		},
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return auth

}

// RegexRouterMatcher 正则路由匹配
func RegexRouterMatcher(router, regex string) bool {
	// 参数校验
	if router == "" || regex == "" {
		return false
	}

	// 编译正则表达式
	pattern, err := regexp.Compile(regex)
	if err != nil {
		log.Printf("正则表达式编译失败: %v", err)
		return false
	}

	hlog.Debug("匹配路由: ", router, " 正则表达式: ", regex)

	// 进行匹配
	return pattern.MatchString(router)
}

func RegexRouterMatcherFunc(args ...interface{}) (interface{}, error) {
	val1 := args[0].(string)
	val2 := args[1].(string)

	return RegexRouterMatcher(val1, val2), nil

}
