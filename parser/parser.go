package main

import (
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func grabCompanies(url string, c chan *Company, wg *sync.WaitGroup) {

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
	if next_page != "" {
		grabCompanies(BASE_URL+next_page, c, wg)
	}
}

func grabCompany(url string, c chan *Company, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println(url)
	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
	}

	name := doc.Find(".company_info .company_name a").First().Text()
	site := doc.Find(".company_info .company_site a").Text()
	about := doc.Find(".company_info .company_about").Text()
	ratingstr := doc.Find(".company_info span.rating").Text()
	address := doc.Find(".company_info .address").Text()
	link, _ := doc.Find(".company_info .company_name a").Attr("href")

	log.Println(name)
	var rating float32
	if len(ratingstr) > 0 {
		value, err := strconv.ParseFloat(ratingstr, 32)
		if err != nil {
			log.Println(err)
		} else {
			rating = float32(value)
		}
	}

	came := companiesList(doc, ".left_section .company_item")
	left := companiesList(doc, ".right_section .company_item")

	company := Company{
		name,
		site,
		about,
		rating,
		address,
		0,
		link,
		left,
		came,
	}
	c <- &company
}

func companiesList(doc *goquery.Document, path string) []Employee {
	var employees []Employee

	doc.Find(path).Each(func(i int, s *goquery.Selection) {
		page, _ := s.Find(".title").Attr("href")
		page = BASE_URL + page
		name := s.Find(".title").Text()
		amountstr := s.Find(".count").Text()
		amountstr = amountstr[:strings.Index(amountstr, " ")]

		var amount int64
		amount, err := strconv.ParseInt(amountstr, 10, 32)
		if err != nil {
			log.Println(err)
		}

		//log.Println(name, amount)
		employees = append(
			employees,
			Employee{
				CompanyPage: page,
				CompanyName: name,
				Amount:      int(amount),
			})
	})

	return employees
}
