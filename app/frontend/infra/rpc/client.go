// @Author Adrian.Wang 2025/1/30 11:24:00
package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/douyin-shop/douyin-shop/app/auth/kitex_gen/auth/authservice"
	"github.com/douyin-shop/douyin-shop/app/frontend/conf"
	"github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user/userservice"
	"github.com/douyin-shop/douyin-shop/common/nacos"
	kitextracing "github.com/kitex-contrib/obs-opentelemetry/tracing"
)

var (
	UserClient userservice.Client
	AuthClient authservice.Client
)

func InitClient() {
	// 通过微服务调用user服务
	resolver := nacos.GetNacosResolver()
	klog.Info("初始化rpc client")
	UserClient = userservice.MustNewClient("user",
		client.WithResolver(resolver),
		client.WithSuite(kitextracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Hertz.Service}),
	)
	AuthClient = authservice.MustNewClient("auth",
		client.WithResolver(resolver),
		client.WithSuite(kitextracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Hertz.Service}),
	)
}
