package routes

import (
    "github.com/gin-gonic/gin"
    "lab6/controllers"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.GET("/matches", controllers.GetMatches)
        api.GET("/matches/:id", controllers.GetMatchByID)
        api.POST("/matches", controllers.CreateMatch)
        api.PUT("/matches/:id", controllers.UpdateMatch)
        api.DELETE("/matches/:id", controllers.DeleteMatch)
    }
}