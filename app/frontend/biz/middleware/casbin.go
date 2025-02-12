// Package middleware @Author Adrian.Wang 2025/1/26 20:26:00
package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/utils"
	"github.com/hertz-contrib/casbin"
	"log"
	"regexp"
	"sync"
)

var (
	casbinMiddleware *casbin.Middleware
	once             sync.Once
)

// GetCasbinMiddleware 获取casbin中间件单例
func GetCasbinMiddleware() *casbin.Middleware {
	once.Do(func() {
		casbinMiddleware = initCasbinMiddleware()
	})
	return casbinMiddleware
}

// CasbinMiddleware casbin中间件
func initCasbinMiddleware() *casbin.Middleware {

	// 注册自定义函数
	utils.GetCachedEnforcer().AddFunction("RegexRouterMatcher", RegexRouterMatcherFunc)

	auth, err := casbin.NewCasbinMiddlewareFromEnforcer(utils.GetCachedEnforcer(),
		func(ctx context.Context, c *app.RequestContext) string {
			// 从ctx中获取user_id
			userID := ctx.Value("user_id")

			// 如果user_id不是string类型, 则返回空字符串
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
