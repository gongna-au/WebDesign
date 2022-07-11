package config

import (
	//"strings"
	"fmt"
	"github.com/WebDesign/pkg/helpers"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
	//"time"
)

type Config struct {
	AppName  string
	LogLevel string
	MySQL    MySQLConfig
	Redis    RedisConfig
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	IP   string
	Port int
}

var GlobalConfig *viper.Viper

func init() {
	//fmt.Printf("Loading configuration logics...\n")
	GlobalConfig = InitConfig()
	go dynamicConfig()

}

func InitConfig() *viper.Viper {
	globalConfig := viper.New()
	//配置文件的位置
	globalConfig.AddConfigPath("conf")
	globalConfig.SetConfigName("config")
	globalConfig.SetConfigType("toml")
	//设置配置文件和可执行二进制文件在用一个目录
	//设置配置文件的搜索目录
	//viper.AddConfigPath("$HOME/.appname"
	err := globalConfig.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}

		log.Fatal(err) // 读取配置文件失败致命错误

	}
	return globalConfig

}

func dynamicConfig() {
	GlobalConfig.WatchConfig()
	/* fmt.Println("redis port before sleep: ", GlobalConfig.Get("redis.port"))
	time.Sleep(time.Second * 20)
	fmt.Println("redis port after sleep: ", GlobalConfig.Get("redis.port")) */
	GlobalConfig.OnConfigChange(func(event fsnotify.Event) {
		fmt.Printf("Detect config change: %s \n", event.String())
	})

}

//提供公开接口函数让外界获取配置信息
func GetConfigInfo() Config {
	var c Config
	GlobalConfig.Unmarshal(&c)
	return c
}

func GetGlobalConfig() *viper.Viper {
	return GlobalConfig
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取，如：app.name
// 第二个参数允许传参默认值

func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !GlobalConfig.IsSet(path) || helpers.Empty(GlobalConfig.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return GlobalConfig.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return GlobalConfig.GetStringMapString(path)
}
