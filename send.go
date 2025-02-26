package main

import (
	"context"
	"log"
	"rabbitmq/utils"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitMQ := utils.GetRabbitMQ("hello")
	defer rabbitMQ.Conn.Close()

	ch := rabbitMQ.Chan
	defer ch.Close()

	q := rabbitMQ.Queue

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err := ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
