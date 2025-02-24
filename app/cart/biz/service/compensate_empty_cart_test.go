package service

import (
	"context"
	cart "github.com/douyin-shop/douyin-shop/app/cart/kitex_gen/cart"
	"testing"
)

func TestCompensateEmptyCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCompensateEmptyCartService(ctx)
	// init req and assert value

	req := &cart.RestoreCartItemsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
