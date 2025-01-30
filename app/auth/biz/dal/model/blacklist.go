// Package model @Author Adrian.Wang 2025/1/27 11:13:00
package model

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type BlackList struct {
	gorm.Model
	UserId int32 `gorm:"index;not null" json:"user_id" binding:"required" label:"用户id"`
	Expire int64 `gorm:"index;not null" json:"expire" binding:"required" label:"黑名单到期时间"`
}

// UserStatus 添加枚举类型
type UserStatus string

const (
	Normal UserStatus = "normal"
	Ban    UserStatus = "ban"
)

func (bl BlackList) TableName() string {
	return "black_list"
}

// GetUserStatusFromDB 从数据库中获取用户状态
// 如果没有找到记录，说明用户状态为正常
// 接收到错误则表示查询数据库失败
func GetUserStatusFromDB(db *gorm.DB, ctx context.Context, userId int32) (status UserStatus, expire int64, err error) {
	var blackList BlackList
	err = db.Where("user_id = ?", userId).First(&blackList).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果没有找到记录，说明用户状态为正常
		return Normal, 0, nil
	}
	if err != nil {
		klog.Error("查询数据库失败:", err)
		return Normal, 0, err
	}

	return Ban, blackList.Expire, nil
}

// AddUserToBlackList 添加用户到黑名单
func AddUserToBlackList(db *gorm.DB, ctx context.Context, userId int32, expire int64) error {
	result := db.Create(&BlackList{
		UserId: userId,
		Expire: expire,
	})
	return result.Error
}

// DeleteFromBlackList 从黑名单中删除用户
func DeleteFromBlackList(db *gorm.DB, ctx context.Context, userId int32) error {
	res := db.Where("user_id = ?", userId).Delete(&BlackList{})
	klog.Debug("DeleteFromBlackList result: ", res.RowsAffected)
	return res.Error
}
