package config

import (
	//"strings"
	"fmt"
	"github.com/fsnotify/fsnotify"
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

var globalConfig *viper.Viper

func init() {
	//fmt.Printf("Loading configuration logics...\n")
	globalConfig = initConfig()
	go dynamicConfig()

}
func initConfig() *viper.Viper {
	globalConfig := viper.New()
	//配置文件的位置
	globalConfig.AddConfigPath("conf")
	globalConfig.SetConfigName("config.toml")
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
	globalConfig.WatchConfig()
	/* fmt.Println("redis port before sleep: ", GlobalConfig.Get("redis.port"))
	time.Sleep(time.Second * 20)
	fmt.Println("redis port after sleep: ", GlobalConfig.Get("redis.port")) */
	globalConfig.OnConfigChange(func(event fsnotify.Event) {
		fmt.Printf("Detect config change: %s \n", event.String())
	})

}

//提供公开接口函数让外界获取配置信息
func GetConfigInfo() Config {
	var c Config
	globalConfig.Unmarshal(&c)
	return c
}
func GetGlobalConfig() *viper.Viper {

	return globalConfig
}
