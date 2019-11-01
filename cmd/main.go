package main

import (
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

	log.Println("Loading application routes")
	routes.LoadRoutes()

	log.Println("Starting server at :9080")
	services.StartServer(":9080")
}


func startMessage() {
	log.Println("-----------------------")
	log.Println("| ZipCodes API v0.0.1 |")
	log.Println("-----------------------")
}