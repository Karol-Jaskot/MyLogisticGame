package handlers

import (
	"MyLogisticGame/backend/entity"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetMaterials godoc
// @Summary Get Materials
// @Tags materials
// @Accept */*
// @Produce json
// @Success 200 {object} []entity.Material
// @Router /materials [get]
func GetMaterials(c echo.Context) error {
	var materials []entity.Material
	Conn.Find(&materials)
	return c.JSON(http.StatusOK, materials)
}

// CreateMaterial godoc
// @Summary Create Material
// @Tags materials
// @Accept json
// @Produce json
// @Param Material body entity.Material true "Add material"
// @Success 201 {object} entity.Material
// @failure 400 {object} echo.HTTPError
// @Router /materials [post]
func CreateMaterial(c echo.Context) error {
	var mat entity.Material

	err := c.Bind(&mat)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Reset ID if exist
	mat.ID = 0

	Conn.Create(&mat)
	return c.JSON(http.StatusCreated, mat)
}

// GetMaterial godoc
// @Summary Get Material
// @Tags materials
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} entity.Material
// @failure 400 {object} echo.HTTPError
// @failure 404 {object} echo.HTTPError
// @Router /materials/{id} [get]
func GetMaterial(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var mat entity.Material
	Conn.First(&mat, id)

	if mat.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Material with ID %d not exist", id))
	}

	return c.JSON(http.StatusOK, mat)
}

// DeleteMaterial godoc
// @Summary Delete material
// @Tags materials
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {string} string
// @Success 204 {string} string
// @failure 405 {object} echo.HTTPError
// @Router /materials/{id} [delete]
func DeleteMaterial(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var mat entity.Material
	Conn.First(&mat, id)

	if mat.ID == 0 {
		return c.String(http.StatusNoContent, fmt.Sprintf("Material with ID %d not exist", id))
	} else {
		Conn.Delete(mat)
		return c.String(http.StatusOK, fmt.Sprintf("Deleted material with ID: %d", id))
	}
}
