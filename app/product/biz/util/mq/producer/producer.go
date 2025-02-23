package producer

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/conf"
)

type MQProducer struct{
	producer rocketmq.Producer
}

// 创建一个新的生产者
func NewProducer() *MQProducer{
	// 使用配置文件中的NameServer创建一个新的生产者
	p,err:= rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.GetConf().RocketMQ.NameServer})),
		producer.WithRetry(2),
	)
	// 如果创建失败，打印错误信息并返回nil
	if err!=nil{
	    klog.Error("fail to create producer, err: %v")
		return nil
	}
	return &MQProducer{
		producer: p,
	}
}

// 向指定的topic发送消息
func (p *MQProducer) PutMessage(topic string, msg []byte) error{
	// 创建一个空的上下文
	ctx:=context.Background()
	// 发送消息
	_, err := p.producer.SendSync(ctx, &primitive.Message{
		// 设置消息的主题
		Topic: topic,
		// 设置消息的内容
		Body:  msg,
	})
	// 如果发送消息出错，返回错误信息
	if err != nil {
		return fmt.Errorf("send message error: %v", err)
	}
	// 发送消息成功，返回nil
	return nil
}

// 关闭MQProducer
func (p *MQProducer) Shutdown() error{
	// 调用producer的Shutdown方法
	return p.producer.Shutdown()
}