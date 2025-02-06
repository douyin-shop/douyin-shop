package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	caerService "github.com/douyin-shop/douyin-shop/app/cart/kitex_gen/cart"
	"github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/product"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	cart "github.com/douyin-shop/douyin-shop/app/frontend/hertz_gen/frontend/cart"
	productService "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	userId := h.Context.Value("user_id")

	hlog.Info("请求用户id:", userId)

	userIdInt, err := strconv.Atoi(userId.(string))
	if err != nil {
		return nil, err
	}

	// 调用购物车微服务获取购物车信息
	getCartResp, err := rpc.CartClient.GetCart(h.Context, &caerService.GetCartReq{UserId: uint32(userIdInt)})
	if err != nil {
		return nil, err
	}

	var items []*cart.CartItem
	var totalProductPrice float32 = 0
	for _, item := range getCartResp.Cart.Items {

		productId := item.ProductId
		quantity := item.Quantity

		// 调用商品微服务获取商品信息
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &productService.GetProductReq{Id: productId})
		if err != nil {
			return nil, err
		}

		productDetail := &product.Product{
			Id:          productId,
			Name:        productResp.Product.Name,
			Description: productResp.Product.Description,
			Picture:     productResp.Product.Picture,
			Price:       productResp.Product.Price,
			Categories:  productResp.Product.Categories,
		}

		// 计算购物车商品总价
		totalProductPrice += productDetail.Price * float32(quantity)

		items = append(items, &cart.CartItem{
			ProductId: productId,
			Quantity:  quantity,
			Product:   productDetail,
		})
	}

	return &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId:     uint32(userIdInt),
			Items:      items,
			TotalPrice: totalProductPrice,
		},
	}, nil
}
