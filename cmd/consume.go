package main

import (
	"github.com/jonatascabral/zipcodes-api/pkg/database"
	"github.com/jonatascabral/zipcodes-api/pkg/rabbitmq"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	log.Println("Connecting to database")
	database.Init()

	log.Println("Connecting to rabbit")
	rabbitmq.Init()

	rabbitmq.Consume("zipcodes", func(message amqp.Delivery) bool {
		log.Println(string(message.Body))
		return true
	})
}