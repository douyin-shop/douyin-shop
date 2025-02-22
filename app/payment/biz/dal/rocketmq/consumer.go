// Package rocketmq @Author Adrian.Wang 2025/2/22 15:27:00
package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Subscribe(topic string, callback func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error)) error {
	// 订阅消息
	err := RocketPushConsumer.Subscribe(topic, consumer.MessageSelector{}, callback)
	if err != nil {
		fmt.Printf("Subscribe error: %s\n", err)
		return err
	}
	// 启动消费者
	err = RocketPushConsumer.Start()
	if err != nil {
		fmt.Printf("Start consumer error: %s\n", err)
		return err
	}

	return nil
}

func PaymentTimeout(ctx context.Context, messages ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

	klog.Debug("接收到PaymentTimeout消息")
	for _, msg := range messages {
		fmt.Println(string(msg.Body))
		// 解析msg为结构体map[string]string
		messageJson := string(msg.Body)

		messageMap := make(map[string]string)

		err := json.Unmarshal([]byte(messageJson), &messageMap)

		if err != nil {
			klog.Error("json.Unmarshal error: ", err)
			return consumer.ConsumeRetryLater, err
		}

		// 获取订单号和交易号
		orderId := messageMap["order_id"]
		transactionId := messageMap["transaction_id"]

		// TODO 通知订单服务取消订单
		klog.Debug("订单号: ", orderId, " 交易号: ", transactionId)

		klog.Debugf("messageMap：%v", messageMap)

	}

	return consumer.ConsumeSuccess, nil

}
