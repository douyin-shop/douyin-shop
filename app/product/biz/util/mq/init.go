package mq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/biz/util/mq/customer"
	"github.com/douyin-shop/douyin-shop/app/product/biz/util/mq/producer"
	"github.com/douyin-shop/douyin-shop/app/product/biz/util/mysql"
	"github.com/douyin-shop/douyin-shop/app/product/conf"
)

var (
	c *customer.MQConsumer
	p *producer.MQProducer
)

// InitMq 初始化 MQ 生产者和消费者，并在后台运行它们
func InitMq() {

	InitProducer()
	InitConsumer()
	mysql.RunListen(p)
}

// InitProducer 初始化生产者并启动监听
func InitProducer() {
	p = producer.NewProducer()
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
	if err := c.Start(topic); err != nil {
		klog.Error("fail to start consumer, err: %v", err)
		// 处理错误，例如记录日志或退出程序
		panic(err)
	}
	klog.Debug("MQ 消费者启动成功")
}

// ShutdownProducer 关闭生产者
func ShutdownProducer() {
	if p != nil {
		if err := p.Shutdown(); err != nil {
			klog.Error("fail to shutdown producer, err: %v", err)
			return
		}
	}
}

// ShutdownConsumer 关闭消费者
func ShutdownConsumer() {
	if c != nil {
		c.Shutdown()
	}
}

// ShutdownMq 关闭 MQ 生产者和消费者
func ShutdownMq() {
	ShutdownProducer()
	ShutdownConsumer()
}
