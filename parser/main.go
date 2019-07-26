package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/robfig/cron"
)

func parser() {
	fmt.Println("Start parser...")
	var wg sync.WaitGroup
	var c chan *Company = make(chan *Company)
	wg.Add(1)
	go send(c, &wg)
	grabCompanies(COMPANIES_URL, c, &wg)
	wg.Wait()
}

func main() {
	c := cron.New()
	err := c.AddFunc("0 */1 * * * *", parser)
	if err != nil {
		log.Fatalf("Error AddFunc: %s", err)
	}
	c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
