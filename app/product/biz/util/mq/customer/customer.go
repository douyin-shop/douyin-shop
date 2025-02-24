package customer

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/es"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/es/document"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/product/biz/util/mysql"
	"github.com/douyin-shop/douyin-shop/app/product/conf"
	"github.com/go-mysql-org/go-mysql/schema"
	"github.com/kr/pretty"
	"strconv"
)

type MQConsumer struct {
	consumer rocketmq.PushConsumer
}

func NewConsumer() (*MQConsumer, error) {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.GetConf().RocketMQ.NameServer})),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(conf.GetConf().RocketMQ.CustomGroup),
	)
	if err != nil {
		klog.Error("fail to create consumer, err: %v", err)
		return nil, err
	}
	return &MQConsumer{consumer: c}, nil
}

func handleEsMessage(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for _, msg := range msgs {
		var eventMessage mysql.EventMessage
		if err := json.Unmarshal(msg.Body, &eventMessage); err != nil {
			return consumer.ConsumeRetryLater, nil
		}
		// 处理消息
		handleEventMessage(eventMessage)
	}
	return consumer.ConsumeSuccess, nil
}

func (c *MQConsumer) Start(topic string) error {

	err := c.consumer.Subscribe(topic, consumer.MessageSelector{}, handleEsMessage)
	if err != nil {
		return err
	}
	// 启动消费者
	err = c.consumer.Start()

	if err != nil {
		klog.Error("fail to start consumer, err: %v", err)
		return err
	}

	return nil
}

func (c *MQConsumer) Shutdown() {
	c.consumer.Shutdown()
}

func handleEventMessage(msg mysql.EventMessage) {
	klog.Debugf("Consumer 收到消息：%v", msg.Table.Name)
	switch msg.Table.Name {
	case "products":
		handleProductEvent(msg)
	case "categories":
		handleCategoryEvent(msg)
	}
}

func handleProductEvent(msg mysql.EventMessage) {

	// 处理 product 表的事件
	newData, _ := parseRowData(msg.Table, msg.NewData)
	oldData, _ := parseRowData(msg.Table, msg.OldData)

	var p1 model.Product
	var p2 model.Product
	json.Unmarshal(newData, &p1)
	json.Unmarshal(oldData, &p2)
	klog.Debug("=====================================")
	klog.Debug("Consumer 进行产品相关事件的更新操作")
	klog.Debug("操作类型：", msg.EventType)
	klog.Debug("表名：", msg.Table.Name)
	klog.Debug("新数据：", pretty.Sprintf("%# v", msg.NewData))
	klog.Debug("旧数据：", pretty.Sprintf("%# v", msg.OldData))
	klog.Debug("-------------------------------------")
	switch msg.EventType {
	case "insert":
		// 创建产品
		document.CreateProduct(es.Client, conf.GetConf().ElasticSearch.IndexName, p1)
	case "update":
		// 更新产品
		document.UpdateProduct(es.Client, conf.GetConf().ElasticSearch.IndexName, p2, p1)
	case "delete":
		// 删除产品
		document.DeleteProduct(es.Client, conf.GetConf().ElasticSearch.IndexName, p2.ID)
	default:
		// 未知事件类型
		klog.Error("unknown event type")
	}
}

func handleCategoryEvent(msg mysql.EventMessage) {
	// 处理 category 表的事件
	newData, _ := parseRowData(msg.Table, msg.NewData)
	oldData, _ := parseRowData(msg.Table, msg.OldData)
	var c1 model.Category
	var c2 model.Category
	json.Unmarshal(newData, &c1)
	json.Unmarshal(oldData, &c2)
	switch msg.EventType {
	case "update":
		document.UpdateCategory(es.Client, conf.GetConf().ElasticSearch.IndexName, c1, c2)
	case "delete":
		document.DeleteCategory(es.Client, conf.GetConf().ElasticSearch.IndexName, c2)
	}
}

// parseRowData 函数用于将每一行的数据转换为 JSON 格式，并处理二进制数据的 Base64 编码
func parseRowData(table *schema.Table, row []interface{}) (json.RawMessage, error) {
	encodedRow := make(map[string]interface{})
	for i, value := range row {
		column := table.Columns[i]
		switch v := value.(type) {
		case string:
			if column.Type == schema.TYPE_DECIMAL {
				v_str := string(v)
				// 将string转化为浮点数，使用64位精度
				if f, err := strconv.ParseFloat(v_str, 64); err == nil {
					encodedRow[column.Name] = f
				} else {
					// 如果转换失败，保留原始字符串
					encodedRow[column.Name] = float64(0)
				}
			} else {
				encodedRow[column.Name] = string(v)
			}

		default:
			encodedRow[column.Name] = v
		}
	}
	return json.Marshal(encodedRow)
}
