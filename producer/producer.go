package producer

import (
	"log"
	"sync"

	"github.com/streadway/amqp"

	"data"
	"settings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Send(c chan *data.Company, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := amqp.Dial(settings.RABBITMQ_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		settings.COMPANIES_QUEUE, // queue name
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

