package models

import (
	"time"

	"gorm.io/gorm"
)

type LeaveDetail struct {
	ID         uint           `gorm:"primaryKey"`
	LeaveID    uint           `gorm:"not null"`                    // Foreign key ke Leave
	ApprovedBy uint           `gorm:"default:null"`                // User ID yang menyetujui
	RejectedBy uint           `gorm:"default:null"`                // User ID yang menolak
	CanceledBy uint           `gorm:"default:null"`                // User ID yang membatalkan
	ApprovedAt time.Time      `gorm:"type:timestamp;default:null"` // Waktu persetujuan
	RejectedAt time.Time      `gorm:"type:timestamp;default:null"` // Waktu penolakan
	CanceledAt time.Time      `gorm:"type:timestamp;default:null"` // Waktu pembatalan
	Note       string         `gorm:"type:text"`                   // Catatan tambahan
	CreatedAt  time.Time      `gorm:"autoCreateTime"`              // Waktu pembuatan otomatis
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`              // Waktu pembaruan otomatis
	DeletedAt  gorm.DeletedAt `gorm:"index"`                       // Soft delete
	CreatedBy  uint           `gorm:"not null"`                    // User ID yang membuat data
	UpdatedBy  uint           `gorm:"default:null"`                // User ID yang terakhir memperbarui data
	DeletedBy  uint           `gorm:"default:null"`                // User ID yang menghapus data
}
