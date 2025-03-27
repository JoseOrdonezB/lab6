package main

import (
    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"lab6/database"
	"lab6/routes"

)

func main() {
    database.Connect()

    router := gin.Default()

	router.Use(cors.Default())

    routes.SetupRoutes(router)

    router.Run(":8080")
}