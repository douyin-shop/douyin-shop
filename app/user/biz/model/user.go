package model

import "gorm.io/gorm"

type User struct{
	gorm.Model
	PassWord string `gorm:"varchar(20);not null;" json:"password" binding:"required,min=6 max=12" label:"用户密码"`
	Email string `gorm:"varchar(30);not null;" json:"email" binding:"required,email" label:"用户邮箱"`
	Role int `gorm:"tinyint;not null;" json:"role" label:"用户权限"`
}