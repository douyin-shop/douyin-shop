package conf

import (
	"github.com/cloudwego/kitex/pkg/klog"
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

	if err := validator.Validate(conf); err != nil {
		klog.Error("validate config error - %v", err)
		panic(err)
	}

	conf.Env = GetEnv()

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

// 获取当前运行的目录相对于根目录的路径
func getRelativePath() string {
	// 获取当前运行的目录
	currentDir, err := os.Getwd()
	if err != nil {
		klog.Error("get current directory error - %v", err)
		panic(err)
	}

	// 获取根目录
	rootDir := "/"

	// 计算相对于根目录的路径
	relativePath, err := filepath.Rel(rootDir, currentDir)
	if err != nil {
		klog.Error("calculate relative path error - %v", err)
		panic(err)
	}

	return relativePath
}
