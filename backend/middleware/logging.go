package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"log"
)

// LoggingMiddleware 记录请求日志
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("%s %s %d %s", c.Request.Method, c.Request.RequestURI, c.Writer.Status(), duration)
	}
}
