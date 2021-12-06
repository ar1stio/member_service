package config

import (
	"member-service/exception"
	"strings"

	"github.com/streadway/amqp"
)

func NewRabbitMq(configuration Config) *amqp.Channel {
	allconnection := configuration.Get("RABBITMQ_URIALL")
	connecarr := strings.Split(allconnection, ",")
	for _, connecarr := range connecarr {
		conn, err := amqp.Dial(connecarr)

		if err == nil {
			ch, _ := conn.Channel()

			return ch
		}

	}
	exception.PanicIfNeeded("connection rabbit mq error")

	conn, _ := amqp.Dial("")

	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	// defer conn.Close()

	ch, _ := conn.Channel()

	// log.Println("error connection", err)

	// defer ch.Close()
	return ch
}

func NewRabbitConsumeMq(configuration Config) *amqp.Connection {
	allconnection := configuration.Get("RABBITMQ_URIALL")
	connecarr := strings.Split(allconnection, ",")
	for _, connecarr := range connecarr {
		conn, err := amqp.Dial(connecarr)

		if err == nil {
			return conn
		}

		exception.PanicIfNeeded(err)
	}

	conn, err := amqp.Dial("")
	exception.PanicIfNeeded(err)

	// defer ch.Close()
	return conn
}
