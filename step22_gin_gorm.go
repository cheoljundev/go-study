package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type UserStep22 struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("DB 연결 실패: " + err.Error())
	}

	// 자동 마이그레이션 (테이블 자동 생성)
	db.AutoMigrate(&UserStep22{})

	r := gin.Default()

	//Create
	r.POST("/users", func(c *gin.Context) {
		var user UserStep22
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusCreated, user)
	})

	//Read (all)
	r.GET("/users", func(c *gin.Context) {
		var users []UserStep22
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	//Read (one)
	r.GET("/users/:id", func(c *gin.Context) {
		var user UserStep22
		if err := db.First(&user, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	//Update
	r.PUT("/users/:id", func(c *gin.Context) {
		var user UserStep22
		id := c.Param("id")

		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		var updatedUser UserStep22
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.Name = updatedUser.Name
		user.Age = updatedUser.Age
		db.Save(&user)
		c.JSON(http.StatusOK, user)
	})

	//Delete
	r.DELETE("/users/:id", func(c *gin.Context) {

		if err := db.Delete(UserStep22{}, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

	})

	r.Run(":8080")

}
