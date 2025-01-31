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

func TestGetCart(t *testing.T) {
	h := server.Default()
	h.GET("cart/get_cart", GetCart)
	path := "cart/get_cart"                                   // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	w := ut.PerformRequest(h.Engine, "GET", path, body, ut.Header{
		Key:   "Content-Type",
		Value: "application/json",
	}, ut.Header{
		Key:   "Authorization",
		Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc5MDU4MjMsInVzZXJfaWQiOjF9.UFoL-XTHXfpT1ZVN1UcU-joB13oRMTH-5g-zfIAXYao",
	})
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestRegister(t *testing.T) {
	h := server.Default()
	h.POST("/register", Register)
	path := "/register"                                       // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}
