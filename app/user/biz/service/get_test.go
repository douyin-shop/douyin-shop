package service

import (
	"context"
	"testing"
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
)

func TestGet_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetService(ctx)
	// init req and assert value

	req := &user.GetReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
