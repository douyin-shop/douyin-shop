// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/tracer/stats"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/dal"
	"github.com/douyin-shop/douyin-shop/app/frontend/infra/rpc"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"io"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/douyin-shop/douyin-shop/app/frontend/biz/router"
	"github.com/douyin-shop/douyin-shop/app/frontend/conf"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"

	"github.com/hertz-contrib/pprof"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// init dal
	dal.Init()
	// 初始化 rpc 客户端
	rpc.InitClient()

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(conf.GetConf().Hertz.Service),
		provider.WithExportEndpoint(conf.GetConf().OpenTelemetry.Address),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	tracer, cfg := hertztracing.NewServerTracer()

	address := conf.GetConf().Hertz.Address
	h := server.New(
		server.WithHostPorts(address),
		tracer,
		server.WithTraceLevel(stats.LevelDetailed), // 开启细粒度的跟踪
	)
	h.Use(hertztracing.ServerMiddleware(cfg))

	registerMiddleware(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	multiWriter := io.MultiWriter(asyncWriter, os.Stdout)
	hlog.SetOutput(multiWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
}
