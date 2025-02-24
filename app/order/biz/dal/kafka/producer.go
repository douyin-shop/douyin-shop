package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
)

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

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: InventoryTopic,
		Value: sarama.ByteEncoder(msgBytes),
	})
	return err
}
