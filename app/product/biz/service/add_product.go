package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/mysql"
	oss "github.com/douyin-shop/douyin-shop/app/product/biz/util/OSS"
	product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
)

type AddProductService struct {
	ctx context.Context
} // NewAddProductService new AddProductService
func NewAddProductService(ctx context.Context) *AddProductService {
	return &AddProductService{ctx: ctx}
}

// Run create note info
func (s *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	var categories []model.Category
	for _, category := range req.Product.Category {
		category := model.Category{
			Id:       category.Id,
			Name:     category.Name,
			ParentId: category.ParentId,
			Level:    category.Level,
		}
		categories = append(categories, category)
	}
	pro := &model.Product{
		Name:        req.Product.Name,
		Price:       req.Product.Price,
		Description: req.Product.Description,
		ImageName:   req.Product.ImageName,
		Category:    categories,
	}
	url, err := oss.UploadFile(req.Product.Image, int64(len(req.Product.Image)))
	if err != nil {
		klog.Error("upload file error", err)
		return nil, code.GetErr(code.UploadFileError)
	}
	pro.ImageURL = url
	err = model.AddProduct(pro, mysql.Db)
	if err != nil {
		klog.Error("add product error", err)
		return nil, code.GetErr(code.AddProductError)
	}
	return &product.AddProductResp{
		Id: uint32(pro.ID),
	}, nil
}
