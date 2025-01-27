package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {

	var resp = map[string]interface{}{
		"code": -1,
		"data": nil,
		"msg":  err.Error(),
	}
	c.JSON(code, resp)
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {

	var resp = map[string]interface{}{
		"code": 0,
		"data": data,
		"msg":  "success",
	}
	c.JSON(code, resp)
}
