package mysql

import (
	"github.com/douyin-shop/douyin-shop/app/cart/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/cart/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&model.Cart{})
	if err != nil {
		panic(err)
	}
}
