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
	position := models.Position{}
	c.Bind(&position)
	result := config.DB.Create(&position)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
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
