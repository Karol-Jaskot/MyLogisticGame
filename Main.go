package main

import (
	"MyLogisticGame/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", api.HomeResponse)

	// Locations
	e.GET("/locations", api.GetAllLocations)
	e.POST("/locations", api.CreateLocation)
	e.GET("/locations/:id", api.GetLocation)
	e.DELETE("/locations/:id", api.DeleteLocation)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
