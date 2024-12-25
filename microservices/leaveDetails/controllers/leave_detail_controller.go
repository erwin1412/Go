package controllers

import (
	"myapp/config"
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllLeaveDetails(c echo.Context) error {
	var leaveDetails models.LeaveDetail

	if err := config.DB.Find(&leaveDetails).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, leaveDetails)
}

func GetLeaveDetailByID(c echo.Context) error {
	var leaveDetail models.LeaveDetail

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&leaveDetail).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}

	return c.JSON(http.StatusOK, leaveDetail)
}

//	type LeaveDetail struct {
//		ID         uint           `gorm:"primaryKey"`
//		LeaveID    uint           `gorm:"not null"`                    // Foreign key ke Leave
//		Leave      Leave          `gorm:"foreignKey:LeaveID"`          // Relasi ke Leave
//		ApprovedBy uint           `gorm:"default:null"`                // User ID yang menyetujui
//		RejectedBy uint           `gorm:"default:null"`                // User ID yang menolak
//		CanceledBy uint           `gorm:"default:null"`                // User ID yang membatalkan
//		ApprovedAt time.Time      `gorm:"type:timestamp;default:null"` // Waktu persetujuan
//		RejectedAt time.Time      `gorm:"type:timestamp;default:null"` // Waktu penolakan
//		CanceledAt time.Time      `gorm:"type:timestamp;default:null"` // Waktu pembatalan
//		Note       string         `gorm:"type:text"`                   // Catatan tambahan
//		CreatedAt  time.Time      `gorm:"autoCreateTime"`              // Waktu pembuatan otomatis
//		UpdatedAt  time.Time      `gorm:"autoUpdateTime"`              // Waktu pembaruan otomatis
//		DeletedAt  gorm.DeletedAt `gorm:"index"`                       // Soft delete
//		CreatedBy  uint           `gorm:"not null"`                    // User ID yang membuat data
//		UpdatedBy  uint           `gorm:"default:null"`                // User ID yang terakhir memperbarui data
//		DeletedBy  uint           `gorm:"default:null"`                // User ID yang menghapus data
//	}
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

func CreateLeaveDetail(c echo.Context) error {
	var input struct {
		LeaveID   uint   `json:"leave_id" validate:"required"`
		Note      string `json:"note"`
		CreatedBy uint   `json:"created_by" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Validasi CompanyID dengan memanggil Division Service
	leaveURL := "http://localhost:8085/api/leaves/" + string(rune(input.LeaveID))
	if err := fetchService(leaveURL); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid Leave ID",
		})
	}

	leaveDetail := models.LeaveDetail{
		LeaveID:   input.LeaveID,
		Note:      input.Note,
		CreatedBy: input.CreatedBy,
	}

	if err := config.DB.Create(&leaveDetail).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, leaveDetail)
}

func UpdateLeaveDetail(c echo.Context) error {
	var leaveDetail models.LeaveDetail

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&leaveDetail).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}

	var input struct {
		Note      string `json:"note"`
		UpdatedBy uint   `json:"updated_by" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	leaveDetail.Note = input.Note
	leaveDetail.UpdatedBy = input.UpdatedBy

	if err := config.DB.Save(&leaveDetail).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, leaveDetail)
}

func DeleteLeaveDetail(c echo.Context) error {
	var leaveDetail models.LeaveDetail

	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&leaveDetail).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Record not found",
		})
	}

	if err := config.DB.Delete(&leaveDetail).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Deleted",
	})
}
