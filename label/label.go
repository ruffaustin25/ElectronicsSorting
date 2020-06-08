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

type viewData struct {
	Part partdata.PartData
}

const templatePath string = "./label/labelTemplate.dymo"
const partParam string = "part"

var compiledTemplate *template.Template
var database *partsdatabase.PartsDatabase

// Init : Load page template
func Init(db *partsdatabase.PartsDatabase) {
	var err error
	templateBase := filepath.Base(templatePath)
	compiledTemplate, err = template.New(templateBase).Funcs(templatefunctions.GetTextFuncMap()).ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Could not load template %s, Error: %s", templatePath, err.Error())
	}
	database = db
}

// Download : get the file
func Download(res http.ResponseWriter, req *http.Request) {
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
		log.Fatalf("Could not execute template in label.go, Error %s", err.Error())
	}
}
