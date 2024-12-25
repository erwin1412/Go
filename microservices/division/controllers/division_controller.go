package controllers

import (
	"errors"
	"myapp/config"
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllDivisions(c echo.Context) error {

	var divisions []models.Division

	if err := config.DB.Find(&divisions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	if len(divisions) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "No division found",
		})
	}
	return c.JSON(http.StatusOK, divisions)
}

func GetDivisionByID(c echo.Context) error {
	var division models.Division

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&division).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}
	return c.JSON(http.StatusOK, division)
}

func fetchService(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid reference ID")
	}
	return nil
}

func CreateDivision(c echo.Context) error {
	var input struct {
		CompanyID   uint   `json:"company_id" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
		CreatedBy   uint   `json:"created_by" validate:"required"`
	}

	// created by = null
	input.CreatedBy = 0
	// Bind input dari request
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Validasi CompanyID dengan memanggil Division Service
	companyURL := "http://localhost:8083/api/companies/" + string(rune(input.CompanyID))
	if err := fetchService(companyURL); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid Company ID",
		})
	}

	// Periksa apakah nama divisi sudah ada
	var division models.Division

	if err := config.DB.Where("name = ?", input.Name).First(&division).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// Tangani error selain ErrRecordNotFound
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to query database",
			})
		}
		// Jika ErrRecordNotFound, artinya data tidak ada, lanjutkan proses
	} else {
		// Jika data ditemukan, kembalikan status 409
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "Division already exists",
		})
	}

	// Buat objek divisi
	division = models.Division{
		CompanyID:   input.CompanyID,
		Name:        input.Name,
		Description: input.Description,
		CreatedBy:   input.CreatedBy,
	}

	// Simpan objek divisi ke database
	if err := config.DB.Save(&division).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to save division",
		})
	}

	return c.JSON(http.StatusCreated, division)
}

func UpdateDivision(c echo.Context) error {
	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
		UpdatedBy   uint   `json:"updated_by" validate:"required"`
	}

	// Bind input dari request
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Ambil id dari parameter
	id := c.Param("id")

	// Periksa apakah divisi ada
	var division models.Division
	if err := config.DB.Where("id = ?", id).First(&division).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}

	// Update data divisi
	if err := config.DB.Model(&division).Updates(models.Division{
		Name:        input.Name,
		Description: input.Description,
		UpdatedBy:   input.UpdatedBy,
	}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update division",
		})
	}

	return c.JSON(http.StatusOK, division)
}

func DeleteDivision(c echo.Context) error {
	var division models.Division

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&division).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}

	if err := config.DB.Delete(&division).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete division",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
