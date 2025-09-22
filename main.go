package main

import (
	"gindemo/config"
	"gindemo/middleware"
	"gindemo/router"

	"github.com/gin-gonic/gin"
)

// http://localhost:8080/v1/product/add?name=acc&price=1000
// http://localhost:8080/v2/sn?name=acc&price=1000
// http://localhost:8080/v2/product/add?name=acc&price=1000&sn=720f52ac9335941f38481c5385746298&ts=1758521240
// http://localhost:8080/v1/member/add

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	//engine := gin.Default()  // 开启控制台输出日志
	engine := gin.New() //关闭控制台输出日志
	engine.Use(middleware.LoggerToFile())
	router.InitRouter(engine) // 设置路由
	engine.Run(config.PORT)
}
