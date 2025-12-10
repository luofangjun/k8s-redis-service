package database

import (
	"context"
	"fmt"
	"log"

	"k8s-redis-service/config"

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
		log.Fatal("配置未加载，请先调用config.LoadConfig()")
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
		log.Printf("警告: Redis连接失败，请检查Redis服务是否启动且配置正确: %v", err)
		log.Println("服务将继续运行，但Redis相关功能将不可用")
	} else {
		fmt.Println("Redis连接成功")
	}
}
