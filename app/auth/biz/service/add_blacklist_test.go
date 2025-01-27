package service

import (
	"context"
	"testing"
	auth "github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth"
)

func TestAddBlacklist_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddBlacklistService(ctx)
	// init req and assert value

	req := &auth.AddBlackListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
