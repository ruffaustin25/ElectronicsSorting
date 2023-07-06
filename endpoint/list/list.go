package list

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ruffaustin25/ElectronicsSorting/partdata"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
	"github.com/ruffaustin25/ElectronicsSorting/templatefunctions"
)

type Page struct {
	compiledTemplate *template.Template
	database         *partsdatabase.PartsDatabase
}

type viewData struct {
	Parts []partdata.PartData
}

const templatePath string = "./templates/list.gohtml"

func (p *Page) Path() string {
	return "/list"
}

// Init : Load page template
func (p *Page) Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	var err error
	layoutBase := filepath.Base(layoutPath)
	p.compiledTemplate, err = template.New(layoutBase).Funcs(templatefunctions.GetHTMLFuncMap()).ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s, %s", layoutPath, templatePath, err)
	}
	p.database = db
}

// Navigate : Present the page
func (p *Page) Navigate(res http.ResponseWriter, req *http.Request) {
	data := viewData{
		Parts: p.database.GetPartsList(),
	}
	err := p.compiledTemplate.Execute(res, data)
	if err != nil {
		log.Fatalf("Could not execute template in list.go, Error %s", err.Error())
	}
}
