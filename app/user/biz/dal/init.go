package dal

import (
	"github.com/douyin-shop/douyin-shop/app/user/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
