package service

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/go-redis/redis/v8"
	"log"
)

type PaymentProcessor struct {
	redis  *redis.Client
	rocket *rocketmq.Producer
}

var (
	RedisClient *redis.Client
)

// ListenForPayments 监听 Redis 频道并处理支付超时
func (p *PaymentProcessor) ListenForPayments(ctx context.Context) {
	// 订阅 Redis 频道
	pubsub := p.redis.Subscribe(ctx, "paymentChannel")
	defer pubsub.Close()

	// 获取消息频道
	ch := pubsub.Channel()

	// 消费消息
	for msg := range ch {
		transactionId := msg.Payload
		log.Printf("Received transaction ID: %s", transactionId)

		// 设置延时消息，30分钟后超时
		p.sendPaymentToQueue(transactionId)

		// 在此处也可以将交易号记录到 Redis 等，等后续的支付回调进行处理
	}
}

// sendPaymentToQueue 将交易号发送到 RocketMQ
func (p *PaymentProcessor) sendPaymentToQueue(transactionId string) {
	// 创建 RocketMQ 消息
	msg := &primitive.Message{
		Topic: "PaymentTimeoutTopic", // 设置主题
		Body:  []byte(transactionId), // 消息体为交易 ID
	}
	q, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
	) //小bug，无法使用p.rocket.SendSync
	// 发送消息到 RocketMQ
	res, err := q.SendSync(context.Background(), msg)
	if err != nil {
		log.Printf("Failed to send RocketMQ message: %v", err)
		return
	}

	log.Printf("Message sent to RocketMQ with transactionId: %s, result: %v", transactionId, res)
}

// ProcessPaymentCallback 支付回调，处理支付成功
func (p *PaymentProcessor) ProcessPaymentCallback(transactionId string) error {
	// 支付回调时从 Redis 删除交易号
	err := p.redis.Del(context.Background(), transactionId)
	if err != nil {
		log.Printf("Failed to delete transactionId from Redis: %v", err)
		return fmt.Errorf("failed to process payment callback")
	}

	// 在此处可以继续处理支付成功后的业务逻辑
	log.Printf("Payment processed successfully for transactionId: %s", transactionId)

	return nil
}

func consumePaymentTimeout() {
	// 创建 RocketMQ 消费者
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"localhost:9876"}),
		consumer.WithGroupName("PaymentTimeoutGroup"),
	)
	if err != nil {
		log.Fatal("Failed to create consumer:", err)
	}
	defer c.Shutdown()

	// 订阅支付超时消息主题
	err = c.Subscribe("PaymentTimeoutTopic", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			transactionId := string(msg.Body)
			log.Printf("Payment timeout for transactionId: %s", transactionId)
			// 删除 Redis 中的记录
			err := RedisClient.Del(ctx, transactionId).Err()
			if err != nil {
				log.Printf("Failed to delete transactionId from Redis: %v", err)
			}
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		log.Fatal("Failed to subscribe:", err)
	}

	// 开始消费
	err = c.Start()
	if err != nil {
		log.Fatal("Failed to start consumer:", err)
	}
	select {}

}
