# Go Redis Web Service

基于Go语言、Gin框架和Redis数据库实现的Web服务，提供键值对存储和查询功能。

## 功能特性

1. Redis Set接口：向Redis存储键值对
2. Redis Get接口：从Redis获取指定键的值
3. 心跳检测接口：服务健康检查

## 接口说明

### 1. Redis Set接口
- **URL**: `/set`
- **方法**: POST
- **参数**: 
  - key: 键
  - value: 值
- **示例**: `curl -X POST "http://localhost:8080/set?key=username&value=john"`

### 2. Redis Get接口
- **URL**: `/get/{key}`
- **方法**: GET
- **参数**: 
  - key: 要查询的键（路径参数）
- **示例**: `curl http://localhost:8080/get/username`

### 3. 心跳检测接口
- **URL**: `/health`
- **方法**: GET
- **示例**: `curl http://localhost:8080/health`

## 安装与运行

1. 确保已安装Go环境（1.16+）
2. 确保Redis服务正在运行，地址：127.0.0.1:6379，密码：123456
3. 克隆项目代码
4. 在项目根目录执行：
   ```bash
   go mod tidy
   go run main.go
   ```

## 依赖

- Gin框架: github.com/gin-gonic/gin
- Redis客户端: github.com/go-redis/redis/v8