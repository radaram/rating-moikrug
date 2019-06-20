package main

import (
	"fmt"
	"os"
)

var ( 
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

