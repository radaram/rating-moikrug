package data


import (
	"encoding/json"
	"fmt"
)

type Company struct {
	Name 	string
	Page  	string
	About 	string
	Address string
	Score 	uint16
}


func (c *Company) ToJson() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
