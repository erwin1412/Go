package migrations

import (
	"myapp/config"
	"myapp/models"
)

func RunMigrations() {
	// Jalankan migrasi untuk semua model
	config.DB.AutoMigrate(
		&models.Company{},
		&models.Division{},
		&models.Position{},
		&models.User{},
		&models.Leave{},
		&models.LeaveDetail{},
	)
}
