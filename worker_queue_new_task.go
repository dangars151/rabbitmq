package main

import (
	"context"
	"log"
	"os"
	"rabbitmq/utils"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitMQ := utils.GetRabbitMQ("task_queue", true)
	defer rabbitMQ.Conn.Close()

	ch := rabbitMQ.Chan
	defer ch.Close()

	q := rabbitMQ.Queue

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := utils.BodyFrom(os.Args)
	err := ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
