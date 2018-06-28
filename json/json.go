package main

import (
	"encoding/json"
	"fmt"
)

type plugin int

type company struct {
	Address   string `json:"address"`
	Name      string `json:"name"`
	NIPNumber int    `json:"nip-number"`
	Phone     string `json:"phone-number"`
}

const exampleJson = `
	{
		"address": "Test Address",
		"name": "Test name",
		"nip-number": 1234567890
	}
`

func (p plugin) Example() {
	var company company
	err := json.Unmarshal([]byte(exampleJson), &company)

	if err != nil {
		panic(err)
	}

	fmt.Println(company)

	company.Phone = "+48 500 600 700"

	jsonCompany, err := json.Marshal(company)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", jsonCompany)
}

var Plugin plugin
