package conf

import (
	"github.com/cloudwego/kitex/pkg/klog"
	common_conf "github.com/douyin-shop/douyin-shop/common/conf"
	"github.com/kitex-contrib/config-nacos/v2/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/kr/pretty"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env string

	Hertz         Hertz         `yaml:"hertz"`
	MySQL         MySQL         `yaml:"mysql"`
	Redis         Redis         `yaml:"redis"`
	Nacos         Nacos         `yaml:"nacos"`
	Jwt           Jwt           `yaml:"jwt"`
	OpenTelemetry OpenTelemetry `yaml:"opentelemetry"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	DB       int    `yaml:"db"`
}

type Hertz struct {
	Service         string `yaml:"service"`
	Address         string `yaml:"address"`
	EnablePprof     bool   `yaml:"enable_pprof"`
	EnableGzip      bool   `yaml:"enable_gzip"`
	EnableAccessLog bool   `yaml:"enable_access_log"`
	LogLevel        string `yaml:"log_level"`
	LogFileName     string `yaml:"log_file_name"`
	LogMaxSize      int    `yaml:"log_max_size"`
	LogMaxBackups   int    `yaml:"log_max_backups"`
	LogMaxAge       int    `yaml:"log_max_age"`
}

type Nacos struct {
	Address             string `yaml:"address"`
	Port                uint64 `yaml:"port"`
	Namespace           string `yaml:"namespace"`
	Group               string `yaml:"group"`
	Username            string `yaml:"username"`
	Password            string `yaml:"password"`
	LogDir              string `yaml:"log_dir"`
	CacheDir            string `yaml:"cache_dir"`
	LogLevel            string `yaml:"log_level"`
	TimeoutMs           uint64 `yaml:"timeout_ms"`
	NotLoadCacheAtStart bool   `yaml:"not_load_cache_at_start"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
}

type OpenTelemetry struct {
	Address string `yaml:"address"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	env := GetEnv()

	prefix := "conf"
	confFileRelPath := filepath.Join(prefix, filepath.Join(env, "conf.yaml"))
	content, err := ioutil.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}

	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}

	err = LoadRemoteConf(env)
	if err != nil {
		klog.Error("load remote config error - %v", err)
		panic(err)
	}

	if err := validator.Validate(conf); err != nil {
		hlog.Error("validate config error - %v", err)
		panic(err)
	}

	conf.Env = GetEnv()

	pretty.Printf("%+v\n", conf)
}

// LoadRemoteConf 从远程加载配置
func LoadRemoteConf(env string) error {

	klog.Info("kitexInit")
	// 从公共配置中加载 Nacos 配置
	nacosConfig := common_conf.GetConf().Nacos

	client, err := nacos.NewClient(nacos.Options{
		Address:     nacosConfig.Address,
		Port:        nacosConfig.Port,
		NamespaceID: nacosConfig.Namespace,
		Group:       nacosConfig.Group,
	})

	if err != nil {
		return err
	}
	client.RegisterConfigCallback(vo.ConfigParam{
		DataId:   "common-config.yaml",
		Group:    env,
		Type:     "yaml",
		OnChange: nil,
	}, func(s string, parser nacos.ConfigParser) {
		klog.Info("远程配置文件变更")

		err = yaml.Unmarshal([]byte(s), conf)
		if err != nil {
			klog.Error("转换配置失败 - %v", err)
		}
		klog.Info("远程配置文件加载成功")
		klog.Debug(pretty.Sprintf("%+v", conf))
	}, 100)

	klog.Info("远程配置文件加载成功")

	return nil
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "test"
	}
	return e
}

func LogLevel() hlog.Level {
	level := GetConf().Hertz.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
