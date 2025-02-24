package code

import "github.com/cloudwego/kitex/pkg/kerrors"

const (
	Success       = 200
	FailedPayment = 50004
	Overtime      = 50005
	Queueerr      = 50006
)

var Message = map[int]string{
	Success:       "success",
	FailedPayment: "交易失败",
	Overtime:      "交易超时",
	Queueerr:      "消息队列错误",
}

func GetMsg(code int) string {
	return Message[code]
}

func GetError(code int) error {
	return kerrors.NewBizStatusError(int32(code), GetMsg(code))
}
