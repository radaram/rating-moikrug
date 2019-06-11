package main

import (
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Read() {
	var company Company
	defer mq.Close()
	defer db.Close()

	ch, err := mq.Channel()
	failOnError(err, "Failed to open a chanel")
	defer ch.Close()

	msgs, err := ch.Consume(
		COMPANIES_QUEUE,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	
	go func() {
		for d := range msgs {
			//log.Printf("Received a message: %s", d.Body)	
			company = Company{}
			err := company.Decode(d.Body)
			if err != nil {
				log.Println(err)
				continue
			}
			go calculate(company)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
