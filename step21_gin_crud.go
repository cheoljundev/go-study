package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserStep21 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []UserStep21{}
var nextId = 1

func main() {
	r := gin.Default()

	//Create
	r.POST("/users", func(c *gin.Context) {
		var newUser UserStep21
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser.ID = nextId
		nextId++
		users = append(users, newUser)
		c.JSON(http.StatusCreated, newUser)
	})

	//Read(all)
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	//Read(one)
	r.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, u := range users {
			if u.ID == id {
				c.JSON(http.StatusOK, u)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	//Update
	r.PUT("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) //strconv.Atoi(...): 문자열을 정수(int)로 바꿈
		var updatedUser UserStep21
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, u := range users {
			if u.ID == id {
				updatedUser.ID = id
				users[i] = updatedUser
				c.JSON(http.StatusOK, updatedUser)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	//Delete
	r.DELETE("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	r.Run(":8080")
}
