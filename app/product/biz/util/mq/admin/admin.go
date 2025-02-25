// @Author Adrian.Wang 2025/2/25 22:33:00
package admin

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/conf"
)

func CreateTopic(topic string) {
	klog.Debug("创建MQ Topic")
	klog.Debug(conf.GetConf().RocketMQ.NameServer)
	// 管理MQ的Admin对象
	mqAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver([]string{conf.GetConf().RocketMQ.NameServer})))
	if err != nil {
		panic(err)
	}

	err = mqAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(conf.GetConf().RocketMQ.BrokerServer),
	)
	if err != nil {
		panic(err)
	}

}
