package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageURL string `json:"image_url"`
}

func TestJSONTagEncode(t *testing.T) {
	product := Product{
		Id:       "1",
		Name:     "Zulfikar",
		Price:    "8000",
		ImageURL: "www.gambar.com",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonRequest := `{"id":"1","name":"Zulfikar","price":"8000","image_url":"www.gambar.com"}`
	jsonBytes := []byte(jsonRequest)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
