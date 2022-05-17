package routes

import (
	"MyLogisticGame/backend/entity"
	"MyLogisticGame/backend/handlers"
	"github.com/labstack/echo/v4"
)

func CreateCompaniesAPI(e *echo.Echo) {
	handlers.Conn.AutoMigrate(&entity.Company{})

	e.GET("/companies", handlers.GetCompanies)
	e.POST("/companies", handlers.CreateCompany)
	e.GET("/companies/:id", handlers.GetCompany)
	e.DELETE("/companies/:id", handlers.DeleteCompany)

	e.POST("/companies/:company_id/locations/:location_id", handlers.CompanyAssignLocation)
	e.DELETE("/companies/:company_id/locations/:location_id", handlers.CompanyRemoveLocation)
}
