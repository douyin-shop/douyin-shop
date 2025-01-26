package dal

import (
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
