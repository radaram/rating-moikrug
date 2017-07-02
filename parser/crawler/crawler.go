package crawler

import (
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"

	"data"
	"parser/settings"
	"parser/producer"
)


func RunCrawler() {
 	var c chan *data.Company = make(chan *data.Company)	
	go producer.Send(c)
	grabCompanies(settings.COMPANIES_URL, c)
	time.Sleep(10 * time.Second)
}

func grabCompanies(url string, c chan *data.Company) {
	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".companies-item").Each(func(i int, s *goquery.Selection) {
		company_page, _ := s.Find(".title").Attr("href")
		company_page = settings.BASE_URL + company_page
	    go grabCompany(company_page, c)
    })

	next_page, _ := doc.Find(".pagination .next_page").Attr("href")
	log.Println(settings.BASE_URL + next_page)
	if (next_page != "") {
		grabCompanies(settings.BASE_URL + next_page, c)
	}
}


func grabCompany(url string, c chan *data.Company) {
	doc, err := goquery.NewDocument(url)

    if err != nil {
        log.Fatal(err)
    }

    name := doc.Find(".company_name a").Text()
    about := doc.Find(".company_about").Text()
    address := doc.Find(".address").Text()

    company := data.Company{name, url, about, address, 0}
    c <- &company
}

