package editcontainer

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ruffaustin25/ElectronicsSorting/containerdata"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

type Page struct {
	compiledTemplate *template.Template
	database         *partsdatabase.PartsDatabase
}

type viewData struct {
	Container containerdata.ContainerData
}

const templatePath string = "./templates/editcontainer.gohtml"
const keyParam string = "key"

func (p *Page) Path() string {
	return "/editcontainer"
}

// Init : Load page template
func (p *Page) Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	var err error
	p.compiledTemplate, err = template.ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s", layoutPath, templatePath)
	}
	p.database = db
}

// Show : Present the page
func (p *Page) Navigate(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	keys := params[keyParam]
	if len(keys) == 0 {
		log.Print("No container value sent as get query")
		return
	}

	container := p.database.GetContainer(keys[0])
	if container == nil {
		log.Printf("No container found for %s", keys[0])
		return
	}
	data := viewData{
		Container: *container,
	}
	p.compiledTemplate.Execute(res, data)
}
