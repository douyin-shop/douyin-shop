package service

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/bxcodec/faker/v3"
	product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.

	products := make([]*product.Product, 10)

	// 生成10个假数据
	for i := 0; i < 10; i++ {

		randomPrice := gofakeit.Price(100, 10000)
		products[i] = &product.Product{
			Id:          uint32(i + 1),
			Name:        faker.Name(),
			Description: faker.Sentence(),
			Price:       float32(randomPrice),                 // 1-1000之间的随机价格
			Categories:  []string{faker.Word(), faker.Word()}, // 随机分类
		}
	}

	resp = &product.ListProductsResp{Products: products}

	return
}
