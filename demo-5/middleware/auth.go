package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// TokenMiddleware token验证中间件
func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			// 如果Authorization头没有token，尝试从Token头获取
			token = c.GetHeader("Token")
		}

		// 检查token格式（Bearer token 或直接token）
		if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
		}

		// 验证token
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "缺少token",
			})
			c.Abort()
			return
		}

		// 这里可以添加更复杂的token验证逻辑
		// 例如：验证token是否有效、是否过期等
		// 这里简单验证token是否为"demo-token-12345"
		if token != "demo-token-12345" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无效的token",
			})
			c.Abort()
			return
		}

		// token验证通过，继续处理请求
		c.Next()
	}
}
