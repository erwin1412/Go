package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text"`
	Division    []Division     `gorm:"foreignKey:CompanyID"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"` // Waktu pembuatan otomatis
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"` // Waktu pembaruan otomatis
	DeletedAt   gorm.DeletedAt `gorm:"index"`          // Soft delete
	CreatedBy   uint           `gorm:"not null"`       // User ID yang membuat data
	UpdatedBy   uint           `gorm:"default:null"`   // User ID yang terakhir memperbarui data
	DeletedBy   uint           `gorm:"default:null"`   // User ID yang menghapus data
}

type CompanyResponse struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   uint      `json:"created_by"`
	UpdatedBy   uint      `json:"updated_by"`
}
