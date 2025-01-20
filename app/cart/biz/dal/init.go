package dal

import (
	"github.com/douyin-shop/douyin-shop/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
