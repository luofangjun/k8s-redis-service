package router

import (
	"k8s-redis-service/handler"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 配置路由
func SetupRoutes(r *gin.Engine) {
	// Redis Set接口：存储键值对
	r.POST("/set", handler.SetKey)

	// Redis Get接口：获取键对应的值
	r.GET("/get/:key", handler.GetKey)

	// 心跳检测接口：健康检查
	r.GET("/health", handler.HealthCheck)

	// 获取应用配置接口
	r.GET("/config", handler.GetAppConfig)

	// 获取应用计数器接口
	r.GET("/config/count", handler.GetAppCount)
}
