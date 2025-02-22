package code

import "github.com/cloudwego/kitex/pkg/kerrors"

const (
	Success       = 200
	FailedPayment = 50004
)

var Message = map[int]string{
	Success:       "success",
	FailedPayment: "交易失败",
}

func GetMsg(code int) string {
	return Message[code]
}

func GetError(code int) error {
	return kerrors.NewBizStatusError(int32(code), GetMsg(code))
}
