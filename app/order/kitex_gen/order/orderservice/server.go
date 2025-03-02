// Code generated by Kitex v0.9.1. DO NOT EDIT.
package orderservice

import (
	klog "github.com/cloudwego/kitex/pkg/klog"
	rpcinfo "github.com/cloudwego/kitex/pkg/rpcinfo"
	server "github.com/cloudwego/kitex/server"
	order "github.com/douyin-shop/douyin-shop/app/order/kitex_gen/order"
	registry "github.com/kitex-contrib/registry-nacos/registry"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler order.OrderService, opts ...server.Option) server.Server {
	var options []server.Option
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		klog.Fatal(err)
	}
	options = append(options, server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "order",
	}))

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler order.OrderService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
