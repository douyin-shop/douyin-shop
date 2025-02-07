package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/cart/kitex_gen/cart"
	"github.com/douyin-shop/douyin-shop/app/checkout/biz/utils/code"
	checkout "github.com/douyin-shop/douyin-shop/app/checkout/kitex_gen/checkout"
	"github.com/douyin-shop/douyin-shop/app/checkout/rpc"
	"github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/product"
	cart2 "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/cart"
	"github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
	"github.com/douyin-shop/douyin-shop/app/payment/kitex_gen/payment"
	productService "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
	"strconv"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {

	cartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, code.GetError(code.CartGetFiled)
	}
	if cartResp.Cart == nil || len(cartResp.Cart.Items) == 0 {
		return nil, code.GetError(code.CartEmpty)
	}

	// 总价
	var totalProductPrice float32
	var orderItems []*order.OrderItem
	for _, cartItem := range cartResp.Cart.Items {
		productId := cartItem.ProductId
		quantity := cartItem.Quantity

		// 调用商品微服务获取商品信息
		productResp, err := rpc.ProductClient.GetProduct(s.ctx, &productService.GetProductReq{Id: productId})
		if err != nil {
			return nil, err
		}

		p := productResp.Product

		productDetail := &product.Product{
			Id:          productId,
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price, // TODO 其实这里价格不应该用浮点数，应该用整数，但是这里为了简化，就直接用浮点数了
			Categories:  productResp.Product.Categories,
		}

		// 计算购物车商品总价
		cost := productDetail.Price * float32(quantity)
		totalProductPrice += cost
		orderItems = append(orderItems, &order.OrderItem{
			Item: &cart2.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity,
			},
			Cost: cost,
		})
	}

	// 创建订单
	zipCode, err := strconv.Atoi(req.Address.ZipCode)
	if err != nil {
		return nil, code.GetError(code.ZipCodeError)
	}
	placeOrderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "",
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       int32(zipCode),
		},
		Email:      req.Email,
		OrderItems: orderItems,
	})
	if err != nil {
		return nil, code.GetError(code.PlaceOrderError)
	}

	orderId := placeOrderResp.Order.OrderId

	// 支付操作
	payReq := &payment.ChargeReq{
		Amount: totalProductPrice,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
		OrderId: orderId,
		UserId:  req.UserId,
	}

	// 清空购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		return nil, code.GetError(code.EmptyCartFiled)
	}

	// 调用支付微服务，生成支付信息
	paymentResp, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, code.GetError(code.PayError)
	}

	klog.Debug("支付成功", paymentResp)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResp.TransactionId,
	}
	return
}
