package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUser(c *gin.Context) {
	var user User
	user.Username = "agussaputran"
	user.Password = "agus34"

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"code":    http.StatusOK,
		"message": "Ok",
	})
}

func PostUser(c *gin.Context) {
	var user User
	c.Bind(&user)

	fmt.Println(user.Username, user.Password)
	c.JSON(http.StatusOK, user)
}
