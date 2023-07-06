package errorpage

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
	"github.com/ruffaustin25/ElectronicsSorting/templatefunctions"
)

type Page struct {
	compiledTemplate *template.Template
}

type viewData struct {
	ErrorMessage string
	Destination  string
}

const templatePath string = "./templates/errorpage.gohtml"
const messageParam string = "message"
const destinationParam string = "destination"

func (p *Page) Path() string {
	return "/error"
}

// Init : Load page template
func (p *Page) Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	var err error
	layoutBase := filepath.Base(layoutPath)
	p.compiledTemplate, err = template.New(layoutBase).Funcs(templatefunctions.GetHTMLFuncMap()).ParseFiles(layoutPath, templatePath)
	if err != nil {
		log.Fatalf("Could not load layout %s or template %s", layoutPath, templatePath)
	}
}

// Show : Present the page
func (p *Page) Navigate(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	message := params[messageParam]
	if len(message) == 0 {
		log.Print("No message value sent as get query")
		message = []string{"Generic Error"}
	}

	destination := params[destinationParam]
	if len(destination) == 0 {
		log.Print("No destination value sent as get query")
		destination = []string{"/list"}
	}

	data := viewData{
		ErrorMessage: message[0],
		Destination:  destination[0],
	}

	err := p.compiledTemplate.Execute(res, data)
	if err != nil {
		log.Fatalf("Could not execute template in error.go, Error %s", err.Error())
	}
}

func DoErrorPage(res http.ResponseWriter, req *http.Request, message string, destination string) {
	if len(message) > 1024 {
		message = message[:1024]
	}
	http.Redirect(res, req, "/error?"+messageParam+"="+message+"&"+destinationParam+"="+destination, http.StatusSeeOther)
}
