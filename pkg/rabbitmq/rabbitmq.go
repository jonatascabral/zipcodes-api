package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	QueueName string
	Connection *amqp.Connection
}

type callback func (message amqp.Delivery) bool

var RabbitConfig *Config

func Connect(host string, port string, user string, password string) *amqp.Connection {
	connection, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%s/%%2Fzipcodes",
		user,
		password,
		host,
		port))
	if err != nil {
		panic(err)
	}

	return connection
}

func Publish(queueName string, body string) (*Config, error) {
	channel, err := RabbitConfig.Connection.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		panic(err)
	}

	err = channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(body),
			DeliveryMode: uint8(2)})
	channel.Close()

	if err != nil {
		log.Printf("Error publising on queue '%s'", queueName)
	} else {
		log.Printf("Message published on queue '%s'", queueName)
	}

	return RabbitConfig, err
}

func Consume(queueName string, consumerCallback callback) {
	channel, err := RabbitConfig.Connection.Channel()
	if err != nil {
		panic(err)
	}
	consumerTag := fmt.Sprintf("%s_%d", queueName, rand.Uint32())

	messages, err := channel.Consume(
		queueName,
		consumerTag,
		false,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		panic(err)
	}
	for message := range messages {
		if consumerCallback(message) {
			channel.Ack(message.DeliveryTag, false)
		} else {
			channel.Nack(message.DeliveryTag, false, true)
		}
	}
	channel.Close()
}

func Init() {
	RabbitConfig = &Config{
		Host: "localhost",
		Port: "5672",
		User: "admin",
		Password: "admin"}

	RabbitConfig.Connection = Connect(
		RabbitConfig.Host,
		RabbitConfig.Port,
		RabbitConfig.User,
		RabbitConfig.Password)

	log.Println("Rabbitmq connected")
}