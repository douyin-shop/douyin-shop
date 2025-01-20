package dal

import (
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
