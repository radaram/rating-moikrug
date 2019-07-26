package main

import (
	"database/sql"
	"log"
)

func calculate(company Company) {
	log.Println(company)
	total_score := company.calculateTotalScore()
	log.Println("total_score", total_score)
	company.Score = total_score
	saveResult(company)
}

func saveResult(company Company) {
	var company_id int
	employeesLeftData, err := company.employeesLeftJsonEncode()
	failOnError(err, "json encoded error")

	employeesCameData, err := company.employeesCameJsonEncode()
	failOnError(err, "json encoded error")

	err = db.QueryRow("SELECT id FROM company WHERE name = $1", company.Name).Scan(&company_id)
	if err == sql.ErrNoRows {
		_, err = db.Exec(
			"INSERT INTO company (name, site, about, rating, address, score, link, employees_left, employees_came) "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
			company.Name, company.Site,
			company.About, company.Rating,
			company.Address, company.Score,
			company.Link,
			string(employeesLeftData), string(employeesCameData),
		)
		failOnError(err, "insert error")
		log.Println("INSERT", company.Name)
	} else if err != nil {
		failOnError(err, "get company error")
	} else {
		_, err = db.Exec("UPDATE company "+
			"SET site = $1, about = $2, rating = $3, address = $4, "+
			"    score = $5, link = $6, employee_left = $7, employee_came = $8 "+
			"WHERE id = $8",
			company.Site, company.About,
			company.Rating, company.Address, company.Score,
			company.Link,
			string(employeesLeftData), string(employeesCameData),
			company_id)
		failOnError(err, "update error")
		log.Println("UPDATE", company.Name)
	}
}
