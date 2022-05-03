package companies

import (
	"MyLogisticGame/api"
	"MyLogisticGame/configs"
	"MyLogisticGame/entity"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var conn = configs.GetConnection()

// GetCompanies godoc
// @Summary Get Companies
// @Tags companies
// @Accept */*
// @Produce json
// @Success 200 {object} []entity.Company
// @Router /companies [get]
func getCompanies(c echo.Context) error {
	var companies []entity.Company
	conn.Preload("Locations").Find(&companies)
	return c.JSON(http.StatusOK, companies)
}

// GetCompany godoc
// @Summary Get Company
// @Tags companies
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} entity.Company
// @failure 400 {object} echo.HTTPError
// @Router /companies/{id} [get]
func getCompany(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var com entity.Company
	conn.First(&com, id)
	return c.JSON(http.StatusOK, com)
}

// CreateCompany godoc
// @Summary Create Company
// @Tags companies
// @Accept json
// @Produce json
// @Param company body entity.Company  true  "Add company"
// @Success 201 {object} entity.Company
// @failure 400 {object} echo.HTTPError
// @Router /companies [post]
func createCompany(c echo.Context) error {
	var com entity.Company
	err := c.Bind(&com)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	conn.Create(&com)
	return c.JSON(http.StatusCreated, com)
}

// DeleteCompanies godoc
// @Summary Delete Companies
// @Tags companies
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} map[string]interface{}
// @failure 500 {object} map[string]interface{}
// @Router /companies/{id} [delete]
func deleteCompany(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	api.CheckErr(err)
	conn.Delete(&entity.Company{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Delete company with id: " + string(rune(id)),
	})
}

// AddLocationToCompany godoc
// @Summary Add Location To Company
// @Tags companies
// @Accept */*
// @Produce json
// @Param companyId path int true "Company ID"
// @Param locationId path int true "Location ID"
// @Success 200 {object} entity.Company
// @failure 500 {object} map[string]interface{}
// @Router /companies/{companyId}/locations/{locationId} [post]
func companyAddLocation(c echo.Context) error {
	companyId := c.Param("company_id")
	locationId := c.Param("location_id")

	var loc entity.Location
	conn.First(&loc, locationId)

	var com entity.Company
	conn.Preload("Locations").First(&com, companyId)

	if len(com.Locations) == 0 {
		com.Locations = []entity.Location{}
	}

	com.Locations = append(com.Locations, loc)
	conn.Save(&com)

	return c.JSON(http.StatusOK, com)
}
