package handler

import (
	"k8s-redis-service/logger"
	"k8s-redis-service/response"

	"github.com/gin-gonic/gin"
)

// HealthCheck 心跳检测接口：健康检查
func HealthCheck(c *gin.Context) {
	logger.Info("HealthCheck接口收到请求")
	data := gin.H{"status": "ok"}
	logger.Info("HealthCheck接口成功响应")
	response.Success(c, data)
}
