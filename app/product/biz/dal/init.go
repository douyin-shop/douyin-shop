package dal

import (
	mysql2 "github.com/douyin-shop/douyin-shop/app/product/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql2.Init()
}
