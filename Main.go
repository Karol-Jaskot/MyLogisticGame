package main

import (
	"MyLogisticGame/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", homeResponse)

	api.CreateLocationsAPI(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func homeResponse(c echo.Context) error {
	c.String(http.StatusOK, "Hello, World!")
	return nil
}
