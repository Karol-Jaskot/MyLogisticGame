package routes

import (
	"MyLogisticGame/backend/entity"
	"MyLogisticGame/backend/handlers"
	"github.com/labstack/echo/v4"
)

func CreateVehiclesAPI(e *echo.Echo) {
	handlers.Conn.AutoMigrate(&entity.Vehicle{})

	e.GET("/vehicles", handlers.GetVehicles)
	e.POST("/vehicles", handlers.CreateVehicle)
	e.GET("/vehicles/:id", handlers.GetVehicle)
	e.DELETE("/vehicles/:id", handlers.DeleteVehicle)
}
