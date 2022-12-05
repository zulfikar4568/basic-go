package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJsonPrint(t *testing.T) {
	logJson("Zulfikar")
	logJson(1)
	logJson(true)
	logJson([]string{"Buku", "Laptop", "HP"})
}

type Address struct {
	Street     string
	City       string
	PostalCode string
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Address   []Address
	Hobies    []string
}

func TestEncodeJSON(t *testing.T) {
	customer := Customer{
		FirstName: "Zulfikar",
		LastName:  "Isnaen",
		Age:       22,
		Address: []Address{
			{
				Street:     "Kebon Kopi",
				City:       "Bandung",
				PostalCode: "40535",
			},
			{
				Street:     "Jl Jendral Sudirman",
				City:       "Jakarta",
				PostalCode: "1234",
			},
		},
		Hobies: []string{"Footbal", "Drawing"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
