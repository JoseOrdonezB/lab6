// @title La Liga Tracker API
// @version 1.0
// @description API para gestionar partidos de fútbol
// @host localhost:8080
// @BasePath /api

package main

import (
	"lab6/database"
	"lab6/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "lab6/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.Connect()

	router := gin.Default()

	router.Use(cors.Default())

	routes.SetupRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
