package middleware

import (
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 检查请求中 JWT 的合法性
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未提供 token"})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "token 无效: " + err.Error()})
			c.Abort()
			return
		}
		// 将用户 ID 存入上下文，后续接口可使用
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
