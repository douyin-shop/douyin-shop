package dal

import (
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
