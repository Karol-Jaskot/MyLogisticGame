package materials

import (
	"MyLogisticGame/entity"
	"github.com/labstack/echo/v4"
)

func CreateMaterialsAPI(e *echo.Echo) {
	conn.AutoMigrate(&entity.Material{})

	e.GET("/materials", getMaterials)
	e.POST("/materials", createMaterial)
	e.GET("/materials/:id", getMaterial)
	e.DELETE("/materials/:id", deleteMaterial)
}
