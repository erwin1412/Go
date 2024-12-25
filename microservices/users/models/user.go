package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"type:varchar(255);not null"`
	Email       string         `gorm:"type:varchar(255);unique;not null"`
	Password    string         `gorm:"type:varchar(255);not null"`
	LeavesQty   int            `gorm:"default:0"`
	JoinDate    string         `gorm:"type:date"`
	Status      int            `gorm:"default:1"`
	PhoneNumber string         `gorm:"type:varchar(15)"`
	DivisionID  uint           `gorm:"null"` // Hanya menyimpan ID referensi ke Division Service
	PositionID  uint           `gorm:"null"` // Hanya menyimpan ID referensi ke Position Service
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	CreatedBy   uint           `gorm:"default:null"`
	UpdatedBy   uint           `gorm:"default:null"`
	DeletedBy   uint           `gorm:"default:null"`
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

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
