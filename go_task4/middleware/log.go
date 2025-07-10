package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Params
		fmt.Println("LogMiddleware")
		// 打印所有路由参数（可选）
		for _, param := range params {
			fmt.Printf("Param: %s = %s\n", param.Key, param.Value)
		}
		c.Next()
		fmt.Println("LogMiddleware end")
	}
}
