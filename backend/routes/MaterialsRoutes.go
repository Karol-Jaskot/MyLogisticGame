package routes

import (
	"MyLogisticGame/backend/entity"
	"MyLogisticGame/backend/handlers"
	"github.com/labstack/echo/v4"
)

func CreateMaterialsAPI(e *echo.Echo) {
	handlers.Conn.AutoMigrate(&entity.Material{})

	e.GET("/materials", handlers.GetMaterials)
	e.POST("/materials", handlers.CreateMaterial)
	e.GET("/materials/:id", handlers.GetMaterial)
	e.DELETE("/materials/:id", handlers.DeleteMaterial)
}
