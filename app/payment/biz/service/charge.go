package service

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/rocketmq"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/utils/code"
	payment "github.com/douyin-shop/douyin-shop/app/payment/kitex_gen/payment"
	"github.com/douyin-shop/douyin-shop/common/topic"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
// 这里是真正的支付业务的逻辑，但是我们这里不会真的对接某个支付平台，会用某个库模拟支付的逻辑
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {

	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	err = card.Validate()
	// 允许测试卡通过
	if req.CreditCard.CreditCardNumber != "4242424242424242" && err != nil {
		klog.Error("credit card validate failed", err)
		return nil, kerrors.NewBizStatusError(code.FailedPayment, err.Error())
	}

	// 交易id使用uuid模拟生成一个
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewBizStatusError(code.FailedPayment, err.Error())
	}

	// TODO 虽然我也在想为什么要存储支付日志，应该可以不存储
	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		Amount:        req.Amount,
		TransactionId: transactionId.String(),
		PayTime:       time.Now(),
	})

	if err != nil {
		return nil, kerrors.NewBizStatusError(code.FailedPayment, err.Error())
	}

	// 调用定时消息队列，定时三十分钟，如果没有支付成功，就会取消订单

	// 订单和交易号存储到json中，然后发送到消息队列
	messageBody := map[string]string{
		"order_id":       req.OrderId,
		"transaction_id": transactionId.String(),
	}

	// 转化为json字符串
	messageJson, err := json.Marshal(messageBody)
	if err != nil {
		return nil, kerrors.NewBizStatusError(code.FailedPayment, err.Error())
	}

	message, err := rocketmq.SendDelayMessage(topic.GetMsg(topic.Payment), string(messageJson), 16)
	if err != nil {
		return nil, err
	}

	klog.Debug("Message sent to RocketMQ with transactionId: %s, result: %v", transactionId, message)

	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}
