package controllers

import (
	"net/http"
	"strconv"

	"lab6/database"
	"lab6/models"

	"github.com/gin-gonic/gin"
)

func GetMatches(c *gin.Context) {
	var matches []models.Match
	database.DB.Find(&matches)
	c.JSON(http.StatusOK, matches)
}

func GetMatchByID(c *gin.Context) {
	id := c.Param("id")
	var match models.Match
	if err := database.DB.First(&match, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}
	c.JSON(http.StatusOK, match)
}

func CreateMatch(c *gin.Context) {
	var match models.Match
	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&match)
	c.JSON(http.StatusCreated, match)
}

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
