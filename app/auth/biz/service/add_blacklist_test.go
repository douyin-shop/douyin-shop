package service

import (
	"context"
	"os"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
)

func init() {
	// 设置当前目录为项目根目录
	err := os.Chdir("../../")
	if err != nil {
		return
	}

	dal.Init()
}

func TestAddBlacklist_Run(t *testing.T) {

	ctx := context.Background()
	s := NewAddBlacklistService(ctx)
	// init req and assert value

	req := &auth.AddBlackListReq{
		Blacklist: &auth.Blacklist{
			UserId: 1,
			Exp:    86401,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	assert.Assert(t, err == nil && resp.Res == true)

	// todo: edit your unit test

}
