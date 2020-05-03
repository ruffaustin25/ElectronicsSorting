package examplepage

import (
	"html/template"
	"log"
	"net/http"
)

type viewData struct {
}

const templatePath string = "./templates/examplepage.html"

var compiledTemplate *template.Template

// Load page template
func Init(layoutPath string) {
	var err error
	compiledTemplate, err = template.ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s", layoutPath, templatePath)
	}
}

// Present the page
func Show(res http.ResponseWriter, req *http.Request) {
	data := viewData{}
	compiledTemplate.Execute(res, data)
}
