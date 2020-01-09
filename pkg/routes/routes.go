package routes

import (
	"github.com/jonatascabral/zipcodes-api/pkg/controllers"
	"github.com/labstack/echo"
)

func LoadRoutes(e *echo.Echo) {
	e.GET("/", controllers.Ping)
	e.POST("/import/csv", controllers.ImportCsv)
}