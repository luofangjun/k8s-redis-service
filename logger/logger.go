package logger

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

// Logger 全局日志记录器
var Logger *slog.Logger

// InitLogger 初始化日志记录器
func InitLogger() error {
	// 动态确定日志目录路径
	// 优先使用环境变量指定的目录，否则使用当前工作目录下的logs目录
	logDir := os.Getenv("LOG_DIR")
	if logDir == "" {
		// 获取当前工作目录
		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}
		// 使用当前工作目录下的logs目录
		logDir = filepath.Join(currentDir, "logs")
	}

	// 创建logs目录
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	// 生成按天的日志文件名
	today := time.Now().Format("2006-01-02")
	logFile := filepath.Join(logDir, today+".log")

	// 打开日志文件（追加模式）
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// 创建JSON格式的日志处理器
	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo, // 日志级别：Info
	})

	// 创建日志记录器
	Logger = slog.New(handler)

	// 记录日志系统启动信息
	Logger.Info("日志系统初始化完成", "log_file", logFile, "log_dir", logDir)

	return nil
}

// Info 记录信息级别日志
func Info(msg string, args ...any) {
	if Logger != nil {
		Logger.Info(msg, args...)
	}
}

// Error 记录错误级别日志
func Error(msg string, args ...any) {
	if Logger != nil {
		Logger.Error(msg, args...)
	}
}

// Warn 记录警告级别日志
func Warn(msg string, args ...any) {
	if Logger != nil {
		Logger.Warn(msg, args...)
	}
}

// Debug 记录调试级别日志
func Debug(msg string, args ...any) {
	if Logger != nil {
		Logger.Debug(msg, args...)
	}
}

// WithContext 添加上下文信息的日志记录
func WithContext(ctx context.Context) *slog.Logger {
	if Logger != nil {
		return Logger
	}
	return slog.Default()
}
