package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jonatascabral/zipcodes-api/pkg/controllers"
)

func LoadRoutes(request *gin.Engine) {
	request.GET("/", controllers.Ping)

	request.POST("/import/csv", controllers.ImportCsv)
}