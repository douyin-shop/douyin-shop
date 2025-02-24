package es

import (
	"fmt"
	"github.com/douyin-shop/douyin-shop/app/product/conf"
	"github.com/olivere/elastic/v7"
)

var Client *elastic.Client

// Init 初始化es
func Init() {
	InitES()
	Client.CreateIndex(conf.GetConf().ElasticSearch.IndexName)
}

func InitES() {
	client, err := elastic.NewClient(
		elastic.SetURL(conf.GetConf().ElasticSearch.Address),
		elastic.SetSniff(false),      //禁止自动检测集群节点
		elastic.SetBasicAuth("", ""), //设置用户名和密码
	)
	if err != nil {
		panic(err)
	}
	Client = client
	fmt.Println("es client init success", client)
}
