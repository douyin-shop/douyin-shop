// Package model @Author Adrian.Wang 2025/2/4 13:08:00
package model

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
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

// AddItem 添加购物车商品
// item 商品信息
func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {

	// 1. 查询购物车是否已存在该商品
	var row Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
		First(&row).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		klog.Error("查询购物车商品失败")
		return err
	}

	// 2. 如果存在则更新数量
	if row.ID > 0 {
		newQuantity := row.Quantity + item.Quantity
		if newQuantity <= 0 {
			return errors.New("更新后的数量不能小于或等于0")
		}
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
			UpdateColumn("quantity", newQuantity).Error
	}

	// 3. 如果不存在则新增
	return db.Create(item).Error
}

// EmptyCart 清空购物车
func EmptyCart(ctx context.Context, db *gorm.DB, userId int32) error {

	if userId == 0 {
		return errors.New("用户id不能为空")
	}

	return db.WithContext(ctx).
		Delete(&Cart{}, "user_id = ?", userId).Error
}

// GetCartByUserId 获取用户购物车
func GetCartByUserId(ctx context.Context, db *gorm.DB, userId int32) ([]*Cart, error) {
	var rows []*Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: userId}).
		Find(&rows).Error
	return rows, err
}
