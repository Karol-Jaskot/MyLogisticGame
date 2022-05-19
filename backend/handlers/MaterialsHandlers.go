package handlers

import (
	"MyLogisticGame/backend/data/materials"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetMaterials godoc
// @Summary Get Materials
// @Tags materials
// @Accept */*
// @Produce json
// @Success 200 {object} []materials.Material
// @failure 500 {object} echo.HTTPError
// @Router /materials [get]
func GetMaterials(c echo.Context) error {
	mat, err := materials.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, mat)
}

// CreateMaterial godoc
// @Summary Create Material
// @Tags materials
// @Accept json
// @Produce json
// @Param Material body materials.Material true "Add material"
// @Success 201 {object} materials.Material
// @failure 400 {object} echo.HTTPError
// @failure 500 {object} echo.HTTPError
// @Router /materials [post]
func CreateMaterial(c echo.Context) error {
	mat, err := materials.CreateFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = materials.Save(mat)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, mat)
}

// GetMaterial godoc
// @Summary Get Material
// @Tags materials
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} materials.Material
// @failure 400 {object} echo.HTTPError
// @failure 404 {object} echo.HTTPError
// @failure 500 {object} echo.HTTPError
// @Router /materials/{id} [get]
func GetMaterial(c echo.Context) error {
	id, err := ParseIdFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	mat, err := materials.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, mat)
}

// DeleteMaterial godoc
// @Summary Delete material
// @Tags materials
// @Accept */*
// @Produce json
// @Param id path int true "ID"
// @Success 204 {string} string
// @failure 405 {object} echo.HTTPError
// @failure 500 {object} echo.HTTPError
// @Router /materials/{id} [delete]
func DeleteMaterial(c echo.Context) error {
	id, err := ParseIdFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	mat, err := materials.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	err = materials.Delete(mat)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}
