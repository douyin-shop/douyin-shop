package kafka

import (
	"github.com/IBM/sarama"
)

var (
	KafkaBrokerList      = []string{"localhost:9092"} // 替换为实际的Kafka broker地址
	InventoryTopic       = "inventory"                // 扣减库存消息主题
	InventoryFailedTopic = "inventory_failed"         // 扣减库存失败消息主题
	producer             sarama.SyncProducer
	err                  error
)

// 库存扣减消息结构体
type InventoryMessage struct {
	ProductId uint32 `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

// 初始化Kafka配置
func InitKafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true                                          // 生产者需要返回成功信息以便确认消息发送成功
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin() // 消费者组再平衡策略
	return config
}

// 初始化Kafka生产者
func InitProducer() (sarama.SyncProducer, error) {
	config := InitKafkaConfig()
	producer, err = sarama.NewSyncProducer(KafkaBrokerList, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// 获取Kafka生产者
func GetProducer() sarama.SyncProducer {
	if producer == nil {
		producer, err := InitProducer()
		if err != nil {
			panic(err)
		}
		return producer
	}
	return producer
}

// 初始化Kafka消费者组
func InitConsumerGroup() (sarama.ConsumerGroup, error) {
	config := InitKafkaConfig()
	consumerGroup, err := sarama.NewConsumerGroup(KafkaBrokerList, "inventory_consumer_group", config)
	if err != nil {
		return nil, err
	}
	return consumerGroup, nil
}
