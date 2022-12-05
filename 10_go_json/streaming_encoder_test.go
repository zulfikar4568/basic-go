package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("sample_output.json")

	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Zulfikar",
		LastName:  "Isnaen",
		Age:       24,
	}

	_ = encoder.Encode(customer)
	fmt.Println(customer)
}
