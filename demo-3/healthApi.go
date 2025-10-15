package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// pingHandler godoc
// @Summary Ping 接口
// @Description 用于测试服务是否可用
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func pingHandler(c *gin.Context) {
	logger.Info("Ping request received",
		zap.String("client_ip", c.ClientIP()),
		zap.String("path", c.Request.URL.Path),
	)
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
