package vehicles

import (
	"MyLogisticGame/entity"
	"github.com/labstack/echo/v4"
)

func CreateVehiclesAPI(e *echo.Echo) {
	conn.AutoMigrate(&entity.Vehicle{})

	e.GET("/vehicles", getVehicles)
	e.POST("/vehicles", createVehicle)
	e.GET("/vehicles/:id", getVehicle)
	e.DELETE("/vehicles/:id", deleteVehicle)
}
