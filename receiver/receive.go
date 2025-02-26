package main

import (
	"log"
	"rabbitmq/utils"
)

func main() {
	rabbitMQ := utils.GetRabbitMQ("hello")
	defer rabbitMQ.Conn.Close()

	ch := rabbitMQ.Chan
	defer ch.Close()

	q := rabbitMQ.Queue

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
