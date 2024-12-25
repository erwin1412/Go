package controllers

import (
	"myapp/config"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FindCompaniesWithPagination(c echo.Context) error {
	search := c.QueryParam("search") // Ambil parameter 'search' dari query string

	var companies []models.Company
	var total int64

	query := config.DB.Model(&models.Company{})
	if search != "" {
		// Tambahkan logika pencarian jika parameter 'search' ada
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Hitung total data (termasuk hasil pencarian jika ada)
	if err := query.Count(&total).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Ambil data berdasarkan limit dan offset
	if err := query.Limit(limit).Offset(offset).Find(&companies).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Kirimkan respons dengan data dan metadata total
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  companies,
		"total": total,
	})
}

func FindCompanies(c echo.Context) error {
	var companies []models.Company

	if err := config.DB.Find(&companies).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, companies)
}

func FindOneCompany(c echo.Context) error {
	var company []models.Company

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&company).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}
	return c.JSON(http.StatusOK, company)
}

func CreateCompany(c echo.Context) error {
	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	// Bind input dari request
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Periksa apakah nama perusahaan sudah ada
	var company models.Company
	if err := config.DB.Where("name = ?", input.Name).First(&company).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			// Tangani error selain ErrRecordNotFound
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to query database",
			})
		}
		// Jika ErrRecordNotFound, artinya data tidak ada, lanjutkan proses
	} else {
		// Jika data ditemukan, kembalikan status 409
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "Company already exists",
		})
	}

	userID := c.Get("userID")
	// Buat data baru jika tidak ada conflict
	newCompany := models.Company{
		Name:        input.Name,
		Description: input.Description,
		CreatedBy:   userID.(uint), // Hardcode user ID sementara
	}
	if err := config.DB.Create(&newCompany).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create company",
		})
	}

	// Bentuk respons
	response := models.CompanyResponse{
		ID:          newCompany.ID,
		Name:        newCompany.Name,
		Description: newCompany.Description,
		CreatedAt:   newCompany.CreatedAt,
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdateCompany(c echo.Context) error {
	var company models.Company

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&company).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}

	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	company.Name = input.Name
	company.Description = input.Description

	if err := config.DB.Save(&company).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update company",
		})
	}

	return c.JSON(http.StatusOK, company)
}

func DeleteCompany(c echo.Context) error {
	var company models.Company

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First((&company)).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}

	if err := config.DB.Delete(&company).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete company",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Company deleted"})
}
