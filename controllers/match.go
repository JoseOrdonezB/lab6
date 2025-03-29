package controllers

import (
	"net/http"
	"strconv"

	"lab6/database"
	"lab6/models"

	"github.com/gin-gonic/gin"
)

// GetMatches obtiene todos los partidos registrados
// @Summary Obtener todos los partidos
// @Description Devuelve la lista completa de partidos
// @Tags Matches
// @Produce json
// @Success 200 {array} models.Match
// @Router /matches [get]
func GetMatches(c *gin.Context) {
	var matches []models.Match
	database.DB.Find(&matches)
	c.JSON(http.StatusOK, matches)
}

// GetMatchByID obtiene un partido por ID
// @Summary Obtener partido por ID
// @Description Devuelve los datos de un partido específico
// @Tags Matches
// @Produce json
// @Param id path int true "ID del partido"
// @Success 200 {object} models.Match
// @Router /matches/{id} [get]
func GetMatchByID(c *gin.Context) {
	id := c.Param("id")
	var match models.Match
	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}
	c.JSON(http.StatusOK, match)
}

// CreateMatch crea un nuevo partido
// @Summary Crear partido
// @Description Crea un nuevo partido con la información proporcionada
// @Tags Matches
// @Accept json
// @Produce json
// @Param match body models.Match true "Datos del partido"
// @Success 201 {object} models.Match
// @Router /matches [post]
func CreateMatch(c *gin.Context) {
	var match models.Match
	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&match)
	c.JSON(http.StatusCreated, match)
}

// UpdateMatch actualiza un partido existente
// @Summary Actualizar partido
// @Description Actualiza los datos de un partido existente
// @Tags Matches
// @Accept json
// @Produce json
// @Param id path int true "ID del partido"
// @Param match body models.Match true "Nuevos datos"
// @Success 200 {object} models.Match
// @Router /matches/{id} [put]
func UpdateMatch(c *gin.Context) {
	id := c.Param("id")
	var match models.Match
	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	var input models.Match
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match.HomeTeam = input.HomeTeam
	match.AwayTeam = input.AwayTeam
	match.MatchDate = input.MatchDate

	database.DB.Save(&match)
	c.JSON(http.StatusOK, match)
}

// DeleteMatch elimina un partido
// @Summary Eliminar partido
// @Description Elimina un partido por su ID
// @Tags Matches
// @Param id path int true "ID del partido"
// @Success 204
// @Router /matches/{id} [delete]
func DeleteMatch(c *gin.Context) {
	id := c.Param("id")
	idNum, _ := strconv.Atoi(id)
	result := database.DB.Delete(&models.Match{}, idNum)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateGoals registra un gol en el partido
// @Summary Registrar gol
// @Description Aumenta en 1 el contador de goles totales
// @Tags Matches
// @Param id path int true "ID del partido"
// @Success 200 {object} models.Match
// @Router /matches/{id}/goals [patch]
func UpdateGoals(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	match.TotalGoals++

	if err := database.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el partido"})
		return
	}

	c.JSON(http.StatusOK, match)
}

// AddYellowCard registra una tarjeta amarilla
// @Summary Registrar tarjeta amarilla
// @Description Aumenta en 1 el número de tarjetas amarillas del partido
// @Tags Matches
// @Param id path int true "ID del partido"
// @Success 200 {object} models.Match
// @Router /matches/{id}/yellowcards [patch]
func AddYellowCard(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	match.YellowCards++

	if err := database.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el partido"})
		return
	}

	c.JSON(http.StatusOK, match)
}

// AddRedCard registra una tarjeta roja
// @Summary Registrar tarjeta roja
// @Description Aumenta en 1 el número de tarjetas rojas del partido
// @Tags Matches
// @Param id path int true "ID del partido"
// @Success 200 {object} models.Match
// @Router /matches/{id}/redcards [patch]
func AddRedCard(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	match.RedCards++

	if err := database.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el partido"})
		return
	}

	c.JSON(http.StatusOK, match)
}

// AddExtraTime registra tiempo extra en el partido
// @Summary Registrar tiempo extra
// @Description Marca que el partido tuvo tiempo extra
// @Tags Matches
// @Param id path int true "ID del partido"
// @Success 200 {object} models.Match
// @Router /matches/{id}/extratime [patch]
func AddExtraTime(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	match.ExtraTime = true

	if err := database.DB.Save(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el partido"})
		return
	}

	c.JSON(http.StatusOK, match)
}
