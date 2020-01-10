package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonatascabral/zipcodes-api/pkg/database"
	"github.com/jonatascabral/zipcodes-api/pkg/rabbitmq"
	"github.com/jonatascabral/zipcodes-api/pkg/routes"
	"github.com/jonatascabral/zipcodes-api/pkg/services"
	"log"
)

func main() {
	startMessage()

	log.Println("Connecting to database")
	database.Init()

	log.Println("Connecting to rabbit")
	rabbitmq.Init()

	request := gin.Default()

	log.Println("Loading application routes")
	routes.LoadRoutes(request)

	log.Println("Starting server at :9080")
	services.StartServer(request, ":9080")
}


func startMessage() {
	log.Println("-----------------------")
	log.Println("| ZipCodes API v0.0.1 |")
	log.Println("-----------------------")
}