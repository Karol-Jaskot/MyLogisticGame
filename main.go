package main

import (
	"MyLogisticGame/backend/routes"
	"MyLogisticGame/configs"
	_ "MyLogisticGame/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/swaggo/echo-swagger"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is my simple game with backend.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @schemes http
func main() {
	// Echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/*", healthCheck)

	routes.CreateLocationsAPI(e)
	routes.CreateCompaniesAPI(e)
	routes.CreateMaterialsAPI(e)
	routes.CreateVehiclesAPI(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":" + configs.ApiPort))
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
