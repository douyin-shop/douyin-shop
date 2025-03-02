package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/redis"
	"time"

	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/mysql"
	product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {

	klog.Debug("GetProductService Run", req)

	productKey := fmt.Sprintf("product:%d", req.Id)

	// 尝试从缓存获取
	if data, err := redis.RedisClient.Get(s.ctx, productKey).Bytes(); err == nil {
		var p product.Product
		if err := json.Unmarshal(data, &p); err == nil {
			klog.Debug("缓存命中！！！加载缓存中的数据：", productKey)
			return &product.GetProductResp{Product: &p}, nil
		}
	}

	// 缓存未命中，从数据库获取
	pro, err := model.GetProduct(int(req.Id), mysql.Db)
	if err != nil {
		klog.Error("GetProduct failed", err)
		return nil, err
	}

	var categories []*product.Category
	for _, category := range pro.Category {
		c := &product.Category{
			Name: category.Name,
		}
		categories = append(categories, c)
	}

	p := &product.Product{
		Id:          uint32(pro.ID),
		Name:        pro.Name,
		Price:       pro.Price,
		Description: pro.Description,
		ImageName:   pro.ImageName,
		ImageUrl:    pro.ImageURL,
		Category:    categories,
	}

	// 设置缓存，过期时间3小时
	if data, err := json.Marshal(p); err == nil {
		redis.RedisClient.Set(s.ctx, productKey, data, 3*time.Hour)
	}

	return &product.GetProductResp{Product: p}, nil
}
