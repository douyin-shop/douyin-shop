package conf

import (
	"github.com/cloudwego/kitex/pkg/klog"
	common_conf "github.com/douyin-shop/douyin-shop/common/conf"
	"github.com/kitex-contrib/config-nacos/v2/nacos"
	"github.com/kr/pretty"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env      string
	Kitex    Kitex    `yaml:"kitex"`
	MySQL    MySQL    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Registry Registry `yaml:"registry"`
	Nacos    Nacos    `yaml:"nacos"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Kitex struct {
	Service       string `yaml:"service"`
	Address       string `yaml:"address"`
	LogLevel      string `yaml:"log_level"`
	LogFileName   string `yaml:"log_file_name"`
	LogMaxSize    int    `yaml:"log_max_size"`
	LogMaxBackups int    `yaml:"log_max_backups"`
	LogMaxAge     int    `yaml:"log_max_age"`
}

type Registry struct {
	RegistryAddress []string `yaml:"registry_address"`
	Username        string   `yaml:"username"`
	Password        string   `yaml:"password"`
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

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {

	// 设置日志级别,防止日志级别未配置
	klog.SetLevel(klog.LevelDebug)

	// 获取当前环境配置
	env := GetEnv()
	klog.Infof("当前环境: %s", env)

	prefix := "conf"
	confFileRelPath := filepath.Join(prefix, filepath.Join(GetEnv(), "conf.yaml"))
	content, err := ioutil.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}
	conf = new(Config)

	err = yaml.Unmarshal(content, conf)
	if err != nil {
		klog.Error("parse yaml error - %v", err)
		panic(err)
	}

	err = LoadRemoteConf(env)
	if err != nil {
		klog.Error("load remote config error - %v", err)
		panic(err)
	}

	if err := validator.Validate(conf); err != nil {
		klog.Error("validate config error - %v", err)
		panic(err)
	}

	conf.Env = GetEnv()

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

func LogLevel() klog.Level {
	level := GetConf().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
