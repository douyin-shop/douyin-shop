package code

import "github.com/cloudwego/kitex/pkg/kerrors"

const (
	Success            = 200
	NotFoundProduct    = 40004
	StructConvertError = 40005
	ListOrderError     = 40006
	PaymentSuccess     = 40007
)

var Message = map[int]string{
	Success:            "success",
	NotFoundProduct:    "商品不存在",
	StructConvertError: "结构体转换错误",
	ListOrderError:     "获取订单列表失败",
	PaymentSuccess:     "订单已经支付成功，无法取消",
}

func GetMsg(code int) string {
	return Message[code]
}

func GetError(code int) error {
	return kerrors.NewBizStatusError(int32(code), GetMsg(code))
}
