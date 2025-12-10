package handler

import (
	"k8s-redis-service/database"
	"k8s-redis-service/logger"
	"k8s-redis-service/response"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// SetKey Redis Set接口：存储键值对
func SetKey(c *gin.Context) {
	// 记录请求开始日志
	logger.Info("SetKey接口收到请求")

	// 检查Redis连接是否可用
	if database.Rdb == nil {
		logger.Error("SetKey接口错误: Redis服务不可用")
		response.InternalServerError(c, "Redis服务不可用，请检查服务配置")
		return
	}

	// 获取请求参数
	key := c.Query("key")
	value := c.Query("value")
	logger.Info("SetKey接口参数", "key", key, "value", value)

	// 检查参数是否完整
	if key == "" || value == "" {
		logger.Error("SetKey接口错误: 缺少必要参数key或value")
		response.BadRequest(c, "缺少必要参数key或value")
		return
	}

	// 执行Redis SET命令
	err := database.Rdb.Set(database.Ctx, key, value, 0).Err()
	if err != nil {
		logger.Error("SetKey接口错误: 存储失败", "error", err)
		response.InternalServerError(c, "存储失败: "+err.Error())
		return
	}

	// 记录成功日志
	logger.Info("SetKey接口成功", "key", key, "value", value)

	// 返回成功响应
	data := gin.H{
		"message": "存储成功",
		"key":     key,
		"value":   value,
	}
	response.Success(c, data)
}

// GetKey Redis Get接口：获取键对应的值
func GetKey(c *gin.Context) {
	// 记录请求开始日志
	logger.Info("GetKey接口收到请求")

	// 检查Redis连接是否可用
	if database.Rdb == nil {
		logger.Error("GetKey接口错误: Redis服务不可用")
		response.InternalServerError(c, "Redis服务不可用，请检查服务配置")
		return
	}

	// 获取URL参数
	key := c.Param("key")
	logger.Info("GetKey接口参数", "key", key)

	// 检查参数是否完整
	if key == "" {
		logger.Error("GetKey接口错误: 缺少必要参数key")
		response.BadRequest(c, "缺少必要参数key")
		return
	}

	// 执行Redis GET命令
	value, err := database.Rdb.Get(database.Ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// key不存在的情况
			logger.Warn("GetKey接口警告: key不存在", "key", key)
			response.NotFound(c, "key不存在: "+key)
			return
		}
		// 其他错误
		logger.Error("GetKey接口错误: 查询失败", "error", err)
		response.InternalServerError(c, "查询失败: "+err.Error())
		return
	}

	// 记录成功日志
	logger.Info("GetKey接口成功", "key", key, "value", value)

	// 返回成功响应
	data := gin.H{
		"key":   key,
		"value": value,
	}
	response.Success(c, data)
}
