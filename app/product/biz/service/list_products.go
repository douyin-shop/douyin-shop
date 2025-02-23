package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/mysql"
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

	klog.Debug(s.ctx, "ListProductsService Run")

	pro, c := model.ListProduct(mysql.Db, int(req.Page), int(req.PageSize))

	if c == code.Error {
		return nil, kerrors.NewGRPCBizStatusError(int32(c), code.GetMessage(c))
	}
	var products []*product.Product
	for _, product := range pro {
		products = append(products, PtoM1(product))
	}

	return &product.ListProductsResp{
		Products: products,
	}, nil
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
