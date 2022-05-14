package handlers

import (
	"MyLogisticGame/backend/entity"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetVehicles godoc
// @Summary Get Vehicles
// @Tags vehicles
// @Accept */*
// @Produce json
// @Success 200 {object} []entity.Vehicle
// @Router /vehicles [get]
func GetVehicles(c echo.Context) error {
	var vehicles []entity.Vehicle
	Conn.Find(&vehicles)
	return c.JSON(http.StatusOK, vehicles)
}

// CreateVehicle godoc
// @Summary Create Vehicle
// @Tags vehicles
// @Accept json
// @Produce json
// @Param Vehicle body entity.Vehicle true "Add vehicle"
// @Success 201 {object} entity.Vehicle
// @failure 400 {object} echo.HTTPError
// @Router /vehicles [post]
func CreateVehicle(c echo.Context) error {
	var veh entity.Vehicle

	err := c.Bind(&veh)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Reset ID if exist
	veh.ID = 0

	Conn.Create(&veh)
	return c.JSON(http.StatusCreated, veh)
}

// GetVehicle godoc
// @Summary Get Vehicle
// @Tags vehicles
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} entity.Vehicle
// @failure 400 {object} echo.HTTPError
// @failure 404 {object} echo.HTTPError
// @Router /vehicles/{id} [get]
func GetVehicle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var veh entity.Vehicle
	Conn.First(&veh, id)

	if veh.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Vehicle with ID %d not exist", id))
	}

	return c.JSON(http.StatusOK, veh)
}

// DeleteVehicle godoc
// @Summary Delete vehicle
// @Tags vehicles
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {string} string
// @Success 204 {string} string
// @failure 405 {object} echo.HTTPError
// @Router /vehicles/{id} [delete]
func DeleteVehicle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var veh entity.Vehicle
	Conn.First(&veh, id)

	if veh.ID == 0 {
		return c.String(http.StatusNoContent, fmt.Sprintf("Vehicle with ID %d not exist", id))
	} else {
		Conn.Delete(veh)
		return c.String(http.StatusOK, fmt.Sprintf("Deleted vehicle with ID: %d", id))
	}
}
