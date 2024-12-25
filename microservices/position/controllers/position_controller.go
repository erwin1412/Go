package controllers

import (
	"myapp/config"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllPositions(c echo.Context) error {
	var positions []models.Position
	result := config.DB.Find(&positions)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, positions)
}

func GetPosition(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var position models.Position
	result := config.DB.First(&position, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, position)
}

func CreatePosition(c echo.Context) error {
	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}
	userID := c.Get("userID")
	position := models.Position{
		Name:        input.Name,
		Description: input.Description,
		CreatedBy:   userID.(uint),
	}

	// Simpan user ke database
	if err := config.DB.Create(&position).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create position",
		})
	}

	return c.JSON(http.StatusOK, position)
}

func UpdatePosition(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	position := models.Position{}

	config.DB.First(&position, id)
	c.Bind(&position)
	result := config.DB.Save(&position)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, position)
}

func DeletePosition(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var position models.Position
	result := config.DB.Delete(&position, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, position)
}
