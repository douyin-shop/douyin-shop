// Package utils @Author Adrian.Wang 2025/1/26 23:10:00
package utils

import (
	"fmt"
	casbin_sdk "github.com/casbin/casbin/v2"
	casbin_model "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/frontend/conf"
)

func AddPermission() {
	casbinAdapter, err := gormadapter.NewAdapterByDB(mysql.DB)
	if err != nil {
		hlog.Fatal(err)
	}
	model, err := casbin_model.NewModelFromFile(fmt.Sprintf("./conf/%s/model.conf", conf.GetConf().Env))

	enforcer, err := casbin_sdk.NewCachedEnforcer(model, casbinAdapter)

	enforcer.AddPolicy("user", "^/cart.*$", ".*")
}
