# 使用官方Go运行时作为基础镜像
FROM golang:1.23.8-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go mod和sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建优化的应用（禁用符号表、调试信息）
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-w -s' -a -installsuffix cgo -o main .

# 使用更小的基础镜像
FROM alpine:latest

# 安装ca证书（用于HTTPS请求）和时区数据
RUN apk --no-cache add ca-certificates tzdata && \
    # 设置上海时区
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 创建非root用户和用户组
RUN addgroup -g 1000 -S appuser && \
    adduser -u 1000 -S appuser -G appuser

# 设置工作目录
WORKDIR /app

# 从builder阶段复制二进制文件
COPY --from=builder --chown=appuser:appuser /app/main .

# 复制配置文件
COPY --from=builder --chown=appuser:appuser /app/conf ./conf

# 创建日志目录并设置权限
RUN mkdir -p /app/logs && chown -R appuser:appuser /app/logs

# 切换到非root用户
USER appuser

# 暴露端口
EXPOSE 8080

# 运行应用（使用exec格式）
CMD ["./main"]