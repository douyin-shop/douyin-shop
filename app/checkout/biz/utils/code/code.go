package code

import "github.com/cloudwego/kitex/pkg/kerrors"

const (
	Success         = 200
	CartGetFiled    = 60001
	CartEmpty       = 60002
	ZipCodeError    = 60003
	PlaceOrderError = 60004
	EmptyCartFiled  = 60005
	PayError        = 60006
)

var Message = map[int]string{
	Success:         "success",
	CartGetFiled:    "获取购物车失败",
	CartEmpty:       "购物车为空",
	ZipCodeError:    "邮编错误",
	PlaceOrderError: "下单失败",
	EmptyCartFiled:  "清空购物车失败",
	PayError:        "支付失败",
}

func GetMsg(code int) string {
	return Message[code]
}

func GetError(code int) error {
	return kerrors.NewBizStatusError(int32(code), GetMsg(code))
}
