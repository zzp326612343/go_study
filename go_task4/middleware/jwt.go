// middleware/jwt.go
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zzp326612343/go_study/go_task4/utils"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 header 中获取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Missing Authorization Header"})
			c.Abort()
			return
		}

		// 提取 token（格式为 Bearer token）
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid Authorization Header"})
			c.Abort()
			return
		}

		// 解析 token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid Token"})
			c.Abort()
			return
		}

		// 保存用户信息到上下文
		c.Set("userName", claims.UserName)
		c.Set("userId", claims.ID)
		c.Next()
	}
}
