package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/mysql"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/redis"
	product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
	redis_core "github.com/redis/go-redis/v9"
	"math/rand"
	"time"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run 获取产品列表，带分页缓存
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	klog.Debug(s.ctx, "ListProductsService Run")

	// 生成唯一缓存键
	cacheKey := fmt.Sprintf("products:list:%d:%d", req.Page, req.PageSize)
	const cacheDuration = 5 * time.Minute

	// 尝试从缓存获取
	if cacheData, err := redis.RedisClient.Get(s.ctx, cacheKey).Bytes(); err == nil {
		var cachedProducts []model.Product
		if err := json.Unmarshal(cacheData, &cachedProducts); err == nil {
			klog.Debug("缓存命中！！！加载缓存中的数据：", cacheKey)
			return convertProducts(cachedProducts), nil
		}
		klog.Errorf("Cache unmarshal error: %v", err)
	} else if !errors.Is(err, redis_core.Nil) {
		klog.Errorf("Cache access error: %v", err)
	}

	// 缓存未命中，查询数据库
	pro, c := model.ListProduct(mysql.Db, int(req.Page), int(req.PageSize))
	if c == code.Error {
		return nil, kerrors.NewGRPCBizStatusError(int32(c), code.GetMessage(c))
	}

	// 异步更新缓存（带随机过期时间防止雪崩）
	go func(data []model.Product) {
		if cacheData, err := json.Marshal(data); err == nil {
			expire := cacheDuration + time.Second*time.Duration(rand.Intn(60))
			if err := redis.RedisClient.Set(s.ctx, cacheKey, cacheData, expire).Err(); err != nil {
				klog.Errorf("Cache set failed: %v", err)
			}
		}
	}(pro)

	return convertProducts(pro), nil
}

// convertProducts 转换模型到响应格式
func convertProducts(pro []model.Product) *product.ListProductsResp {
	products := make([]*product.Product, 0, len(pro))
	for _, p := range pro {
		products = append(products, PtoM1(p))
	}
	return &product.ListProductsResp{
		Products: products,
	}
}
func PtoM(c model.Category) *product.Category {
	return &product.Category{
		Id:       uint64(c.Id),
		Level:    uint64(c.Level),
		Name:     c.Name,
		ParentId: uint64(c.ParentId),
	}
}

func PtoM1(p model.Product) *product.Product {
	var c []*product.Category
	for _, category := range p.Category {
		c = append(c, PtoM(category))
	}
	return &product.Product{
		Id:          uint32(p.ID),
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
		ImageName:   p.ImageName,
		ImageUrl:    p.ImageURL,
		Category:    c,
	}
}
