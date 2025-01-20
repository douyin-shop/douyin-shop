package dal

import (
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
