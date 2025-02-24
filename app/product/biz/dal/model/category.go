package model

import (
	"github.com/douyin-shop/douyin-shop/app/product/biz/code"
	"gorm.io/gorm"
)

type Category struct {
	Id       uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	ParentId uint64 `gorm:"column:parent_id;not null" json:"parent_id"`
	Level    uint64 `gorm:"column:level;not null" json:"level"`
}

func CheckCategoryExist(name string, tx *gorm.DB) int {
	var category *Category
	tx.Where("name = ?", name).First(&category)
	if category.Id != 0 {
		return code.CategoryNotExist
	}
	return code.CategoryExist
}

func AddCategory(category *Category, tx *gorm.DB) int {
	err := tx.Create(category).Error
	if err != nil {
		return code.Error
	}
	return code.Success
}
