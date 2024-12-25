package models

import (
	"time"

	"gorm.io/gorm"
)

type Leave struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"not null"`           // Foreign key ke User
	StartDate time.Time      `gorm:"type:date;not null"` // Waktu mulai cuti
	EndDate   time.Time      `gorm:"type:date;not null"` // Waktu selesai cuti
	Reason    string         `gorm:"type:text"`          // Alasan cuti
	Status    int            `gorm:"default:0"`          // Status cuti (0 = pending)
	Qty       int            `gorm:"not null"`           // Jumlah hari cuti
	CreatedAt time.Time      `gorm:"autoCreateTime"`     // Waktu pembuatan otomatis
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`     // Waktu pembaruan otomatis
	DeletedAt gorm.DeletedAt `gorm:"index"`              // Soft delete
	CreatedBy uint           `gorm:"not null"`           // User ID yang membuat data
	UpdatedBy uint           `gorm:"default:null"`       // User ID yang terakhir memperbarui data
	DeletedBy uint           `gorm:"default:null"`       // User ID yang menghapus data
}
