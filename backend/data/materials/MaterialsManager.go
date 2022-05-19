package materials

import (
	"MyLogisticGame/configs"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
)

var conn = configs.GetConnection()

func CreateFromContext(c echo.Context) (Material, error) {
	var mat Material
	err := c.Bind(&mat)

	// Reset ID if exist
	mat.ID = 0

	return mat, err
}

func Save(mat Material) error {
	err := conn.Create(&mat).Error
	return err
}

func GetAll() ([]Material, error) {
	var materials []Material
	err := conn.Find(&materials).Error
	return materials, err
}

func GetById(id int) (Material, error) {
	var mat Material
	err := conn.First(&mat, id).Error

	if mat.ID == 0 {
		return mat, errors.New(fmt.Sprintf(" material with ID %d not exist", id))
	}

	return mat, err
}

func Delete(mat Material) error {
	return conn.Delete(mat).Error
}
