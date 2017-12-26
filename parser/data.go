package main


import (
	"encoding/json"
	"fmt"
)

type company struct {
	Name 	string
	Page  	string
	About 	string
	Address string
	Score 	uint16
}


func (c *company) ToJson() []byte {
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
