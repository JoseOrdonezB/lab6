package routes

import (
	"lab6/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas de la API REST
func SetupRoutes(router *gin.Engine) {
	// Agrupamos las rutas bajo el prefijo /api
	api := router.Group("/api")
	{
		// Endpoints CRUD principales para gestionar partidos
		api.GET("/matches", controllers.GetMatches)         // Obtener todos los partidos
		api.GET("/matches/:id", controllers.GetMatchByID)   // Obtener un partido por ID
		api.POST("/matches", controllers.CreateMatch)       // Crear un nuevo partido
		api.PUT("/matches/:id", controllers.UpdateMatch)    // Actualizar un partido existente
		api.DELETE("/matches/:id", controllers.DeleteMatch) // Eliminar un partido

		// Endpoints adicionales para eventos del partido
		api.PATCH("/matches/:id/goals", controllers.UpdateGoals)         // Registrar un gol
		api.PATCH("/matches/:id/yellowcards", controllers.AddYellowCard) // Registrar tarjeta amarilla
		api.PATCH("/matches/:id/redcards", controllers.AddRedCard)       // Registrar tarjeta roja
		api.PATCH("/matches/:id/extratime", controllers.AddExtraTime)    // Registrar tiempo extra
	}
}
