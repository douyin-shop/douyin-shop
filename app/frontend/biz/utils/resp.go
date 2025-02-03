package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {

	if err == nil {
		var resp = map[string]interface{}{
			"code": -1,
			"data": nil,
			"msg":  "unknown error",
		}
		c.JSON(code, resp)
		return
	}

	bizErr, ok := kerrors.FromBizStatusError(err)

	msg := err.Error()

	if ok {
		msg = bizErr.BizMessage()
	}

	var resp = map[string]interface{}{
		"code": -1,
		"data": nil,
		"msg":  msg,
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
