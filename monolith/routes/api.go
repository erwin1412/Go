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

	// Routes for Company
	protected.GET("/companies", controllers.FindCompanies)
	protected.GET("/companies/paginated", controllers.FindCompaniesWithPagination)
	protected.GET("/companies/:id", controllers.FindOneCompany)
	protected.POST("/companies", controllers.CreateCompany)
	protected.PUT("/companies/:id", controllers.UpdateCompany)
	protected.DELETE("/companies/:id", controllers.DeleteCompany)

	// Routes for Division
	protected.GET("/divisions", controllers.GetAllDivisions)
	protected.GET("/divisions/:id", controllers.GetDivisionByID)
	protected.POST("/divisions", controllers.CreateDivision)
	protected.PUT("/divisions/:id", controllers.UpdateDivision)
	protected.DELETE("/divisions/:id", controllers.DeleteDivision)

	// Routes for Position
	protected.GET("/positions", controllers.GetAllPositions)
	protected.GET("/positions/:id", controllers.GetPosition)
	protected.POST("/positions", controllers.CreatePosition)
	protected.PUT("/positions/:id", controllers.UpdatePosition)
	protected.DELETE("/positions/:id", controllers.DeletePosition)

	// Routes for Leave
	protected.GET("/leaves", controllers.GetAllLeaves)
	protected.GET("/leaves/:id", controllers.GetLeaveByID)
	protected.POST("/leaves", controllers.CreateLeave)
	protected.PUT("/leaves/:id", controllers.UpdateLeave)
	protected.DELETE("/leaves/:id", controllers.DeleteLeave)

	// Routes for Leave Detail
	protected.GET("/leave-details", controllers.GetAllLeaveDetails)
	protected.GET("/leave-details/:id", controllers.GetLeaveDetailByID)
	protected.POST("/leave-details", controllers.CreateLeaveDetail)
	protected.PUT("/leave-details/:id", controllers.UpdateLeaveDetail)
	protected.DELETE("/leave-details/:id", controllers.DeleteLeaveDetail)

}
