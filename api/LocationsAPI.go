package api

import (
	"MyLogisticGame/entity"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// CreateLocationsAPI ==========================================================
func CreateLocationsAPI(e *echo.Echo) {
	// Routes
	e.GET("/locations", getLocations)
	e.POST("/locations", createLocation)
	e.GET("/locations/:id", getLocation)
	e.DELETE("/locations/:id", deleteLocation)
}

// ==========================================================
// Handlers
func getLocations(c echo.Context) error {
	encjson, _ := json.Marshal(entity.GetAllLocations())
	return c.String(http.StatusOK, string(encjson))
}

func getLocation(c echo.Context) error {
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	var loc = entity.GetLocationById(id)
	encjson, _ := json.Marshal(loc)
	return c.String(http.StatusOK, string(encjson))
}

func createLocation(c echo.Context) error {
	var loc entity.Location
	loc.Name = c.FormValue("name")
	entity.SaveLocation(loc)
	return c.JSON(http.StatusCreated, loc)
}

func deleteLocation(c echo.Context) error {
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	entity.DeleteLocation(id)
	mess := "Delete location with id: " + i
	return c.String(http.StatusOK, mess)
}
