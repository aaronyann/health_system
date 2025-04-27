package utils

import "github.com/gin-gonic/gin"

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

func ResponseError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
