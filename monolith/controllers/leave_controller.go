package controllers

import (
	"myapp/config"
	"myapp/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// leave controller

func GetAllLeaves(c echo.Context) error {
	var leaves models.Leave

	if err := config.DB.Find(&leaves).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, leaves)
}

func GetLeaveByID(c echo.Context) error {
	var leave models.Leave

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&leave).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}

	return c.JSON(http.StatusOK, leave)
}

func CreateLeave(c echo.Context) error {
	var input struct {
		UserID    uint   `json:"user_id" validate:"required"`
		StartDate string `json:"start_date" validate:"required"`
		EndDate   string `json:"end_date" validate:"required"`
		Reason    string `json:"reason" validate:"required"`
		Qty       int    `json:"qty" validate:"required"`
		CreatedBy uint   `json:"created_by" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid start date",
		})
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid end date",
		})
	}

	leave := models.Leave{
		UserID:    input.UserID,
		StartDate: startDate,
		EndDate:   endDate,
		Reason:    input.Reason,
		Qty:       input.Qty,
		CreatedBy: input.CreatedBy,
	}

	if err := config.DB.Save(&leave).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to save leave",
		})
	}

	return c.JSON(http.StatusCreated, leave)
}

func UpdateLeave(c echo.Context) error {
	var input struct {
		UserID    uint   `json:"user_id" validate:"required"`
		StartDate string `json:"start_date" validate:"required"`
		EndDate   string `json:"end_date" validate:"required"`
		Reason    string `json:"reason" validate:"required"`
		Qty       int    `json:"qty" validate:"required"`
		UpdatedBy uint   `json:"updated_by" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	id := c.Param("id")

	var leave models.Leave
	if err := config.DB.Where("id = ?", id).First(&leave).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid start date",
		})
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid end date",
		})
	}

	if err := config.DB.Model(&leave).Updates(models.Leave{
		UserID:    input.UserID,
		StartDate: startDate,
		EndDate:   endDate,
		Reason:    input.Reason,
		Qty:       input.Qty,
		UpdatedBy: input.UpdatedBy,
	}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update leave",
		})
	}

	return c.JSON(http.StatusOK, leave)
}

func DeleteLeave(c echo.Context) error {
	var leave models.Leave

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&leave).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}

	if err := config.DB.Delete(&leave).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete leave",
		})
	}

	return c.JSON(http.StatusOK, leave)
}
