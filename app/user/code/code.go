package code

const (
	Success          = 200
	UserNotExist     = 1001
	UserExist        = 1002
	UserCreateFailed = 1003
	PassWordError    = 1004
	DeleteError      = 1005
)

var CodeMessage = map[int]string{
	Success:          "success",
	UserNotExist:     "user not exist",
	UserExist:        "user exist",
	UserCreateFailed: "user create failed",
	PassWordError:    "password error",
	DeleteError:      "delete error",
}

func GetMsg(code int) string {
	return CodeMessage[code]
}
