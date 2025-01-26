package mysql

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/douyin-shop/douyin-shop/app/frontend/conf"
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

	// 自动迁移
	err := DB.AutoMigrate(&gormadapter.CasbinRule{})
	if err != nil {
		panic(err)
	}
}
