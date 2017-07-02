package producer

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"

	"data"
	"parser/settings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Send(c chan *data.Company) {
	conn, err := amqp.Dial(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%s",
			settings.RABBITMQ_USER,
			settings.RABBITMQ_PASSWORD,
			settings.RABBITMQ_HOST,
			settings.RABBITMQ_PORT,
		),
	)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"companies", // queue name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for {
		select {
			case company := <- c:
			    body := company.ToJson()
				err = ch.Publish(
					"",
					q.Name,
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body: []byte(body),
					})
				failOnError(err, "Failed to publish a message")
		}
	}
}

