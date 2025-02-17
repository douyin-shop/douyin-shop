package model

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/user/biz/utils/code"
	"github.com/douyin-shop/douyin-shop/app/user/conf"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PassWord string `gorm:"varchar(20);not null;" json:"password" binding:"required,min=6 max=12" label:"用户密码"`
	Email    string `gorm:"varchar(30);not null;" json:"email" binding:"required,email" label:"用户邮箱"`
	Phone    string `gorm:"varchar(20);not null;" json:"phone" binding:"required" label:"用户电话"`
	Nickname string `gorm:"varchar(20);not null;" json:"nick_name" binding:"required" label:"用户昵称"`
	Address  string `gorm:"varchar(100);not null;" json:"address" binding:"required" label:"用户地址"`
	Role     int    `gorm:"tinyint;not null;" json:"role" label:"用户权限"`
}

// BeforeSave 触发器:负责在添加用户信息时,对用户密码进行加密
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.PassWord = ScriptPassWord(u.PassWord)
	return
}

// CheckUserExist 检查用户是否存在
func CheckUserExist(db *gorm.DB, email string) (int, *User) {
	var user User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return code.UserExist, &user
	}
	return code.UserNotExist, nil
}

func CreateUser(db *gorm.DB, user *User) (userId uint, err error) {
	err = db.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func ScriptPassWord(password string) string {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(password), conf.GetConf().Bcrypt.Cost)
	if err != nil {
		klog.Error("password encrypt failed", err)
	}
	return string(hashpass)
}

// UpdateUserInfo 更新用户信息
func (u *User) UpdateUserInfo(db *gorm.DB, newInfo map[string]interface{}) error {
	// 使用 GORM 的 Updates 方法更新用户信息
	return db.Model(u).Updates(newInfo).Error
}
