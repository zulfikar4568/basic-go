package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")

	if first_name == "" && last_name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s %s", first_name, last_name)
	}
}

func MultipleParameter(writer http.ResponseWriter, request *http.Request) {
	var query url.Values = request.URL.Query()

	var names []string = query["name"]
	fmt.Fprintln(writer, strings.Join(names, ","))
}

func TestMultiParamQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Budi&name=Zul", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestParamQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Zulfikar&last_name=Isnaen", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
