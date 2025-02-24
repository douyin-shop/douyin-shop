// Package code @Author Adrian.Wang 2025/2/23 13:18:00
package code

import "errors"

const (
	Success         = 200  // 成功
	ImageRequired   = 1001 // 需要上传图片
	ImageOpenFailed = 1002 // 打开图片失败
	ImageUploadFail = 1003 // 上传图片失败
)

var CodeMessage = map[int]string{
	Success:         "success",
	ImageRequired:   "需要上传图片！",
	ImageOpenFailed: "图片存在问题！",
	ImageUploadFail: "上传图片失败！",
}

func GetMsg(code int) string {
	return CodeMessage[code]
}

func GetErr(code int) error {
	return errors.New(CodeMessage[code])
}
