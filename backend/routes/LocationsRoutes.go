package routes

import (
	"MyLogisticGame/backend/entity"
	"MyLogisticGame/backend/handlers"
	"github.com/labstack/echo/v4"
)

func CreateLocationsAPI(e *echo.Echo) {
	handlers.Conn.AutoMigrate(&entity.Location{})

	e.GET("/locations", handlers.GetLocations)
	e.POST("/locations", handlers.CreateLocation)
	e.GET("/locations/:id", handlers.GetLocation)
	e.DELETE("/locations/:id", handlers.DeleteLocation)
}
