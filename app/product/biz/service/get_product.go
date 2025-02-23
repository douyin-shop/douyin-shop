package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
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

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	pro,c:=model.GetProduct(int(req.Id),mysql.Db)
	if c==code.ProductNotExist{
		return nil,kerrors.NewGRPCBizStatusError(int32(c),code.GetMessage(c))
	}
	if c==code.Error{
		return nil,kerrors.NewGRPCBizStatusError(int32(c),code.GetMessage(c))
	}

	var categories []*product.Category

	for _,category:=range pro.Category{
		c:=&product.Category{
		    Name: category.Name,
		}
		categories=append(categories,c)
	}

	p:=&product.Product{
	    Id:uint32(pro.ID),
		Name:pro.Name,
		Price:pro.Price,
		Description: pro.Description,
		ImageName: pro.ImageName,
		ImageUrl: pro.ImageURL,
		Category: categories,
	}
	return &product.GetProductResp{Product:p},nil
}
