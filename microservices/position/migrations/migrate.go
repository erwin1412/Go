package migrations

import (
	"myapp/config"
	"myapp/models"
)

func RunMigrations() {
	// Jalankan migrasi untuk semua model
	config.DB.AutoMigrate(
		&models.Position{},
	)
}
