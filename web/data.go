package main


import (
	"encoding/json"
)


type Employee struct {
	CompanyPage	string
	CompanyName string
	Amount int
}


type Company struct {
	Name 	string
	Site  	string
	About 	string
	Rating  float32
	Address string
	Score 	int
	EmployeesLeft []Employee
	EmployeesCame []Employee
	ID      int
}


func (c *Company) Decode(data []byte) error {
    if err := json.Unmarshal(data, &c); err != nil {
        return err
    }
	return nil
}

func (c *Company) employeesLeftJsonEncode() ([]byte, error) {
	return json.Marshal(c.EmployeesLeft)
}

func (c *Company) employeesCameJsonEncode() ([]byte, error) {
	return json.Marshal(c.EmployeesCame)
}


