package locations

import (
	"MyLogisticGame/api"
	"MyLogisticGame/configs"
	"MyLogisticGame/entity"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var conn = configs.GetConnection()

// GetLocations godoc
// @Summary Get Locations
// @Tags locations
// @Accept */*
// @Produce json
// @Success 200 {object} []entity.Location
// @failure 500 {object} map[string]interface{}
// @Router /locations [get]
func getLocations(c echo.Context) error {
	var locations []entity.Location
	conn.Find(&locations)
	return c.JSON(http.StatusOK, locations)
}

// GetLocation godoc
// @Summary Get Location
// @Tags locations
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} entity.Location
// @failure 500 {string}  string    "ERROR"
// @Router /locations/{id} [get]
func getLocation(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	api.CheckErr(err)
	var loc entity.Location
	conn.First(&loc, id)
	return c.JSON(http.StatusOK, loc)
}

// GetLocations godoc
// @Summary Get Locations
// @Tags locations
// @Accept json
// @Produce json
// @Param location body entity.Location true "Add location"
// @Success 200 {object} entity.Location
// @failure 500 {object} map[string]interface{}
// @Router /locations [post]
func createLocation(c echo.Context) error {
	var loc entity.Location
	err := c.Bind(&loc)
	api.CheckErr(err)
	conn.Create(&loc)
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
func deleteLocation(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, err)
	}

	var count int
	conn.Raw("SELECT count(1) FROM locations WHERE id = ?", id).Scan(&count)

	if count > 0 {
		return c.String(http.StatusOK, fmt.Sprintf("Deleted location with id: %d", id))
	} else {
		return c.String(http.StatusOK, fmt.Sprintf("Location with id %d not exist", id))
	}
}
