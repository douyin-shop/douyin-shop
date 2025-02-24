// @Author Adrian.Wang 2025/2/22 12:42:00
package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

// SendMessage 发送消息
func SendMessage(topic, body string) (result *primitive.SendResult, err error) {
	// 发送的消息
	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(body),
	}
	// 发送消息
	result, err = RocketProducer.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("Send message error: %s\n", err)
		return nil, err
	}

	return result, nil
}

// SendDelayMessage 发送延时消息
func SendDelayMessage(topic, body string, delayLevel int) (result *primitive.SendResult, err error) {
	// 发送的消息
	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte(body),
	}
	// 设置延时级别
	// 延迟级别：1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h
	// 延迟30分钟
	msg.WithDelayTimeLevel(delayLevel)
	// 发送消息
	result, err = RocketProducer.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("Send message error: %s\n", err)
		return nil, err
	}

	return result, nil
}
