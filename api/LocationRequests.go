package api

import (
	"MyLogisticGame/entity"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func HomeResponse(c echo.Context) error {
	c.String(http.StatusOK, "Hello, World!")
	return nil
}

func GetAllLocations(c echo.Context) error {
	encjson, _ := json.Marshal(entity.GetAllLocations())
	return c.String(http.StatusOK, string(encjson))
}

func GetLocation(c echo.Context) error {
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	var loc = entity.GetLocationById(id)
	encjson, _ := json.Marshal(loc)
	return c.String(http.StatusOK, string(encjson))
}

func CreateLocation(c echo.Context) error {
	var loc entity.Location
	loc.Name = c.FormValue("name")
	entity.SaveLocation(loc)
	return c.JSON(http.StatusCreated, loc)
}

func DeleteLocation(c echo.Context) error {
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	entity.DeleteLocation(id)
	mess := "Delete location with id: " + i
	return c.String(http.StatusOK, mess)
}
