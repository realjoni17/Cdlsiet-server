package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/realjoni17/Cdlsiet/database"
	"github.com/realjoni17/Cdlsiet/models"
)

func GetSyllabus(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Syllabus

	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot find Syllabus with id: " + err.Error(),
		})
		return
	}
	c.JSON(200, p)

}
