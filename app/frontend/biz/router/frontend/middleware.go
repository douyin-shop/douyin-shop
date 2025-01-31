// Code generated by hertz generator.

package frontend

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _cartMw() []app.HandlerFunc {

	return []app.HandlerFunc{
		middleware.VerifyTokenMiddleware(),
		middleware.GetCasbinMiddleware().RequiresPermissions("/cart:*"),
	}
}

func _getcartMw() []app.HandlerFunc {

	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}
