package code

import "github.com/cloudwego/kitex/pkg/kerrors"

const (
	Success          = 200
	Error            = 500
	ProductExist     = 1001
	ProductNotExist  = 1002
	AddProductError  = 1003
	CategoryExist    = 2001
	CategoryNotExist = 2002
	UploadFileError  = 3001
	ESSearchError    = 4001
)

var CodeMsg = map[int]string{
	Success:          "success",
	Error:            "error",
	ProductExist:     "product exist",
	ProductNotExist:  "product not exist",
	AddProductError:  "add product error",
	CategoryExist:    "category exist",
	CategoryNotExist: "category not exist",
	UploadFileError:  "upload file error",
	ESSearchError:    "es search error",
}

// CodeMsgMap 获取CodeMsg
func GetMessage(code int) string {
	return CodeMsg[code]
}

func GetErr(code int) error {
	return kerrors.NewGRPCBizStatusError(int32(code), GetMessage(code))
}
