package main

import (
	"fmt"
	"os"
)

var (
	BASE_URL      string = "https://moikrug.ru"
	COMPANIES_URL string = "https://moikrug.ru/companies?with_vacancies=1"
	
	RABBITMQ_USER string = os.Getenv("RABBITMQ_USER")
	RABBITMQ_PASSWORD string = os.Getenv("RABBITMQ_PASSWORD")
	RABBITMQ_HOST string = os.Getenv("RABBITMQ_HOST")
	RABBITMQ_PORT string = os.Getenv("RABBITMQ_PORT")

	COMPANIES_QUEUE = os.Getenv("COMPANIES_QUEUE")

	RABBITMQ_URL string = fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		RABBITMQ_USER,
		RABBITMQ_PASSWORD,
		RABBITMQ_HOST,
		RABBITMQ_PORT,
	)
)

