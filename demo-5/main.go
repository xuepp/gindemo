package main

import (
	"log"

	"demo-5/config"
	"demo-5/handlers"
	"demo-5/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化GORM数据库连接
	err := config.InitDB()
	if err != nil {
		log.Printf("数据库初始化失败: %v", err)
		log.Println("将使用内存数据模式")
	} else {
		// 自动迁移数据库表
		err = config.AutoMigrate()
		if err != nil {
			log.Printf("数据库表迁移失败: %v", err)
		} else {
			log.Println("数据库表迁移成功")
		}
	}

	r := gin.Default()

	// 配置CORS中间件
	r.Use(cors.New(config.GetCORSConfig()))

	// 注册路由（应用token中间件）
	r.GET("/dict/getsetting", middleware.TokenMiddleware(), handlers.GetSettings)

	// 启动服务器
	log.Println("服务器启动在 :8081 端口")
	r.Run(":8081")
}
