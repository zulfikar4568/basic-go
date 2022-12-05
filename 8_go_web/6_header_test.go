package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Zulfikar Isnaen")
	fmt.Fprint(writer, "OK")
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	poweredBy := recorder.Header().Get("x-powered-by")
	fmt.Println(string(poweredBy))
}
