package companies

import (
	"MyLogisticGame/configs"
	"MyLogisticGame/entity"
	"fmt"
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
// @failure 404 {object} echo.HTTPError
// @Router /companies/{id} [get]
func getCompany(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var com entity.Company
	conn.First(&com, id)

	if com.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Company with ID %d not exist", id))
	}

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
// @failure 405 {object} echo.HTTPError
// @Router /companies [post]
func createCompany(c echo.Context) error {
	var com entity.Company
	err := c.Bind(&com)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Reset ID if exist
	com.ID = 0

	// Check unique GLN
	if com.Gln == 0 {
		return c.String(http.StatusMethodNotAllowed, fmt.Sprintf("GLN cannot be 0"))
	}
	var comGln entity.Company
	conn.Where(&entity.Company{Gln: com.Gln}).First(&comGln)
	if comGln.ID != 0 {
		return c.String(http.StatusMethodNotAllowed, fmt.Sprintf("Company with GLN %d is exist", com.Gln))
	}

	conn.Create(&com)
	return c.JSON(http.StatusCreated, com)
}

// DeleteCompany godoc
// @Summary Delete Company
// @Tags companies
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {string} string
// @Success 204 {string} string
// @failure 400 {object} echo.HTTPError
// @Router /companies/{id} [delete]
func deleteCompany(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var com entity.Company
	conn.First(&com, id)

	if com.ID == 0 {
		return c.String(http.StatusNoContent, fmt.Sprintf("Company with ID %d not exist", id))
	} else {
		// Clear reference from locations
		for _, loc := range com.Locations {
			loc.CompanyRefer = 0
			conn.Save(loc)
		}

		conn.Delete(com)
		return c.String(http.StatusOK, fmt.Sprintf("Deleted company with ID: %d", id))
	}

}

// AssignLocationToCompany godoc
// @Summary Assign Location To Company
// @Tags companies
// @Accept */*
// @Produce json
// @Param companyId path int true "Company ID"
// @Param locationId path int true "Location ID"
// @Success 200 {object} entity.Company
// @failure 400 {object} echo.HTTPError
// @failure 405 {object} echo.HTTPError
// @Router /companies/{companyId}/locations/{locationId} [post]
func companyAssignLocation(c echo.Context) error {
	companyId, err := strconv.Atoi(c.Param("company_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	locationId, err := strconv.Atoi(c.Param("location_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if companyId == 0 || locationId == 0 {
		return c.String(http.StatusMethodNotAllowed, fmt.Sprintf("Company or location ID cannot be 0"))
	}

	// Get and check location
	var loc entity.Location
	conn.First(&loc, locationId)
	if loc.ID == 0 {
		return c.String(http.StatusMethodNotAllowed, fmt.Sprintf("Location with ID %d not exist", locationId))
	} else if loc.CompanyRefer != 0 && (int(loc.CompanyRefer)) != companyId {
		return c.String(http.StatusMethodNotAllowed, fmt.Sprintf("This location is assign to other company"))
	}

	// Get and check company
	var com entity.Company
	conn.Preload("Locations").First(&com, companyId)
	if com.ID == 0 {
		return c.String(http.StatusMethodNotAllowed, fmt.Sprintf("Company with ID %d not exist", companyId))
	}

	// Init list if is empty
	if len(com.Locations) == 0 {
		com.Locations = []entity.Location{}
	}

	com.Locations = append(com.Locations, loc)
	conn.Save(&com)

	return c.JSON(http.StatusOK, com)
}
