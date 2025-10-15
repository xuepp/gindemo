package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/xuepp/gindemo/demo-3/docs"
	"github.com/zsais/go-gin-prometheus"
)

// 全局 zap 日志器
var logger *zap.Logger

// @title GoPeck Stress Test API
// @version 1.0
// @description API to simulate CPU and memory load
// @host localhost:6081
// @BasePath /
func main() {
	// 初始化 zap 日志
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	// 初始化 Gin
	r := gin.New()

	// 创建 Prometheus 实例
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r) // 自动注册 /metrics 路由

	r.Use(GinZapLogger(logger)) // ✅ 加入日志中间件
	r.Use(gin.Recovery())

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 接口注册
	r.GET("/ping", pingHandler)
	r.GET("/cpu", cpuHandler)
	r.GET("/memory", memoryHandler)

	logger.Info("Server started on :6081")
	r.Run(":6081")
}

// GinZapLogger 日志中间件
func GinZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		logger.Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.String("client_ip", c.ClientIP()),
			zap.Duration("latency", latency),
		)
	}
}
