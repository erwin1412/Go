package models

import (
	"time"

	"gorm.io/gorm"
)

type Position struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"` // Waktu pembuatan otomatis
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"` // Waktu pembaruan otomatis
	DeletedAt   gorm.DeletedAt `gorm:"index"`          // Soft delete
	CreatedBy   uint           `gorm:"default:null"`   // User ID yang membuat data
	UpdatedBy   uint           `gorm:"default:null"`   // User ID yang terakhir memperbarui data
	DeletedBy   uint           `gorm:"default:null"`   // User ID yang menghapus data
}
