package part

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ruffaustin25/ElectronicsSorting/partdata"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
	"github.com/ruffaustin25/ElectronicsSorting/templatefunctions"
)

type viewData struct {
	Part partdata.PartData
}

const templatePath string = "./templates/part.gohtml"
const partParam string = "part"

var compiledTemplate *template.Template
var database *partsdatabase.PartsDatabase

// Init : Load page template
func Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	var err error
	layoutBase := filepath.Base(layoutPath)
	compiledTemplate, err = template.New(layoutBase).Funcs(templatefunctions.GetHTMLFuncMap()).ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s", layoutPath, templatePath)
	}
	database = db
}

// Show : Present the page
func Show(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	partValue := params[partParam]
	if len(partValue) == 0 {
		log.Print("No part value sent as get query")
		return
	}

	part := database.GetPart(partValue[0])
	if part == nil {
		log.Printf("No part found for %s", partValue)
		return
	}
	data := viewData{
		Part: *part,
	}

	err := compiledTemplate.Execute(res, data)
	if err != nil {
		log.Fatalf("Could not execute template in part.go, Error %s", err.Error())
	}
}
