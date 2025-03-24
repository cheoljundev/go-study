package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserStep20 struct {
	Name string `json:"name" binding:"required"`     // 필수
	Age  int    `json:"age" binding:"gte=1,lte=120"` // 1세 이상, 120세 이하
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var user UserStep20

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("%s (%d세)님, 환영합니다!", user.Name, user.Age),
		})
	})

	r.Run(":8080")
}
