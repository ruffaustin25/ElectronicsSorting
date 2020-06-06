package list

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ruffaustin25/ElectronicsSorting/common"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

type viewData struct {
	Parts []common.PartData
}

const templatePath string = "./templates/list.html"

var compiledTemplate *template.Template

// Init : Load page template
func Init(layoutPath string) {
	var err error
	layoutBase := filepath.Base(layoutPath)
	compiledTemplate, err = template.New(layoutBase).Funcs(common.GetCommonFuncMap()).ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s", layoutPath, templatePath)
	}
}

// Show : Present the page
func Show(res http.ResponseWriter, req *http.Request) {
	data := viewData{
		Parts: partsdatabase.GetPartsList(),
	}
	err := compiledTemplate.Execute(res, data)
	if err != nil {
		log.Fatalf("Could not execute template in list.go, Error %s", err.Error())
	}
}
