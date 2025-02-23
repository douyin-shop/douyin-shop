package dal

import (
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/redis"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/rocketmq"
)

func Init() {
	redis.Init()
	mysql.Init()
	rocketmq.Init()
}
