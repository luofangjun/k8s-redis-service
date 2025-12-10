# Kubernetes部署指南

本指南将指导您如何在本地Kubernetes环境中部署k8s-redis-service应用，支持基于路径的路由功能。

## 目录结构

```
k8s/
├── configmap.yaml              # 应用配置文件（独立）
├── secret.yaml                 # 敏感信息配置（独立）
├── all-in-one.yaml            # 主应用服务配置（Deployment + Service + Ingress）
└── goods-all-in-one.yaml      # 商品服务配置（Deployment + Service）
```

## 部署前准备

1. 确保您的本地已安装并运行Docker Desktop with Kubernetes
2. 确保本地已部署ingress-nginx-controller
3. 确保本地Redis服务正在运行并监听6379端口

## 构建Docker镜像

在部署到Kubernetes之前，需要先构建Docker镜像：

```powershell
# 构建应用镜像
cd d:\python\k8s
docker build -t k8s-redis-service:latest .

# 验证镜像构建成功
docker images | findstr "k8s-redis-service"

# 如果需要推送到Docker Hub或其他镜像仓库（可选）
# docker tag k8s-redis-service:latest yourusername/k8s-redis-service:latest
# docker push yourusername/k8s-redis-service:latest
```

## 部署步骤

### 方法一：使用分离的配置文件部署（推荐用于开发环境）

```powershell
# 应用配置和密钥
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/secret.yaml

# 应用主服务
kubectl apply -f k8s/all-in-one.yaml

# 应用商品服务
kubectl apply -f k8s/goods-all-in-one.yaml
```

## 验证部署

### 1. 检查资源状态

```bash
# 检查Deployment状态
kubectl get deployments

# 检查Pod状态
kubectl get pods

# 检查Service状态
kubectl get services

# 检查Ingress状态
kubectl get ingress

# 检查ConfigMap
kubectl get configmaps

# 检查Secret
kubectl get secrets
```

### 2. 查看详细信息

```bash
# 查看Deployment详细信息
kubectl describe deployment app-deployment

# 查看Pod详细信息
kubectl describe pod <pod-name>

# 查看日志
kubectl logs <pod-name>
```

### 3. 测试应用功能

#### 本地测试（开发环境）
```powershell
# 获取应用配置信息
curl http://localhost:8080/config

# 获取应用计数器
curl http://localhost:8080/config/count

# 健康检查
curl http://localhost:8080/health

# Redis操作
curl -X POST "http://localhost:8080/set?key=testKey&value=testValue"
curl http://localhost:8080/get/testKey
```

#### Kubernetes环境测试（生产环境）
首先，将域名添加到hosts文件中：
```
127.0.0.1 k8s-redis.local
```

然后使用curl命令测试路径路由：

```powershell
# 测试 /user 路径路由到主服务
curl http://k8s-redis.local/user/config
curl http://k8s-redis.local/user/config/count
curl http://k8s-redis.local/user/health

# 测试 /goods 路径路由到商品服务
curl http://k8s-redis.local/goods

# 测试主服务的其他功能
curl -X POST "http://k8s-redis.local/user/set?key=testKey&value=testValue"
curl http://k8s-redis.local/user/get/testKey
```

## 资源清理

如需删除所有资源，请执行：

```powershell
# 删除所有资源（使用分离文件方式部署）
kubectl delete -f k8s/configmap.yaml
kubectl delete -f k8s/secret.yaml
kubectl delete -f k8s/all-in-one.yaml
kubectl delete -f k8s/goods-all-in-one.yaml

```

## 配置说明

### ConfigMap（独立文件）
- 挂载配置文件`conf.yaml`到Pod中
- 配置了Redis地址为`host.docker.internal:6379`
- 方便单独修改应用配置

### Secret（独立文件）
- 安全存储Redis密码
- 使用base64编码存储敏感信息
- 便于权限控制和安全管理

### 主应用服务（all-in-one.yaml）
- **Deployment**：使用k8s-redis-service:latest镜像运行应用，挂载ConfigMap，注入Secret密码，配置资源限制和健康检查
- **Service**：为Deployment提供ClusterIP类型的Service，端口80映射到8080
- **Ingress**：配置基于路径的路由规则

### 商品服务（goods-all-in-one.yaml）
- **Deployment**：使用nginx:alpine镜像运行商品服务，配置资源限制和健康检查
- **Service**：为商品服务提供ClusterIP类型的Service

### Ingress路径路由功能
- 域名：`k8s-redis.local`
- 路径规则：
  - `/user(/|$)(.*)` → 路由到主应用服务（app-service）
  - `/goods(/|$)(.*)` → 路由到商品服务（goods-service）
- 使用`nginx.ingress.kubernetes.io/rewrite-target: /$2`注解实现路径重写
- 支持路径前缀匹配和捕获组重写
