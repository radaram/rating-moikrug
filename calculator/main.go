package main

import (
    _ "github.com/lib/pq"
    "database/sql"
	"github.com/streadway/amqp"
)


var db *sql.DB
var mq *amqp.Connection

func initdb() {
	var err error
	db, err = sql.Open(DATABASE_DRIVER_NAME, DATABASE)
	failOnError(err, "Failed to connect to Postgres")
	
	err = db.Ping()
    failOnError(err, "Failed to connect to Postgres")
}


func initmq() {
	var err error
	mq, err = amqp.Dial(RABBITMQ_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
}


func init() {
	initdb()
	initmq()
}


func main() {
	Read()
}
