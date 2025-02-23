// @Author Adrian.Wang 2025/2/22 12:55:00
package topic

const (
	Success = 1
	Payment = 2
)

var Message = map[int]string{
	Success: "success",
	Payment: "payment",
}

func GetMsg(code int) string {
	return Message[code]
}
