package routes

import (
	"lab6/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/matches", controllers.GetMatches)
		api.GET("/matches/:id", controllers.GetMatchByID)
		api.POST("/matches", controllers.CreateMatch)
		api.PUT("/matches/:id", controllers.UpdateMatch)
		api.DELETE("/matches/:id", controllers.DeleteMatch)

		api.PATCH("/matches/:id/goals", controllers.UpdateGoals)
		api.PATCH("/matches/:id/yellowcards", controllers.AddYellowCard)
		api.PATCH("/matches/:id/redcards", controllers.AddRedCard)
		api.PATCH("/matches/:id/extratime", controllers.AddExtraTime)
	}
}
