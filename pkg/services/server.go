package services

import (
	"github.com/labstack/echo"
)

func StartServer(e *echo.Echo, port string) {
	e.Logger.Fatal(e.Start(port))
}
