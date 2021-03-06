package rabbitmq

import (
	"errors"
	"hcc/violin/lib/config"
	"hcc/violin/lib/logger"
	"strconv"

	"github.com/streadway/amqp"
)

// Connection : RabbitMQ connection variable
var Connection *amqp.Connection

// Channel : RabbitMQ channel variable
var Channel *amqp.Channel

func prepareChannel() error {
	Connection, err := amqp.Dial("amqp://" + config.RabbitMQ.ID + ":" + config.RabbitMQ.Password + "@" +
		config.RabbitMQ.Address + ":" + strconv.Itoa(int(config.RabbitMQ.Port)))
	if err != nil {
		return errors.New("failed to connect to RabbitMQ server")
	}
	logger.Logger.Println("Connected to RabbitMQ server")

	Channel, err = Connection.Channel()
	if err != nil {
		return errors.New("failed to open a RabbitMQ channel")
	}
	logger.Logger.Println("Opened RabbitMQ channel.")

	return nil
}
