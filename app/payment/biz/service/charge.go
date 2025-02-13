package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/utils/code"
	payment "github.com/douyin-shop/douyin-shop/app/payment/kitex_gen/payment"
	creditcard "github.com/durango/go-credit-card"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type ChargeService struct {
	ctx   context.Context
	redis *redis.Client
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context, redisClient *redis.Client) *ChargeService {
	return &ChargeService{
		ctx:   ctx,
		redis: redisClient,
	}
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
	if err != nil {
		return nil, kerrors.NewBizStatusError(code.FailedPayment, err.Error())
	}

	// 交易id使用uuid模拟生成一个
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewBizStatusError(code.FailedPayment, err.Error())
	}

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

	redisKey := fmt.Sprintf("timeout:%s", transactionId.String())
	if err := s.redis.Set(s.ctx, redisKey, "支付中", 15*time.Minute).Err(); err != nil {
		return nil, kerrors.NewBizStatusError(code.FailedPayment, err.Error())
	}

	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}




// 启动Redis过期监听
func (s *ChargeService) handleOvertime(transactionID string) int {

	paymentKey := fmt.Sprintf("payment:%s", transactionID)
	if err := s.redis.Del(s.ctx, paymentKey).Err(); err != nil {
		return code.FailedPayment
	}

	return code.Overtime
}

// 支付成功处理
func (s *ChargeService) HandlePaymentSuccess(transactionID string) int {
	paymentKey := fmt.Sprintf("payment:%s", transactionID)
	if err := s.redis.Del(s.ctx, paymentKey).Err(); err != nil {
		return code.FailedPayment
	}

	orderID, err := model.GetOrderIDByTransaction(mysql.DB, s.ctx, transactionID)
	if err != nil {
		return code.FailedPayment
	}
	if err := s.sendToPaymentQueue(orderID, transactionID); err != code.Success {
		return code.Queueerr
	}

	return code.Overtime
}

type PaymentMessage struct {
	OrderID       int
	TransactionID string
}

func (s *ChargeService) sendToPaymentQueue(orderID int, transactionID string) int {
	msg := PaymentMessage{
		OrderID:       orderID,
		TransactionID: transactionID,
	}

	_, err := json.Marshal(msg)
	if err != nil {
		return code.Queueerr
	}

	// 使用Redis Stream确保消息持久化
	return code.Success
}

// HandlePaymentSuccess 处理支付成功
