package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()

	if err != nil {
		panic(err)
	}

	// request.PostFormValue("firstName") // Jika langsung

	firstName := request.PostForm.Get("firstName")
	lastName := request.PostForm.Get("lastName")

	fmt.Fprintf(writer, "%s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstName=Zulfikar&lastName=Isnaen")
	request := httptest.NewRequest("POST", "http://localhost", requestBody)

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
