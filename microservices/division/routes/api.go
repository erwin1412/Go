package routes

import (
	"myapp/controllers"
	"myapp/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	api := e.Group("/api/divisions")
	api.Use(middleware.JWTMiddleware)

	// Routes for Division
	api.GET("", controllers.GetAllDivisions)
	api.GET("/:id", controllers.GetDivisionByID)
	api.POST("", controllers.CreateDivision)
	api.PUT("/:id", controllers.UpdateDivision)
	api.DELETE("/:id", controllers.DeleteDivision)

}
