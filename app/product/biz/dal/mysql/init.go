package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/douyin-shop/douyin-shop/app/product/biz/dal/model"
	"github.com/douyin-shop/douyin-shop/app/product/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func Init() {
	conf := conf.GetConf()
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.MySQL.Username,
		conf.MySQL.Password,
		conf.MySQL.Address,
		conf.MySQL.Port,
		conf.MySQL.DbName,
	)

	fmt.Println(dns)

	Db, err = gorm.Open(mysql.Open(dns),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	var sqlDB *sql.DB
	sqlDB, err = Db.DB()

	_ = Db.AutoMigrate(model.Product{}, model.Category{})

	//设置连接池最大连接数量
	sqlDB.SetMaxOpenConns(100)
	//设置连接池最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	//设置连接连接可重用的最大时长
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	return
}
