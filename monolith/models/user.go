package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"type:varchar(255);not null"`        // Nama user
	Email       string         `gorm:"type:varchar(255);unique;not null"` // Email unik
	Password    string         `gorm:"type:varchar(255);not null"`        // Password (hashed)
	LeavesQty   int            `gorm:"default:0"`                         // Jumlah cuti
	JoinDate    string         `gorm:"type:date"`                         // Tanggal bergabung
	Status      int            `gorm:"default:1"`                         // Status user
	PhoneNumber string         `gorm:"type:varchar(15)"`                  // Nomor telepon
	DivisionID  uint           `gorm:"not null"`                          // Foreign key ke Division
	Division    Division       `gorm:"foreignKey:DivisionID"`             // Relasi ke Division
	PositionID  uint           `gorm:"not null"`                          // Foreign key ke Position
	Position    Position       `gorm:"foreignKey:PositionID"`             // Relasi ke Position
	Leaves      []Leave        `gorm:"foreignKey:UserID"`                 // One-to-Many ke Leave
	CreatedAt   time.Time      `gorm:"autoCreateTime"`                    // Waktu pembuatan otomatis
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`                    // Waktu pembaruan otomatis
	DeletedAt   gorm.DeletedAt `gorm:"index"`                             // Soft delete
	CreatedBy   uint           `gorm:"not null"`                          // User ID yang membuat data
	UpdatedBy   uint           `gorm:"default:null"`                      // User ID yang terakhir memperbarui data
	DeletedBy   uint           `gorm:"default:null"`                      // User ID yang menghapus data
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// GenerateToken membuat JWT token
func (u *User) GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"user_id":     u.ID,
		"username":    u.Name,
		"division_id": u.DivisionID,
		"position_id": u.PositionID,
		"exp":         time.Now().Add(time.Hour * 72).Unix(), // Token berlaku 72 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key"))
}

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err

}
