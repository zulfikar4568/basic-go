package goweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

//go:embed resources/index.html
var resourceOK string

//go:embed resources/not-found.html
var resourceEmpty string

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/index.html")
	} else {
		http.ServeFile(writer, request, "./resources/not-found.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Println(resourceOK, resourceEmpty)
		fmt.Fprint(writer, resourceOK)
	} else {
		fmt.Fprint(writer, resourceEmpty)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
