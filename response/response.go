package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构体
type Response struct {
	Code int         `json:"code"` // 状态码：0表示成功，其他表示失败
	Msg  string      `json:"msg"`  // 消息：失败原因或成功信息
	Data interface{} `json:"data"` // 数据：没数据就是{}
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "成功",
		Data: data,
	})
}

// SuccessWithMsg 带自定义消息的成功响应
func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: struct{}{},
	})
}

// ErrorWithData 带数据的错误响应
func ErrorWithData(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// BadRequest 参数错误响应
func BadRequest(c *gin.Context, msg string) {
	Error(c, 400, msg)
}

// InternalServerError 服务器内部错误响应
func InternalServerError(c *gin.Context, msg string) {
	Error(c, 500, msg)
}

// NotFound 资源不存在响应
func NotFound(c *gin.Context, msg string) {
	Error(c, 404, msg)
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context, msg string) {
	Error(c, 401, msg)
}
