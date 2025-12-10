package main

import (
	"fmt"

	"k8s-redis-service/config"
	"k8s-redis-service/database"
	"k8s-redis-service/router"

	"github.com/gin-gonic/gin"
)

func main() {
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
	fmt.Printf("服务器启动在 :%d\n", config.GlobalConfig.Server.Port)
	r.Run(fmt.Sprintf(":%d", config.GlobalConfig.Server.Port))
}
