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

	// Routes for Leave Detail
	api.GET("/leave-details", controllers.GetAllLeaveDetails)
	api.GET("/leave-details/:id", controllers.GetLeaveDetailByID)
	api.POST("/leave-details", controllers.CreateLeaveDetail)
	api.PUT("/leave-details/:id", controllers.UpdateLeaveDetail)
	api.DELETE("/leave-details/:id", controllers.DeleteLeaveDetail)

}
