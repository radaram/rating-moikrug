package main

import (
    _ "github.com/lib/pq"
    "database/sql"
	"log"
	"net/http"
)


var db *sql.DB


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


func initdb() {
	var err error
	db, err = sql.Open(DATABASE_DRIVER_NAME, DATABASE)
	failOnError(err, "Failed to connect to Postgres")
	
	err = db.Ping()
    failOnError(err, "Failed to connect to Postgres")
}


func init() {
	initdb()
}


func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7000", nil)
}
