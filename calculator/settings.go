package main

import (
	"fmt"
	"os"
)

var (
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
   
    
	POSTGRES_USER string = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD string = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST string = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT string = os.Getenv("POSTGRES_PORT")
	POSTGRES_DB string = os.Getenv("POSTGRES_DB")

    DATABASE_DRIVER_NAME = "postgres"
	DATABASE string = fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		POSTGRES_USER,
		POSTGRES_PASSWORD,
		POSTGRES_HOST,
		POSTGRES_DB,
	)

)

