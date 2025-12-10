package handler

import (
	"k8s-redis-service/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAppConfig 获取应用配置接口
func GetAppConfig(c *gin.Context) {
	// 返回应用配置信息
	c.JSON(http.StatusOK, gin.H{
		"app_count":     config.GlobalConfig.App.Count,
		"server_port":   config.GlobalConfig.Server.Port,
		"redis_address": config.GlobalConfig.Redis.Address,
	})
}

// GetAppCount 获取应用计数器接口
func GetAppCount(c *gin.Context) {
	// 返回应用计数器值
	c.JSON(http.StatusOK, gin.H{
		"count": config.GlobalConfig.App.Count,
	})
}
