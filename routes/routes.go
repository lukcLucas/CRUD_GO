package routes

import (
	"go-mvc-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	bookGroup := r.Group("/books")
	{
		bookGroup.GET("/", controllers.GetBooks)
		bookGroup.GET("/:id", controllers.GetBook)
		bookGroup.POST("/", controllers.CreateBook)
		bookGroup.PUT("/:id", controllers.UpdateBook)
		bookGroup.DELETE("/:id", controllers.DeleteBook)
	}
}
