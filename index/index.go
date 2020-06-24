package index

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

type Page struct {
	compiledTemplate *template.Template
}

type viewData struct {
}

const templatePath string = "./templates/index.gohtml"

func (p *Page) Path() string {
	return "/"
}

// Load page template
func (p *Page) Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	var err error
	p.compiledTemplate, err = template.ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s", layoutPath, templatePath)
	}
}

// Present the page
func (p *Page) Navigate(res http.ResponseWriter, req *http.Request) {
	data := viewData{}
	p.compiledTemplate.Execute(res, data)
}
