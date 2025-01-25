package frontend

import (
	"bytes"
	"os"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestLogin(t *testing.T) {

	// 设置当前目录为项目根目录
	err := os.Chdir("../../../")
	if err != nil {
		return
	}

	h := server.Default()
	h.POST("/login", Login)
	path := "/login"
	bodyStr := `{"email":"wyz17601402786@gmail.com","password":"123456"}`
	body := &ut.Body{Body: bytes.NewBufferString(bodyStr), Len: len(bodyStr)} // todo: you can customize body
	header := ut.Header{
		Key:   "Content-Type",
		Value: "application/json",
	} // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}
