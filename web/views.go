package main

import (
	"html/template"
	"net/http"
	"strconv"
)


func getCompanies(page int, limit int) ([]Company, error) {
	var company Company
	var raw_employees_left, raw_employees_came []byte
	
	offset := limit * (page - 1)
	rows, err := db.Query("SELECT * FROM company ORDER BY score DESC OFFSET $1 LIMIT $2", offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	companies := make([]Company, 0)
	for rows.Next() {
    	company = Company{}
		
		err = rows.Scan(
			&company.Name, 
			&company.Site, 
			&company.About, 
			&company.Rating, 
			&company.Address, 
			&company.Score,
			&company.Link,
			&raw_employees_left,
			&raw_employees_came,
			&company.ID,
		)
		failOnError(err, "read company error")
		if err != nil {
			return nil, err
		}

		company.employeesLeftDecode(raw_employees_left)
		company.employeesCameDecode(raw_employees_came)
	
		companies = append(companies, company)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return companies, nil
}


type PageData struct {
	Companies []Company
    CurrentPage int
	NextPage int
	PreviousPage int
}


func listCompaniesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	
	var page int
	pages, ok := r.URL.Query()["page"]
	if !ok || len(pages[0]) < 1 {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(pages[0])
		if err != nil {
			page = 1
		}
	}
    pageLimit := 15
	companies, err := getCompanies(page, pageLimit)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	data := PageData{}
	data.Companies = companies
	data.CurrentPage = page
	data.NextPage = page + 1
	data.PreviousPage = page - 1
	tmpl.Execute(w, data)
}
