package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
	payment "github.com/douyin-shop/douyin-shop/app/payment/kitex_gen/payment"
	"github.com/douyin-shop/douyin-shop/app/payment/rpc"
)

type PaymentCallbackService struct {
	ctx context.Context
} // NewPaymentCallbackService new PaymentCallbackService
func NewPaymentCallbackService(ctx context.Context) *PaymentCallbackService {
	return &PaymentCallbackService{ctx: ctx}
}

// Run create note info
func (s *PaymentCallbackService) Run(req *payment.PaymentCallbackReq) (resp *payment.PaymentCallbackResp, err error) {

	// 获取签名类型和签名
	signType := req.SignType
	sign := req.Sign

	klog.Debug("签名类型: ", signType, " 签名: ", sign)

	// 验证签名
	//TODO 使用签名验证订单信息
	//crypto.RSAVerify(publicKey, req.Sign)

	klog.Debug("验证签名通过(实际上没有，这只是个测试)")

	// 获取订单号和交易号
	orderId := req.OrderId
	transactionId := req.TransactionId

	klog.Debug("订单号: ", orderId, " 交易号: ", transactionId)

	//  通知订单服务支付成功，修改订单状态，订单微服务如果发现订单已经被取消，那么就不需要修改订单状态，并返回错误
	//  这样第三方就会进行退款操作
	_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{OrderId: orderId})
	if err != nil {
		klog.Error("订单服务调用失败: ", err)
		return nil, err
	}

	resp = &payment.PaymentCallbackResp{
		Success: true,
		Message: "支付成功",
	}

	return
}
