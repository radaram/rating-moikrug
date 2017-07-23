package main

import (
	"sync"

	"data"
	"consumer"
	"crawler"
	"producer"
	"settings"
)


func main() {
	var wg sync.WaitGroup
	var c chan *data.Company = make(chan *data.Company)	
	wg.Add(1)
	go consumer.Read(&wg)
	wg.Add(1)
	go producer.Send(c, &wg)
	crawler.GrabCompanies(settings.COMPANIES_URL, c, &wg)
	wg.Wait()
}

