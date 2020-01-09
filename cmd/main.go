package main

import (
	"github.com/jonatascabral/zipcodes-api/pkg/database"
	"github.com/jonatascabral/zipcodes-api/pkg/rabbitmq"
	"github.com/jonatascabral/zipcodes-api/pkg/routes"
	"github.com/jonatascabral/zipcodes-api/pkg/services"
	"github.com/labstack/echo"
	"log"
)

func main() {
	startMessage()

	log.Println("Connecting to database")
	database.Init()

	log.Println("Connecting to rabbit")
	rabbitmq.Init()

	e := echo.New()

	log.Println("Loading application routes")
	routes.LoadRoutes(e)

	log.Println("Starting server at :9080")
	services.StartServer(e, ":9080")
}


func startMessage() {
	log.Println("-----------------------")
	log.Println("| ZipCodes API v0.0.1 |")
	log.Println("-----------------------")
}