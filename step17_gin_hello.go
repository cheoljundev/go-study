package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() // 기본 설정 (로깅+복구 포함)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin Backend!",
		})
	})
	r.Run(":8080")
}
