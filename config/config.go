package config

import (
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// AppConfig 应用配置
type AppConfig struct {
	Count int `mapstructure:"count"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// Config 全局配置
type Config struct {
	App    AppConfig    `mapstructure:"app"`
	Server ServerConfig `mapstructure:"server"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

// GlobalConfig 全局配置变量
var GlobalConfig *Config

// visitAndReplacePlaceholders 遍历配置并替换环境变量占位符
func visitAndReplacePlaceholders() {
	// 遍历所有配置键
	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		// 检查是否为环境变量占位符格式 ${ENV_VAR}
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			// 提取环境变量名称
			envVar := value[2 : len(value)-1]
			// 从环境变量获取实际值
			envValue := os.Getenv(envVar)
			if envValue != "" {
				// 替换占位符为实际值
				viper.Set(key, envValue)
			}
		}
	}
}

// LoadConfig 加载配置
func LoadConfig() {
	// 配置Viper
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 处理环境变量占位符
	visitAndReplacePlaceholders()

	// 解析配置到结构体
	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	// 启用配置文件监听
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件发生变化，重新加载配置...")

		newConfig := &Config{}
		if err := viper.Unmarshal(newConfig); err != nil {
			log.Printf("重新加载配置失败: %v", err)
			return
		}
		GlobalConfig = newConfig
		log.Println("配置重新加载成功")
		log.Printf("应用计数器: %d", GlobalConfig.App.Count)
	})

	log.Println("配置加载成功")
	log.Printf("应用计数器: %d", GlobalConfig.App.Count)
	log.Printf("Redis地址: %s", GlobalConfig.Redis.Address)
}
