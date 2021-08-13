package main

import (
	"belajar-be-dasar-gin/controller"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)

	r.POST("/ping", func(c *gin.Context) {
		name := c.Query("name")

		c.JSON(200, gin.H{
			"name":    name,
			"message": "post pong",
		})

	})

	r.GET("/user", controller.GetUser)
	r.POST("/user", controller.PostUser)

	port := ":9000"
	r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
