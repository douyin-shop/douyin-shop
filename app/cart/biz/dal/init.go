package dal

import (
	"github.com/douyin-shop/douyin-shop/app/cart/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
