package model

import (
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(80);not null" json:"name" binding:"required,len=9" label:"商品名称"`
	Description string     `gorm:"type:text;not null" json:"description" binding:"required,len=9" label:"商品描述"`
	Price       float64    `gorm:"type:decimal(10,2);not null" json:"price" binding:"required,len=9" label:"商品价格"`
	ImageName   string     `gorm:"type:varchar(255)" json:"image_name" binding:"omitempty" label:"商品图片名称"`
	ImageURL    string     `gorm:"type:varchar(255)" json:"image_url" binding:"omitempty,url" label:"商品图片URL"`
	Stock       int32      `gorm:"type:int;not null" json:"stock" binding:"required" label:"商品库存"`
	Category    []Category `gorm:"many2many:product_categories;"`
}

func CheckProductExist(name string, tx *gorm.DB) int {
	var product *Product
	tx.Where("name = ?", name).First(&product)
	if product.ID != 0 {
		return code.ProductExist
	}
	return code.ProductNotExist
}

func CheckImage(id uint, imageName string, tx *gorm.DB) bool { //判断在更新商品时要不要更新图片url
	var product *Product
	tx.Where("id = ?", id).First(&product)
	if product.ImageName != imageName {
		return true
	}
	return false
}

func AddProduct(product *Product, tx *gorm.DB) error {
	c := CheckProductExist(product.Name, tx)
	if c == code.ProductExist {
		return code.GetErr(code.ProductExist)
	}
	//查看商品分类是否存在
	for _, category := range product.Category {
		c = CheckCategoryExist(category.Name, tx)
		if c == code.CategoryNotExist {
			AddCategory(&category, tx)
		}
	}
	err := tx.Create(product).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int, tx *gorm.DB) int {
	var product *Product
	tx.Where("id = ?", id).First(&product)
	if product.ID == 0 {
		return code.ProductNotExist
	}
	err := tx.Delete(product).Error
	if err != nil {
		return code.Error
	}
	return code.Success
}

func UpdateProduct(product *Product, tx *gorm.DB) int { //考虑到图片的更新，所以需要返回flag来判断是否需要更新图片
	var p *Product
	tx.Where("id = ?", product.ID).First(&p)
	if p.ID == 0 {
		return code.ProductNotExist
	}
	if err := tx.Model(&p).Updates(product).Error; err != nil {
		return code.Error
	}
	return code.Success
}

func GetProduct(id int, tx *gorm.DB) (*Product, int) {
	var product *Product
	err := tx.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, code.Error
	}
	if product.ID == 0 {
		return nil, code.ProductNotExist
	}
	return product, code.Success
}

func ListProduct(tx *gorm.DB, PageNum, PageSize int) ([]Product, int) {
	var products []Product
	err := tx.Debug().Preload("Category").
		Offset((PageNum - 1) * PageSize).
		Limit(PageSize).
		Find(&products).Error
	if err != nil {
		return nil, code.Error
	}

	klog.Debug("products", products)

	return products, code.Success
}

func UpdateProductStock(tx *gorm.DB, id uint32, quantity int32) (bool, error) {
	var product Product
	// 获取商品信息
	if err := tx.First(&product, id).Error; err != nil {
		return false, err
	}

	// 检查库存是否充足
	if product.Stock < quantity {
		return false, kerrors.NewGRPCBizStatusError(code.StockDecreaseError, code.GetMessage(code.StockDecreaseError))
	}

	// 扣减库存
	product.Stock -= quantity
	// 更新商品库存信息
	if err := tx.Save(&product).Error; err != nil {
		return false, err
	}

	return true, nil
}
