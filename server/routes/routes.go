package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/realjoni17/Cdlsiet/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("cdlsiet")
	{
		lectures := main
		notes := main
		books := main
		pyq := main
		syllabus := main
		{
			lectures.GET("/lectures", controllers.GetLectures)
			notes.GET("/notes", controllers.GetNotes)
			books.GET("/books", controllers.GetBooks)
			pyq.GET("/pyq", controllers.GetPyq)
			syllabus.GET("/syllabus", controllers.GetSyllabus)

		}

	}

	return router
}
