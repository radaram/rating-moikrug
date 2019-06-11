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


func (c *Company) getScoreForLeft() int {
	total := 0
	for _, employee := range c.EmployeesLeft {
		total += employee.Amount
	}
	return total
}

func (c *Company) getScoreForCame() int {
	total := 0
	for _, employee := range c.EmployeesCame {
		total += employee.Amount
	}
	return total
}

func (c *Company) getScoreForAddress() int {
	score := 0
	if len(c.Address) > 0 {
		score = 2
	}
	return score
}

func (c *Company) getScoreForSite() int {
	score := 0
	if len(c.Site) > 0 {
		score = 3
	}
	return score
}

func (c *Company) getScoreForRating() int {
	return int(c.Rating)
}

func (c *Company) calculateTotalScore() int {
	return c.getScoreForAddress() + c.getScoreForSite() + c.getScoreForRating() + c.getScoreForCame() - c.getScoreForLeft()
}


