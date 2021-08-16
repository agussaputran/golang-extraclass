package controller

import (
	"movie-review/model"

	"github.com/gin-gonic/gin"
)

func (db *DB) GetAllMovie(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": db.Movies,
	})
}

func (db *DB) PostMoview(c *gin.Context) {
	var movie model.Movie

	c.Bind(&movie)

	c.JSON(200, gin.H{
		"data": movie,
	})

	db.Movies = append(db.Movies, movie)

}
