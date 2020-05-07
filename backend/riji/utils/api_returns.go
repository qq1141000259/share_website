package utils

import "github.com/gin-gonic/gin"

// API 正常返回
func Response(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "",
		"data": data,
	})
}

// API 错误返回
func ValidateError(c *gin.Context, code int, msg string) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}
