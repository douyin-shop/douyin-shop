package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/douyin-shop/douyin-shop/app/user/biz/dal"
	"github.com/douyin-shop/douyin-shop/app/user/conf"
	"github.com/douyin-shop/douyin-shop/app/user/kitex_gen/user/userservice"
	"github.com/douyin-shop/douyin-shop/common/custom_logger"
	"github.com/douyin-shop/douyin-shop/common/nacos"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// 读取环境变量
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("环境变量文件加载失败", err)
	}

	opts := kitexInit()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {

	dal.Init()

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
	logger.Logger().SetReportCaller(true)
	logger.Logger().SetFormatter(&custom_logger.CustomFormatter{})
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
