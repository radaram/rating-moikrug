package main

import (
	"log"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func grabCompanies(url string, c chan *company, wg *sync.WaitGroup) {
	
	log.Println(url)
	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
	}
    
	doc.Find(".companies-item").Each(func(i int, s *goquery.Selection) {
		company_page, _ := s.Find(".title").Attr("href")
		company_page = BASE_URL + company_page
	    wg.Add(1)
		go grabCompany(company_page, c, wg)
    })

	next_page, _ := doc.Find(".pagination .next_page").Attr("href")
	if (next_page != "") {
		grabCompanies(BASE_URL + next_page, c, wg)
	}
}


func grabCompany(url string, c chan *company, wg *sync.WaitGroup) {
	defer wg.Done()
	
	log.Println(url)
	doc, err := goquery.NewDocument(url)

    if err != nil {
        log.Fatal(err)
    }

    name := doc.Find(".company_name a").Text()
    about := doc.Find(".company_about").Text()
    address := doc.Find(".address").Text()

    company := company{name, url, about, address, 0}
    c <- &company
}

