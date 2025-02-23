package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/es"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/es/document"
	"github.com/douyin-shop/douyin-shop/app/product/conf"
	product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info

/**
这里搜索商品基于ES的定义,一共有两种搜索方式,模糊搜索和精确搜索,涉及参数为：
Keyword     string 
PageNum     int     
PageSize    int    
CategoryName string 
MinPrice    float64 
MaxPrice    float64
这里我们通过KeyWord和CategoryName来确定搜索方式,如果KeyWord不为空,
则进行模糊搜索,如果CategoryName不为空,则进行精确搜索,如果两个都不为空,
则进行混合搜索,混合搜索的规则为:先进行模糊搜索,再对模糊搜索的结果进行过滤,
**/

func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	pros,c:=document.CombinedSearchProduct(es.Client, conf.GetConf().ElasticSearch.IndexName,
	req.SearchQuery.KeyWord, req.SearchQuery.CategoryName, 
	req.SearchQuery.MinPrice, req.SearchQuery.MaxPrice, 
	int(req.SearchQuery.PageNum),int(req.SearchQuery.PageSize))
	if c==code.Error{
		return nil,kerrors.NewGRPCBizStatusError(int32(c),code.GetMessage(c))
	}
	var products []*product.Product
	for _,product:=range pros{
		products=append(products,PtoM1(product))
	}
	return &product.SearchProductsResp{
		Results:products,
	}, nil
}