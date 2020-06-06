package list

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ruffaustin25/ElectronicsSorting/Common"
)

type ViewData struct {
	Parts   []Common.PartData
	StrTest string
}

const templatePath string = "./templates/list.html"

var compiledTemplate *template.Template

// Init : Load page template
func Init(layoutPath string) {
	var err error
	compiledTemplate, err = template.ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s", layoutPath, templatePath)
	}
}

// Show : Present the page
func Show(res http.ResponseWriter, req *http.Request) {
	data := ViewData{}
	data.StrTest = "lkasjdf"
	data.Parts = append(data.Parts, Common.PartData{Name: "Part name 1", Container: 0})
	data.Parts = append(data.Parts, Common.PartData{Name: "Part name 2", Container: 0})
	data.Parts = append(data.Parts, Common.PartData{Name: "Part name 3", Container: 1})
	err := compiledTemplate.Execute(res, data)
	if err != nil {
		log.Fatalf("Could not execute template in list.go")
	}
}
