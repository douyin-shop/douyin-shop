// @Author Adrian.Wang 2025/2/22 15:39:00
package rocketmq_test

import (
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/rocketmq"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/utils/code"
	"github.com/douyin-shop/douyin-shop/common/custom_logger"
	"github.com/douyin-shop/douyin-shop/common/topic"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"os"
	"testing"
	"time"
)

func init() {
	// 设置当前目录为项目根目录
	err := os.Chdir("../../../")
	if err != nil {
		return
	}

	dal.Init()
}

func TestConsumer(t *testing.T) {
	// klog
	logger := kitexlogrus.NewLogger()
	logger.Logger().SetReportCaller(true)
	logger.Logger().SetFormatter(&custom_logger.CustomFormatter{})
	klog.SetLogger(logger)
	klog.SetLevel(klog.LevelDebug)
	// 订单和交易号存储到json中，然后发送到消息队列
	messageBody := map[string]string{
		"order_id":       "1",
		"transaction_id": "1",
		"time":           time.Now().String(),
	}

	// 转化为json字符串
	messageJson, err := json.Marshal(messageBody)
	if err != nil {
		t.Error(kerrors.NewBizStatusError(code.FailedPayment, err.Error()))
	}

	message, err := rocketmq.SendDelayMessage(topic.GetMsg(topic.Payment), string(messageJson), 3)
	if err != nil {
		t.Error(kerrors.NewBizStatusError(code.FailedPayment, err.Error()))
	}

	t.Logf("message: %v", message)

	// 阻止测试结束
	select {}
}
