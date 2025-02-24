package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/biz/util/mq/producer"
	snoyflake "github.com/douyin-shop/douyin-shop/app/product/biz/util/snowflake"
	"github.com/douyin-shop/douyin-shop/app/product/conf"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/schema"
)

// EventMessage 结构体定义
type EventMessage struct {
	Table     *schema.Table
	EventType string        `json:"event_type"`
	ID        uint64        `json:"id"`
	OldData   []interface{} `json:"old_data"`
	NewData   []interface{} `json:"new_data"`
}

// MyEventHandler 结构体定义
type MyEventHandler struct {
	canal.DummyEventHandler
	mqProducer *producer.MQProducer
}

func GetId() uint64 {
	id, err := snoyflake.GetID()
	if err != nil {
		klog.Error("get id error,err:", err)
		return 0
	}
	return id
}

func serializeEventMessage(msg EventMessage) []byte {
	data, _ := json.Marshal(msg)
	return data
}

// OnRow 方法处理不同类型的事件
// 定义MyEventHandler结构体的OnRow方法，用于处理canal.RowsEvent事件
func (m *MyEventHandler) OnRow(event *canal.RowsEvent) error {
	// 根据事件类型进行不同的处理
	switch event.Action {
	case canal.InsertAction:
		// 处理插入事件
		m.handleInsert(event)
	case canal.UpdateAction:
		// 处理更新事件
		m.handleUpdate(event)
	case canal.DeleteAction:
		// 处理删除事件
		m.handleDelete(event)
	}
	// 返回nil表示处理成功
	return nil
}

// handleInsert 处理插入事件
func (m *MyEventHandler) handleInsert(e *canal.RowsEvent) {
	for _, row := range e.Rows {
		msg := EventMessage{
			Table:     e.Table,
			ID:        GetId(), //生成唯一ID
			EventType: "insert",
			NewData:   row,
		}
		klog.Debug("监听到插入事件", msg)
		err := m.mqProducer.PutMessage(conf.GetConf().RocketMQ.Topic, serializeEventMessage(msg))
		if err != nil {
			klog.Error("put message error,err:", err)
		}
	}
}

// handleUpdate 处理更新事件
func (m *MyEventHandler) handleUpdate(e *canal.RowsEvent) {
	if len(e.Rows) < 2 {
		return
	}

	for i := 0; i < len(e.Rows)/2; i++ {
		oldRow := e.Rows[i*2]
		newRow := e.Rows[i*2+1]
		msg := EventMessage{
			Table:     e.Table,
			ID:        GetId(), // 生成唯一ID
			EventType: "update",
			OldData:   oldRow,
			NewData:   newRow,
		}
		klog.Debug("监听到更新事件", msg)
		err := m.mqProducer.PutMessage(conf.GetConf().RocketMQ.Topic, serializeEventMessage(msg))
		if err != nil {
			klog.Error("put message error,err:", err)
		}
	}
}

// handleDelete 处理删除事件
func (m *MyEventHandler) handleDelete(e *canal.RowsEvent) {
	for _, row := range e.Rows {
		msg := EventMessage{
			Table:     e.Table,
			ID:        GetId(),
			EventType: "delete",
			OldData:   row,
		}
		klog.Debug("监听到删除事件", msg)
		err := m.mqProducer.PutMessage(conf.GetConf().RocketMQ.Topic, serializeEventMessage(msg))
		if err != nil {
			klog.Error("put message error,err:", err)
		}
	}
}

// RunListen 启动监听
func RunListen(producer *producer.MQProducer) {
	// 初始化配置
	cfg := initCfg()
	// 创建一个canal实例
	c, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}

	// 注册回调函数
	c.SetEventHandler(&MyEventHandler{mqProducer: producer})

	if err = c.Run(); err != nil {
		klog.Error("Canal 运行失败", err)
	}

}

// 初始化配置函数
func initCfg() *canal.Config {
	// 创建一个默认的canal配置
	cfg := canal.NewDefaultConfig()
	// 设置canal的地址
	cfg.Addr = fmt.Sprintf("%s:%d", conf.GetConf().MySQL.Address, conf.GetConf().MySQL.Port)
	// 设置canal的用户名
	cfg.User = conf.GetConf().MySQL.Username
	// 设置canal的密码
	cfg.Password = conf.GetConf().MySQL.Password
	// 设置canal的执行路径
	cfg.Dump.ExecutionPath = ""
	// 设置canal的服务器ID
	cfg.ServerID = conf.GetConf().MySQL.ServerId

	// 设置监听的表
	tables := []string{
		fmt.Sprintf("%s\\.product", conf.GetConf().MySQL.DbName),
		fmt.Sprintf("%s\\.category", conf.GetConf().MySQL.DbName),
	}

	cfg.IncludeTableRegex = tables
	return cfg
}
