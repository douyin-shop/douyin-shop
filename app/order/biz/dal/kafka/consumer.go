package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/mysql"
)

// 消费者处理消息的函数
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 消费消息
	for message := range claim.Messages() {
		// 处理消息
		var msg InventoryMessage
		err := json.Unmarshal(message.Value, &msg)
		if err != nil {
			// 消费失败，记录日志
			log.Info("kafka consumer unmarshal message failed, err: %v", err)
			continue
		}
		processMessages(msg)
		// 消费成功，提交消息
		session.MarkMessage(message, "")
	}
	return nil
}

type consumerGroupHandler struct{}

func processMessages(messages InventoryMessage) {
	// 这里实现具体的库存扣减逻辑
	model.UpdateProductStock(mysql.Db, messages.ProductId, messages.Quantity)
}
