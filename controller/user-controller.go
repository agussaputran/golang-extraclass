package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

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

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	basePath, _ := os.Getwd()

	fileLoc := filepath.Join(basePath, "files", file.Filename)
	// fmt.Println(filename)
	if err := c.SaveUploadedFile(file, fileLoc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
	})
}
