package mq

import (
	"github.com/douyin-shop/douyin-shop/app/product/conf"
	"github.com/douyin-shop/douyin-shop/app/product/biz/util/mq/customer"
	"github.com/douyin-shop/douyin-shop/app/product/biz/util/mq/producer"
	"github.com/douyin-shop/douyin-shop/app/product/biz/util/mysql"
)

var (
	c *customer.MQConsumer
	p *producer.MQProducer
)

// InitProducer 初始化生产者并启动监听
func InitProducer() {
	p = producer.NewProducer()
	// 不要在这里使用 defer p.Shutdown(), 因为这会导致在函数返回时立即关闭生产者
	mysql.RunListen(p)
}

// InitConsumer 初始化消费者并启动消费
func InitConsumer() {
	var err error
	c, err = customer.NewConsumer()
	if err != nil {
		// 处理错误，例如记录日志或退出程序
		panic(err)
	}
	// 同样不要在这里使用 defer c.Shutdown()
	topic := conf.GetConf().RocketMQ.Topic
	c.Start(topic)
}

// ShutdownProducer 关闭生产者
func ShutdownProducer() {
	if p != nil {
		p.Shutdown()
	}
}

// ShutdownConsumer 关闭消费者
func ShutdownConsumer() {
	if c != nil {
		c.Shutdown()
	}
}

// InitMq 初始化 MQ 生产者和消费者，并在后台运行它们
func InitMq() {
	go func() {
		InitProducer()
	}()

	go func() {
		InitConsumer()
	}()
}

// ShutdownMq 关闭 MQ 生产者和消费者
func ShutdownMq() {
	ShutdownProducer()
	ShutdownConsumer()
}