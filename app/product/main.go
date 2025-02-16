package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/douyin-shop/douyin-shop/common/nacos"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/douyin-shop/douyin-shop/app/product/biz/dal"
	"github.com/douyin-shop/douyin-shop/app/product/biz/util/mq"
	snoyflake "github.com/douyin-shop/douyin-shop/app/product/biz/util/snowflake"
	"github.com/douyin-shop/douyin-shop/app/product/conf"
	"github.com/douyin-shop/douyin-shop/app/product/kitex_gen/product/productcatalogservice"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 读取环境变量
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("环境变量文件加载失败", err)
	}
	opts := kitexInit()

	// init model
	dal.Init()

	// init mq
	mq.InitMq()
	defer mq.ShutdownMq()

	//init snowflake
	snoyflake.Init(conf.GetConf().Snowflake.StartTime,conf.GetConf().Snowflake.MachineId)


	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {

	// OpenTelemetry
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(conf.GetConf().Kitex.Service),
		provider.WithExportEndpoint(conf.GetConf().OpenTelemetry.Address),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	opts = append(opts, server.WithSuite(tracing.NewServerSuite()))

	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// nacos 注册中心
	r := nacos.GetNacosRegistry()
	opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	// 创建一个 MultiWriter，同时写入文件和控制台
	multiWriter := io.MultiWriter(asyncWriter, os.Stdout)
	klog.SetOutput(multiWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
