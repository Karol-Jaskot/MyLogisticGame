package companies

import (
	"MyLogisticGame/entity"
	"github.com/labstack/echo/v4"
)

func CreateCompaniesAPI(e *echo.Echo) {
	conn.AutoMigrate(&entity.Company{})

	e.GET("/companies", getCompanies)
	e.POST("/companies", createCompany)
	e.GET("/companies/:id", getCompany)
	e.DELETE("/companies/:id", deleteCompany)

	e.POST("/companies/:company_id/locations/:location_id", companyAddLocation)
}
