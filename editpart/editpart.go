package editpart

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ruffaustin25/ElectronicsSorting/partdata"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

type Page struct {
	compiledTemplate *template.Template
	database         *partsdatabase.PartsDatabase
}

type viewData struct {
	Part partdata.PartData
}

const templatePath string = "./templates/editpart.gohtml"
const keyParam string = "key"

func (p Page) Path() string {
	return "/editpart"
}

// Init : Load page template
func (p Page) Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	var err error
	p.compiledTemplate, err = template.ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s", layoutPath, templatePath)
	}
	p.database = db
}

// Show : Present the page
func (p Page) Navigate(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	keys := params[keyParam]
	if len(keys) == 0 {
		log.Print("No part value sent as get query")
		return
	}

	part := p.database.GetPart(keys[0])
	if part == nil {
		log.Printf("No part found for %s", keys[0])
		return
	}
	data := viewData{
		Part: *part,
	}
	p.compiledTemplate.Execute(res, data)
}
