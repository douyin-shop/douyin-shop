package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/payment/biz/utils/code"
	payment "github.com/douyin-shop/douyin-shop/app/payment/kitex_gen/payment"
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

	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}
