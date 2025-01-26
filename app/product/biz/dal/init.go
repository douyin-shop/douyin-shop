package dal

import (
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
