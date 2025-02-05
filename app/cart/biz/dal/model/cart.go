// Package model @Author Adrian.Wang 2025/2/4 13:08:00
package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    int32 `gorm:"type:int(11);not null;index:idx_user_id" json:"user_id" binding:"required" label:"用户id"`
	ProductId int32 `gorm:"type:int(11);not null" json:"product_id" binding:"required" label:"商品id"`
	Quantity  int32 `gorm:"type:int(11);not null" json:"quantity" binding:"required" label:"商品数量"`
}

func (bl Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {

	// 1. 查询购物车是否已存在该商品
	var row Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
		First(&row).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 2. 如果存在则更新数量
	if row.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
			UpdateColumn("quantity", gorm.Expr("quantity+?", item.Quantity)).Error
	}

	// 3. 如果不存在则新增
	return db.Create(item).Error
}
