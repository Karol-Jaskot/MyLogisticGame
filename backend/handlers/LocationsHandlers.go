package handlers

import (
	"MyLogisticGame/backend/entity"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetLocations godoc
// @Summary Get Locations
// @Tags locations
// @Accept */*
// @Produce json
// @Success 200 {object} []entity.Location
// @Router /locations [get]
func GetLocations(c echo.Context) error {
	var locations []entity.Location
	Conn.Find(&locations)
	return c.JSON(http.StatusOK, locations)
}

// GetLocation godoc
// @Summary Get Location
// @Tags locations
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} entity.Location
// @failure 400 {object} echo.HTTPError
// @failure 404 {object} echo.HTTPError
// @Router /locations/{id} [get]
func GetLocation(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var loc entity.Location
	Conn.First(&loc, id)

	if loc.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Location with ID %d not exist", id))
	}

	return c.JSON(http.StatusOK, loc)
}

// CreateLocation godoc
// @Summary Create Location
// @Tags locations
// @Accept json
// @Produce json
// @Param location body entity.Location true "Add location"
// @Success 201 {object} entity.Location
// @failure 400 {object} echo.HTTPError
// @failure 405 {object} echo.HTTPError
// @Router /locations [post]
func CreateLocation(c echo.Context) error {
	var loc entity.Location

	err := c.Bind(&loc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Reset ID if exist
	loc.ID = 0

	// Check the company if it is set
	if loc.CompanyRefer != 0 {
		var com entity.Company
		Conn.First(&com, loc.CompanyRefer)
		if com.ID == 0 {
			return c.String(http.StatusMethodNotAllowed, fmt.Sprintf("Company with ID %d not exist", loc.CompanyRefer))
		}
	}

	Conn.Create(&loc)
	return c.JSON(http.StatusCreated, loc)
}

// DeleteLocation godoc
// @Summary Delete Location
// @Tags locations
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {string} string
// @Success 204 {string} string
// @failure 405 {object} echo.HTTPError
// @Router /locations/{id} [delete]
func DeleteLocation(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var loc entity.Location
	Conn.First(&loc, id)

	if loc.ID == 0 {
		return c.String(http.StatusNoContent, fmt.Sprintf("Location with ID %d not exist", id))
	} else if loc.CompanyRefer != 0 {
		return c.String(http.StatusMethodNotAllowed, fmt.Sprintf("Location with ID %d is assing to company with ID %d", id, loc.CompanyRefer))
	} else {
		Conn.Delete(loc)
		return c.String(http.StatusOK, fmt.Sprintf("Deleted location with ID: %d", id))
	}
}
