// @Author Adrian.Wang 2025/1/26 23:16:00
package utils

import (
	"os"
	"testing"

	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/dal"
)

func TestAddPermission(t *testing.T) {
	// 设置当前目录为项目根目录
	err := os.Chdir("../../")
	if err != nil {
		return
	}
	dal.Init()
	AddPermission(mysql.DB)
	t.Log("success")
}
