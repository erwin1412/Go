package routes

import (
	"myapp/controllers"
	"myapp/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Group route dengan middleware JWT
	api := e.Group("/api/companies")
	api.Use(middleware.JWTMiddleware)

	// Routes for Company
	api.GET("", controllers.FindCompanies)
	api.GET("/paginated", controllers.FindCompaniesWithPagination)
	api.GET("/:id", controllers.FindOneCompany)
	api.POST("", controllers.CreateCompany)
	api.PUT("/:id", controllers.UpdateCompany)
	api.DELETE("/:id", controllers.DeleteCompany)

}
