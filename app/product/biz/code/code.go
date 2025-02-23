package code

const (
	Success = 200
	Error   = 500
	ProductExist= 1001
	ProductNotExist= 1002
	CategoryExist= 2001
	CategoryNotExist= 2002
)

var CodeMsg = map[int]string{
	Success: "success",
	Error:   "error",
	ProductExist: "product exist",
	ProductNotExist: "product not exist",
	CategoryExist: "category exist",
	CategoryNotExist: "category not exist",
}

// CodeMsgMap 获取CodeMsg
func GetMessage(code int) string {
	return CodeMsg[code]
}