package model

import (
	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
	"gorm.io/gorm"
)

type Product struct{
	gorm.Model
	Name string `gorm:"type:varchar(80);not null" json:"product-name" binding:"required,len=9" label:"商品名称"`
	Description string `gorm:"type:text;not null" json:"product-description" binding:"required,len=9" label:"商品描述"`
	Price float64 `gorm:"type:decimal(10,2);not null" json:"product-price" binding:"required,len=9" label:"商品价格"`
	ImageName string  `gorm:"type:varchar(255)" json:"image-name" binding:"omitempty" label:"商品图片名称"`
	ImageURL    string  `gorm:"type:varchar(255)" json:"image-url" binding:"omitempty,url" label:"商品图片URL"`
	Category []Category  `gorm:"many2many:product_categories;"`
}

func CheckProductExist(name string,tx *gorm.DB) int {
	var product *Product
	tx.Where("name = ?",name).First(&product)
	if product.ID != 0 {
	    return code.ProductExist
	}
	return code.ProductNotExist
}

func AddProduct(product *Product,tx *gorm.DB) int{
	c:= CheckProductExist(product.Name,tx)
	if c == code.ProductExist {
	    return c
	}
	err:=tx.Create(product).Error
	if err != nil {
	    return code.Error
	}
	return code.Success
}

func DeleteProduct(id int,tx *gorm.DB) int {
	var product *Product
	tx.Where("id = ?",id).First(&product)
	if product.ID == 0 {
	    return code.ProductNotExist
	}
	tx.Delete(product)
	return code.Success
}

func UpdateProduct(product *Product,tx *gorm.DB) int {
	var p *Product
	tx.Where("id = ?",product.ID).First(&p)
	if p.ID == 0 {
	    return code.ProductNotExist
	}
	tx.Save(product)
	return code.Success
}

func GetProduct(id int,tx *gorm.DB) (*Product,int){
	var product *Product
	err:=tx.Where("id = ?",id).First(&product).Error
	if err != nil {
	    return nil,code.Error
	}
	return product,code.Success
}

func ListProduct(tx *gorm.DB,PageSize, PageNum int) ([]Product,int) {
	var products []Product
	err:=tx.Preload("Category").Offset((PageNum - 1) * PageSize).Limit(PageSize).Find(&products).Error
	if err != nil {
	    return nil,code.Error
	}
	return products,code.Success
}


