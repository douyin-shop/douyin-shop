package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

const maxRetries = 3 // 最大重试次数

// 生产者发送扣减库存失败消息到Kafka队列
func SendInventoryFailedMessage(producer sarama.SyncProducer, productId uint32, quantity int32) error {
	msg := InventoryMessage{
		ProductId: productId,
		Quantity:  quantity,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: InventoryFailedTopic,
		Value: sarama.ByteEncoder(msgBytes),
	})
	return err
}

// 生产者发送扣减库存消息到Kafka队列
func SendInventoryMessage(producer sarama.SyncProducer, productId uint32, quantity int32) error {
	msg := InventoryMessage{
		ProductId: productId,
		Quantity:  quantity,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// 尝试发送消息，并在失败时重试
	for attempt := 0; attempt < maxRetries; attempt++ {
		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: InventoryTopic,
			Value: sarama.ByteEncoder(msgBytes),
		})

		if err == nil {
			// 发送成功，退出循环
			return nil
		}

		// 如果发送失败，等待一段时间后重试
		time.Sleep(time.Duration(attempt+1) * time.Second) // 等待时间随重试次数增加
	}

	// 所有重试都失败了，返回错误
	return fmt.Errorf("failed to send inventory message after %d attempts: %w", maxRetries, err)
}
