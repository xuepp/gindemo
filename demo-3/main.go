package main

import (
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/xuepp/gindemo/demo-3/docs" // 重要：导入自动生成的 docs 包
)

// @title GoPeck Stress Test API
// @version 1.0
// @description API to simulate CPU and memory load
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// Swagger UI 路径
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CPU 密集型接口
	r.GET("/cpu", cpuHandler)

	// 内存密集型接口
	r.GET("/memory", memoryHandler)

	r.Run(":6081")
}

// CPURequest 请求参数
type CPURequest struct {
	Iterations int `form:"iterations" example:"10000000"` // 可通过 query 参数传入
}

// cpuHandler CPU 压力接口
// @Summary CPU Stress Test
// @Description Perform CPU intensive calculation
// @Tags stress
// @Accept json
// @Produce json
// @Param iterations query int false "Number of iterations" default(10000000)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /cpu [get]
func cpuHandler(c *gin.Context) {
	var req CPURequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Iterations <= 0 {
		req.Iterations = 1e7
	}

	start := time.Now()
	result := 0.0
	for i := 0; i < req.Iterations; i++ {
		result += math.Sqrt(float64(i)) * math.Sin(float64(i))
	}
	duration := time.Since(start)

	c.JSON(http.StatusOK, gin.H{
		"iterations": req.Iterations,
		"result":     result,
		"duration":   duration.String(),
	})
}

// MemRequest 请求参数
type MemRequest struct {
	SizeMB int `form:"sizeMB" example:"100"` // 可通过 query 参数传入
}

// memoryHandler 内存压力接口
// @Summary Memory Stress Test
// @Description Allocate memory to simulate memory usage
// @Tags stress
// @Accept json
// @Produce json
// @Param sizeMB query int false "Memory size in MB" default(100)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /memory [get]
func memoryHandler(c *gin.Context) {
	var req MemRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.SizeMB <= 0 {
		req.SizeMB = 100
	}

	start := time.Now()
	size := req.SizeMB * 1024 * 1024
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}
	duration := time.Since(start)

	c.JSON(http.StatusOK, gin.H{
		"allocated_MB": req.SizeMB,
		"duration":     duration.String(),
	})
}
