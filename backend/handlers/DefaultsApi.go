package handlers

import (
	"MyLogisticGame/configs"
	"errors"
	"github.com/labstack/echo/v4"
	"strconv"
)

var Conn = configs.GetConnection()

func ParseIdFromContext(c echo.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		if id < 0 {
			err = errors.New(" id cannot by lower that 1")
		}
	}
	return id, err
}
