package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/mysql"
	oss "github.com/douyin-shop/douyin-shop/app/product/biz/util/OSS"
	product "github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product"
	"gorm.io/gorm"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	var categories []model.Category
	for _, category := range req.UpdatedProduct.Category {
		category := model.Category{
			Id:       category.Id,
			Name:     category.Name,
			ParentId: category.ParentId,
			Level:    category.Level,
		}
		categories = append(categories, category)
	}
	pro := &model.Product{
		Model: gorm.Model{
			ID: uint(req.UpdatedProduct.Id),
		},
		Name:        req.UpdatedProduct.Name,
		Price:       req.UpdatedProduct.Price,
		Description: req.UpdatedProduct.Description,
		ImageName:   req.UpdatedProduct.ImageName,
		Category:    categories,
	}
	flag := model.CheckImage(uint(req.UpdatedProduct.Id), req.UpdatedProduct.ImageName, mysql.Db) //判断图片是否需要重新上传

	if flag {
		url, err := oss.UploadFile(req.UpdatedProduct.Image, int64(len(req.UpdatedProduct.Image)))
		if err != nil {
			klog.Error("upload file error", err)
			return nil, code.GetErr(code.UploadFileError)
		} else {
			pro.ImageURL = url
		}
	}
	c := model.UpdateProduct(pro, mysql.Db)
	if c == code.Error {
		return nil, kerrors.NewGRPCBizStatusError(int32(c), code.GetMessage(c))
	}
	return &product.UpdateProductResp{
		Success: true,
	}, nil
}
