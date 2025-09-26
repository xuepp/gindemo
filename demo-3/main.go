package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// swagger 相关
	_ "github.com/xuepp/gindemo/demo-3/docs" // 这一行要导入你生成的 docs 包
)

func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/user", User)
	router.Run(PORT) // 默认监听 0.0.0.0:8080

}

const (
	PORT = ":8081"
)

// ListUsers godoc
// @Summary 用户列表
// @Success 200 {array} User
// @Router /users [get]
func User(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "8888888",
	})
}
