package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type UserStep19 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var user UserStep19

		//JSON 바인딩
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Println("에러 발생:", err)
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		c.JSON(200, gin.H{
			"message": user.Name + " (" + fmt.Sprint(user.Age) + "세)님, 환영합니다!",
		})
	})

	r.Run(":8080")
}
