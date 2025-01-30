package service

import (
	"context"
	"github.com/douyin-shop/douyin-shop/app/user/biz/dal"
	user "github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user"
	"testing"
)

func TestDelete_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewDeleteService(ctx)
	// init req and assert value

	req := &user.DeleteReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
