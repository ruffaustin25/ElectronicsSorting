package list

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ruffaustin25/ElectronicsSorting/common"
	"github.com/ruffaustin25/ElectronicsSorting/partdata"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

type viewData struct {
	Parts []partdata.PartData
}

const templatePath string = "./templates/list.gohtml"

var compiledTemplate *template.Template
var database *partsdatabase.PartsDatabase

// Init : Load page template
func Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	var err error
	layoutBase := filepath.Base(layoutPath)
	compiledTemplate, err = template.New(layoutBase).Funcs(common.GetCommonFuncMap()).ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s", layoutPath, templatePath)
	}
	database = db
}

// Show : Present the page
func Show(res http.ResponseWriter, req *http.Request) {
	data := viewData{
		Parts: database.Parts,
	}
	err := compiledTemplate.Execute(res, data)
	if err != nil {
		log.Fatalf("Could not execute template in list.go, Error %s", err.Error())
	}
}
