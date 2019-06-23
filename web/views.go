package main

import (
	"html/template"
	"net/http"
)


func getCompanies() ([]Company, error) {
	var company Company
	var raw_employees_left, raw_employees_came []byte

	rows, err := db.Query("SELECT * FROM company LIMIT 15")
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


func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	companies, err := getCompanies()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	tmpl.Execute(w, companies)
}
