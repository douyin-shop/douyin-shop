package mysql

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/douyin-shop/douyin-shop/app/auth/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/auth/conf"

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

	err := DB.AutoMigrate(&model.BlackList{})
	if err != nil {
		klog.Fatal("AutoMigrate BlackList error: ", err)
		panic(err)
	}
}
