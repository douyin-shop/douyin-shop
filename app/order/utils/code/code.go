package code

import "github.com/cloudwego/kitex/pkg/kerrors"

const (
	Success            = 200
	NotFoundProduct    = 40004
	StructConvertError = 40005
)

var Message = map[int]string{
	Success:            "success",
	NotFoundProduct:    "商品不存在",
	StructConvertError: "结构体转换错误",
}

func GetMsg(code int) string {
	return Message[code]
}

func GetError(code int) error {
	return kerrors.NewBizStatusError(int32(code), GetMsg(code))
}
