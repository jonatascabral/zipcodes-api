package services

import (
	"github.com/gin-gonic/gin"
)

func StartServer(e *gin.Engine, port string) {
	e.Run(port)
}
