package code

const (
	Success             = 200
	InvalidReq          = 1001
	LockError           = 2001
	AcquireLockFailed   = 2002
	StockDecreaseFailed = 2003
	InternalError       = 4003
)

var CodeMessage = map[int]string{
	Success:             "success",
	InvalidReq:          "the request isn't valid",
	LockError:           "redis lock error",
	AcquireLockFailed:   "acquire lock failed",
	StockDecreaseFailed: "stock decrease error",
	InternalError:       "internal error",
}

func GetMsg(code int) string {
	return CodeMessage[code]
}
