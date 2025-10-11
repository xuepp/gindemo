package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// swagger 相关
	_ "github.com/xuepp/gindemo/demo-3/docs" // 这一行要导入你生成的 docs 包
)

func cc() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/users", getUser)
	router.Run(PORT) // 默认监听 0.0.0.0:8080

}

const (
	PORT = ":8081"
)

// User 用户信息
type User struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"Tom"`
	Email string `json:"email" example:"tom@example.com"`
}

// @Summary 获取用户列表
// @Description 返回所有用户
// @Tags user
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUser(c *gin.Context) {
	users := []User{
		{ID: 1, Name: "Tom", Email: "tom@163.com"},
		{ID: 2, Name: "Jerry", Email: "jerry@163.com"},
		{ID: 3, Name: "Alice", Email: "Alice@163.com"},
	}
	c.JSON(200, users)
}
