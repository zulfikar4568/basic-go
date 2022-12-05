package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonRequest := `{"FirstName":"Zulfikar","LastName":"Isnaen","Age":22,"Address":[{"Street":"Kebon Kopi","City":"Bandung","PostalCode":"40535"},{"Street":"Jl Jendral Sudirman","City":"Jakarta","PostalCode":"1234"}],"Hobies":["Footbal","Drawing"]}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestDecodeJSONSlice(t *testing.T) {
	jsonRequest := `[{"Street":"Kebon Kopi","City":"Bandung","PostalCode":"40535"},{"Street":"Jl Jendral Sudirman","City":"Jakarta","PostalCode":"1234"}]`
	jsonBytes := []byte(jsonRequest)

	address := &[]Address{}
	err := json.Unmarshal(jsonBytes, address)
	if err != nil {
		panic(err)
	}

	fmt.Println(address)
}
