package config

import (
	"github.com/gin-contrib/cors"
)

// GetCORSConfig 获取CORS配置
func GetCORSConfig() cors.Config {
	return cors.Config{
		// 允许的源（域名），这里配置为允许来自localhost:3000的请求
		AllowOrigins: []string{"http://localhost:3000"},
		// 允许的HTTP方法
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 允许的请求头
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Token"},
		// 允许客户端访问的响应头
		ExposeHeaders: []string{"Content-Length"},
		// 是否允许发送凭据（如cookies、认证信息等）
		AllowCredentials: true,
		// 预检请求的缓存时间（秒），这里设置为12小时
		MaxAge: 12 * 60 * 60, // 12 hours
	}
}
