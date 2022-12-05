package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type User struct {
	Name string
	Age  int
}

type Page struct {
	Title string
	Name  string
	User  User
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/name2.gohtml")
	if err != nil {
		panic(err)
	}
	// t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
	// 	"Title": "Test Title",
	// 	"Name":  "Test data Name",
	// })
	t.ExecuteTemplate(writer, "name2.gohtml", Page{
		Title: "Test Title",
		Name:  "Test Data Name",
		User: User{
			Age:  22,
			Name: "Zulfikar",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
