package main

import (
	"encoding/json"
	"github.com/jonatascabral/zipcodes-api/pkg/database"
	"github.com/jonatascabral/zipcodes-api/pkg/models"
	"github.com/jonatascabral/zipcodes-api/pkg/rabbitmq"
	"github.com/jonatascabral/zipcodes-api/pkg/services"
	"github.com/streadway/amqp"
	"log"
)

var (
	ACK = true
	NACK = false
)

func main() {
	log.Println("Connecting to database")
	dbConfig := database.Init()

	log.Println("Connecting to rabbit")
	rabbitmq.Init()

	rabbitmq.Consume("zipcodes", func(message amqp.Delivery) bool {
		var address *models.Address
		json.Unmarshal([]byte(message.Body), &address)
		_, err := services.GetAddress(address)

		if err != nil {
			return NACK
		}

		dbConfig.DB.Create(address)

		return ACK
	})
}