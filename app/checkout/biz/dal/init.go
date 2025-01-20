package dal

import (
	"github.com/douyin-shop/douyin-shop/app/checkout/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
