package goweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	t, err := template.New("SIMPLE").Parse(templateText)

	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "index.gohtml", "Hello HTML Template")
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "index.gohtml", "Hello HTML Template")
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(writer, "index.gohtml", "Hello HTML Template")
}

func TestSimpleTemplateHTML(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestSimpleTemplateHTMLFile(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
