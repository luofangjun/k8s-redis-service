package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck 心跳检测接口：健康检查
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}