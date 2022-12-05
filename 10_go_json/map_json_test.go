package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonRequest := `{"id": "123", "name": "Mac", "price": 8000}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "123",
		"name":  "Mac",
		"price": 8000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
