package container

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ruffaustin25/ElectronicsSorting/containerdata"
	"github.com/ruffaustin25/ElectronicsSorting/partdata"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
	"github.com/ruffaustin25/ElectronicsSorting/templatefunctions"
)

type Page struct {
	compiledTemplate *template.Template
	database         *partsdatabase.PartsDatabase
}

type viewData struct {
	Container containerdata.ContainerData
	Parts     [][][]*partdata.PartData
}

const templatePath string = "./templates/container.gohtml"
const containerParam string = "container"

func (p *Page) Path() string {
	return "/container"
}

// Init : Load page template
func (p *Page) Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	var err error
	layoutBase := filepath.Base(layoutPath)
	p.compiledTemplate, err = template.New(layoutBase).Funcs(templatefunctions.GetHTMLFuncMap()).ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s", layoutPath, templatePath)
	}
	p.database = db
}

// Show : Present the page
func (p *Page) Navigate(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	containerKey := params[containerParam]
	if len(containerKey) == 0 {
		log.Print("No container value sent as get query")
		return
	}

	container := p.database.GetContainer(containerKey[0])
	if container == nil {
		log.Printf("No container found for %s", containerKey)
		return
	}
	data := viewData{
		Container: *container,
		Parts:     p.database.GetPartsInContainer(container),
	}

	err := p.compiledTemplate.Execute(res, data)
	if err != nil {
		log.Fatalf("Could not execute template in container.go, Error %s", err.Error())
	}
}
