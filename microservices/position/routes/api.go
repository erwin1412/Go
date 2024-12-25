package routes

import (
	"myapp/controllers"
	"myapp/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Group route dengan middleware JWT
	api := e.Group("/api/positions")
	api.Use(middleware.JWTMiddleware)

	api.GET("", controllers.GetAllPositions)       // GET /api/positions
	api.GET("/:id", controllers.GetPosition)       // GET /api/positions/:id
	api.POST("", controllers.CreatePosition)       // POST /api/positions
	api.PUT("/:id", controllers.UpdatePosition)    // PUT /api/positions/:id
	api.DELETE("/:id", controllers.DeletePosition) // DELETE /api/positions/:id
}
