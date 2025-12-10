package main

import (
	"fmt"

	"k8s-redis-service/config"
	"k8s-redis-service/database"
	"k8s-redis-service/logger"
	"k8s-redis-service/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志系统
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("初始化日志系统失败: %v\n", err)
		panic("初始化日志系统失败: " + err.Error())
	}
	defer logger.Info("服务器正常关闭")

	// 加载配置
	config.LoadConfig()

	// 初始化Redis
	database.InitRedis()

	// 初始化Gin引擎
	r := gin.Default()

	// 添加 recovery 中间件，用于捕获 panic 并恢复，确保服务稳定性
	r.Use(gin.Recovery())

	// 配置路由
	router.SetupRoutes(r)

	// 启动服务器
	logger.Info("服务器启动", "port", config.GlobalConfig.Server.Port)
	if err := r.Run(":" + fmt.Sprintf("%d", config.GlobalConfig.Server.Port)); err != nil {
		logger.Error("服务器启动失败", "error", err)
	}
}
