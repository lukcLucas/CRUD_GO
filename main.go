package main

import (
	"go-mvc-api/config"
	"go-mvc-api/models"
	"go-mvc-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Book{})

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
