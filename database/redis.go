package database

import (
	"context"

	"k8s-redis-service/config"
	"k8s-redis-service/logger"

	"github.com/go-redis/redis/v8"
)

// Ctx 全局上下文
var Ctx = context.Background()

// Rdb Redis客户端实例
var Rdb *redis.Client

// InitRedis 初始化Redis客户端
func InitRedis() {
	// 确保配置已加载
	if config.GlobalConfig == nil {
		logger.Error("配置未加载，请先调用config.LoadConfig()")
		return
	}

	// 初始化Redis客户端
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.GlobalConfig.Redis.Address,  // Redis服务器地址和端口
		Password: config.GlobalConfig.Redis.Password, // Redis密码
		DB:       config.GlobalConfig.Redis.DB,       // 使用默认DB
	})

	// 测试Redis连接
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		logger.Warn("Redis连接失败，请检查Redis服务是否启动且配置正确", "error", err)
		logger.Info("服务将继续运行，但Redis相关功能将不可用")
	} else {
		logger.Info("Redis连接成功")
	}
}
