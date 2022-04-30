package locations

import (
	"MyLogisticGame/entity"
	"github.com/labstack/echo/v4"
)

func CreateLocationsAPI(e *echo.Echo) {
	conn.AutoMigrate(&entity.Location{})

	e.GET("/locations", getLocations)
	e.POST("/locations", createLocation)
	e.GET("/locations/:id", getLocation)
	e.DELETE("/locations/:id", deleteLocation)
}
