package routes

import (
	"github.com/jonatascabral/zipcodes-api/pkg/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LoadRoutes(e *echo.Echo) {
	e.GET("/", controllers.Ping)

	g := e.Group("/import", middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == "test" && password == "test" {
			return true, nil
		}
		return false, nil
	}))
	{
		g.POST("/csv", controllers.ImportCsv)
	}
}