// Package rocketmq @Author Adrian.Wang 2025/2/22 12:17:00
package rocketmq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/douyin-shop/douyin-shop/app/payment/conf"
)

var (
	RocketProducer rocketmq.Producer
)

func Init() {

	RocketProducer, _ = rocketmq.NewProducer(
		producer.WithNameServer([]string{conf.GetConf().RocketMQ.Address}), // 接入点地址
		producer.WithRetry(conf.GetConf().RocketMQ.RetryTimes),             // 重试次数
		producer.WithGroupName(conf.GetConf().RocketMQ.GroupName),          // 分组名称
	)
	err := RocketProducer.Start()
	if err != nil {
		return
	}

}
