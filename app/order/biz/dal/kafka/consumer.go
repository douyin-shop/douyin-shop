package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

// 消费者处理消息的函数
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var messages []InventoryMessage
	for message := range claim.Messages() {
		var msg InventoryMessage
		err := json.Unmarshal(message.Value, &msg)
		if err != nil {
			log.Printf("解析消息失败: %v", err)
			continue
		}
		messages = append(messages, msg)
		// 当收到一定数量（这里假设为5条）的消息后进行处理
		if len(messages) >= 5 {
			processMessages(messages)
			// 标记已处理的消息
			for _, m := range messages {
				session.MarkMessage(message, "")
			}
			messages = messages[:0]
		}
	}
	return nil
}

type consumerGroupHandler struct{}

func processMessages(messages []InventoryMessage) {
	// 这里实现具体的库存扣减逻辑
	// 开启事务
}
