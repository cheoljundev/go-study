package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/greet", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "익명"
		}

		c.JSON(200, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	r.Run(":8080")
}
