// @Author Adrian.Wang 2025/1/26 23:16:00
package utils

import (
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/dal"
	"os"
	"testing"
)

func TestAddPermission(t *testing.T) {
	// 设置当前目录为项目根目录
	err := os.Chdir("../../")
	if err != nil {
		return
	}
	dal.Init()
	AddPermission()
	t.Log("success")
}
