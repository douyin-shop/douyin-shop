package code

const (
	Success             = 200
	InvalidRequest      = 1000
	LockError           = 1001
	AcquireLockFailed   = 1002
	StockDecreaseFailed = 1003
	CreateOrderError    = 1004
	DecreaseStockError  = 1005
	UnLockError         = 1006
	InternalError       = 4003
)

var CodeMessage = map[int]string{
	Success:             "success",
	InvalidRequest:      "the request isn't valid",
	LockError:           "redis lock error",
	AcquireLockFailed:   "acquire lock failed",
	StockDecreaseFailed: "stock decrease error",
	CreateOrderError:    "create order error",
	DecreaseStockError:  "decrease stock error",
	UnLockError:         "code error",
	InternalError:       "internal error",
}

func GetMsg(code int) string {
	return CodeMessage[code]
}
