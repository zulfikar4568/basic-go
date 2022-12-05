package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HeloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHeloHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/helo", nil)
	recorder := httptest.NewRecorder()

	HeloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
