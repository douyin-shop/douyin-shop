// Package utils @Author Adrian.Wang 2025/1/26 23:10:00
package utils

import (
	"fmt"

	casbin_sdk "github.com/casbin/casbin/v2"
	casbin_model "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

var (
	cachedEnforcer *casbin_sdk.CachedEnforcer
)

// InitCasbin 初始化casbin
func InitCasbin(db *gorm.DB, modelFile string) error {
	casbinAdapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		hlog.Fatal(err)
	}
	model, err := casbin_model.NewModelFromFile(modelFile)

	if err != nil {
		hlog.Fatal(err)
	}

	cachedEnforcer, err = casbin_sdk.NewCachedEnforcer(model, casbinAdapter)

	if err != nil {
		hlog.Fatal(err)
	}

	cachedEnforcer.EnableAutoSave(true)

	return err
}

// GetCachedEnforcer 获取缓存的enforcer
func GetCachedEnforcer() *casbin_sdk.CachedEnforcer {
	return cachedEnforcer
}

// AddPermission 为角色添加权限
// Deprecated: 这里只是作为测试使用
func AddPermission(db *gorm.DB) {

	cachedEnforcer.AddPolicy("user", "^/cart.*$", ".*")

	cachedEnforcer.InvalidateCache()

}

// 为用户添加角色
func AddRoleToUser(db *gorm.DB, userId int32, role string) {

	res, err := cachedEnforcer.AddRoleForUser(fmt.Sprintf("%d", userId), role)
	if err != nil {
		hlog.Fatal(err)
		return
	}

	// 清空缓存
	err = cachedEnforcer.InvalidateCache()
	if err != nil {
		hlog.Fatal(err)
	}

	hlog.Debug("AddRoleForUser res: ", res)

}
