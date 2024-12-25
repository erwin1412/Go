package main

import (
	"myapp/config"
	"myapp/migrations"
	"myapp/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:8081",
			"http://localhost:8082",
			"http://localhost:8083",
			"http://localhost:8084",
			"http://localhost:8085",
			"http://localhost:8086",
		}, // Ganti dengan origin klien Anda
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	// Connect to database
	config.ConnectDatabase()

	// Run migrations
	migrations.RunMigrations()

	// Setup routes
	routes.SetupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8083"))
}
