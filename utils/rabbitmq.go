package utils

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitMQ struct {
	Conn  *amqp.Connection
	Chan  *amqp.Channel
	Queue amqp.Queue
}

func GetRabbitMQ(queueName string) *RabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	return &RabbitMQ{conn, ch, q}
}
