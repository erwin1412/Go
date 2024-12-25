package routes

import (
	"myapp/controllers"
	"myapp/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api")

	// Routes for Auth (Tanpa middleware JWT)
	api.POST("/login", controllers.Login)
	api.POST("/register", controllers.Register)

	// Protected routes (Dengan middleware JWT)
	protected := api.Group("")              // Buat grup baru untuk rute terproteksi
	protected.Use(middleware.JWTMiddleware) // Terapkan middleware di grup ini
}
