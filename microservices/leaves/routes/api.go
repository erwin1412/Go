package routes

import (
	"myapp/controllers"
	"myapp/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Group route dengan middleware JWT
	api := e.Group("/api/leaves")
	api.Use(middleware.JWTMiddleware)

	// Routes for Leave
	api.GET("", controllers.GetAllLeaves)
	api.GET("/:id", controllers.GetLeaveByID)
	api.POST("", controllers.CreateLeave)
	api.PUT("/:id", controllers.UpdateLeave)
	api.DELETE("/:id", controllers.DeleteLeave)

}
