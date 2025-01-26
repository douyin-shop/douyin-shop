package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
	"os"
	"testing"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {
	// 设置当前目录为项目根目录
	err := os.Chdir("../../")
	if err != nil {
		return
	}

	ctx := context.Background()
	s := NewDeliverTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.DeliverTokenReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	verreq := &auth.VerifyTokenReq{
		Token: resp.Token,
	}

	vers := NewVerifyTokenByRPCService(ctx)
	verresp, err := vers.Run(verreq)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", verresp)

	assert.Assert(t, verresp.Res == true)

}
