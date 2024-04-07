package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realjoni17/Cdlsiet/database"
	"github.com/realjoni17/Cdlsiet/models"
)

func GetPost(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Post

	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot find Book with id: " + err.Error(),
		})
		return
	}
	c.JSON(200, p)

}

func CreatePost(c *gin.Context) {
	db := database.GetDatabase()

	// Parse JSON body into a Post struct
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert new post into the database
	if err := db.Create(&newPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// Respond with the created post
	c.JSON(http.StatusCreated, newPost)
}
