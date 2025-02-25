package service

import (
	"context"
	order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
	"testing"
)

func TestMarkOrderCanceled_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMarkOrderCanceledService(ctx)
	// init req and assert value

	req := &order.MarkOrderCanceledReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
