package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	CompanyPage string
	CompanyName string
	Amount      int
}

type Company struct {
	Name          string
	Site          string
	About         string
	Rating        float32
	Address       string
	Score         int
	Link          string
	EmployeesLeft []Employee
	EmployeesCame []Employee
}

func (c *Company) ToJson() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
