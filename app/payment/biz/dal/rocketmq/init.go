// Package rocketmq @Author Adrian.Wang 2025/2/22 12:17:00
package rocketmq

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/douyin-shop/douyin-shop/app/payment/conf"
	"github.com/douyin-shop/douyin-shop/common/topic"
)

var (
	RocketProducer     rocketmq.Producer
	RocketPushConsumer rocketmq.PushConsumer
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

	RocketPushConsumer, _ = rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{conf.GetConf().RocketMQ.Address}), // 接入点地址
		consumer.WithGroupName(conf.GetConf().RocketMQ.GroupName),          // 分组名称
	)

	err = RocketPushConsumer.Subscribe(topic.GetMsg(topic.Payment), consumer.MessageSelector{}, PaymentTimeout)
	if err != nil {
		fmt.Printf("Subscribe error: %s\n", err)
		return
	}
	// 启动消费者
	err = RocketPushConsumer.Start()
	if err != nil {
		fmt.Printf("Start consumer error: %s\n", err)
		return
	}

}
