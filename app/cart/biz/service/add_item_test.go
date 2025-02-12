package service

import (
	"context"
	"github.com/douyin-shop/douyin-shop/app/cart/biz/dal"
	cart "github.com/douyin-shop/douyin-shop/app/cart/kitex_gen/cart"
	"github.com/douyin-shop/douyin-shop/app/cart/rpc"
	"os"
	"testing"
)

func init() {
	// 设置当前目录为项目根目录
	err := os.Chdir("../../")
	if err != nil {
		return
	}

	dal.Init()
	rpc.InitClient()
}

func TestAddItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	req := &cart.AddItemReq{
		UserId: 1,
		Item: &cart.CartItem{
			ProductId: 1,
			Quantity:  1,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
