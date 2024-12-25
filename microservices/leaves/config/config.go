package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Sesuaikan `dsn` dengan kredensial PostgreSQL Anda
	dsn := "host=localhost user=postgres password=1234 dbname=absen_user_leaves port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	log.Println("Database connected successfully")
}
