package handler

import (
	"log"
	"net/http"

	"k8s-redis-service/database"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// SetKey Redis Set接口：存储键值对
func SetKey(c *gin.Context) {
	// 记录请求开始日志
	log.Println("SetKey接口收到请求")

	// 检查Redis连接是否可用
	if database.Rdb == nil {
		log.Println("SetKey接口错误: Redis服务不可用")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Redis服务不可用，请检查服务配置",
		})
		return
	}

	// 获取请求参数
	key := c.Query("key")
	value := c.Query("value")
	log.Printf("SetKey接口参数: key=%s, value=%s", key, value)

	// 检查参数是否完整
	if key == "" || value == "" {
		log.Println("SetKey接口错误: 缺少必要参数key或value")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "缺少必要参数key或value",
		})
		return
	}

	// 执行Redis SET命令
	err := database.Rdb.Set(database.Ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("SetKey接口错误: 存储失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "存储失败: " + err.Error(),
		})
		return
	}

	// 记录成功日志
	log.Printf("SetKey接口成功: key=%s, value=%s", key, value)

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "存储成功",
		"key":     key,
		"value":   value,
	})
}

// GetKey Redis Get接口：获取键对应的值
func GetKey(c *gin.Context) {
	// 记录请求开始日志
	log.Println("GetKey接口收到请求")

	// 检查Redis连接是否可用
	if database.Rdb == nil {
		log.Println("GetKey接口错误: Redis服务不可用")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Redis服务不可用，请检查服务配置",
		})
		return
	}

	// 获取请求参数
	key := c.Param("key")
	log.Printf("GetKey接口参数: key=%s", key)

	// 检查参数是否完整
	if key == "" {
		log.Println("GetKey接口错误: 缺少必要参数key")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "缺少必要参数key",
		})
		return
	}

	// 执行Redis GET命令
	val, err := database.Rdb.Get(database.Ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// 键不存在
			log.Printf("GetKey接口警告: 键不存在 key=%s", key)
			c.JSON(http.StatusNotFound, gin.H{
				"error": "键不存在",
			})
		} else {
			// 其他错误
			log.Printf("GetKey接口错误: 查询失败: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "查询失败: " + err.Error(),
			})
		}
		return
	}

	// 记录成功日志
	log.Printf("GetKey接口成功: key=%s, value=%s", key, val)

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"key":   key,
		"value": val,
	})
}
