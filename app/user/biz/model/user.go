package model

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/mysql"
<<<<<<< HEAD
	"github.com/douyin-shop/douyin-shop/app/user/biz/utils/code"
=======
<<<<<<< Updated upstream
	"github.com/douyin-shop/douyin-shop/app/user/code"
=======
	"github.com/douyin-shop/douyin-shop/app/user/biz/utils/code"
>>>>>>> Stashed changes
>>>>>>> ae6c4a5 (测试)
	"github.com/douyin-shop/douyin-shop/app/user/conf"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	PassWord string `gorm:"varchar(20);not null;" json:"password" binding:"required,min=6 max=12" label:"用户密码"`
	Email string `gorm:"varchar(30);not null;" json:"email" binding:"required,email" label:"用户邮箱"`
	Role int `gorm:"tinyint;not null;" json:"role" label:"用户权限"`
}

// 检查用户是否存在
func CheckUserExist(email string) (int,*User) {
	var user User
	mysql.DB.Where("email = ?", email).First(&user)
	if(user.ID!=0){
		return code.UserExist,&user
	}
	return code.UserNotExist,nil
}

//触发器:负责在添加用户信息时,对用户密码进行加密
func (u *User)BeforeSave(tx *gorm.DB) (err error) {
	u.PassWord=ScriptPassWord(u.PassWord)
	return
}

func CreateUser(user *User) (userId uint, err error) {
	err = mysql.DB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func ScriptPassWord(password string)string{
	hashpass,err:=bcrypt.GenerateFromPassword([]byte(password),conf.GetConf().Bcrypt.Cost)
	if err!=nil{
	    klog.Error("password encrypt failed",err)
	}
	return string(hashpass)
}