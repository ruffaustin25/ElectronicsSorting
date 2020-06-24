package label

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/ruffaustin25/ElectronicsSorting/partdata"
	"github.com/ruffaustin25/ElectronicsSorting/templatefunctions"

	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

type Page struct {
	compiledTemplate *template.Template
	database         *partsdatabase.PartsDatabase
}

type viewData struct {
	Part partdata.PartData
}

const templatePath string = "./label/labelTemplate.dymo"
const keyParam string = "key"

func (p *Page) Path() string {
	return "/label"
}

// Init : Load page template
func (p *Page) Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	var err error
	templateBase := filepath.Base(templatePath)
	p.compiledTemplate, err = template.New(templateBase).Funcs(templatefunctions.GetTextFuncMap()).ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Could not load template %s, Error: %s", templatePath, err.Error())
	}
	p.database = db
}

// Download : get the file
func (p *Page) Navigate(res http.ResponseWriter, req *http.Request) {
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

	err := p.compiledTemplate.Execute(res, data)
	if err != nil {
		log.Fatalf("Could not execute template in label.go, Error %s", err.Error())
	}
}
