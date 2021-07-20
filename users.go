package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Lat       int    `json:"lat"`
	Log       int    `json:"log"`
}

func GetUsers(c *gin.Context) {
	db := GetMySql()
	var users []User

	err := db.Client.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, "Failed to execute request")
	}

	c.JSON(http.StatusOK, users)
}


func CreateUser(c *gin.Context) {
	user := &User{}
	err := c.ShouldBind(user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "Could not understand request")
		return
	}

	db := GetMySql()
	err = db.Client.Create(user).Error
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, "Failed to execute request")
	}
	c.JSON(http.StatusCreated, user)
}


func DeleteUser(c *gin.Context) {
}


func UpdateUser(c *gin.Context) {
}
