package editcontainersubmit

import (
	"log"
	"net/http"

	"github.com/ruffaustin25/ElectronicsSorting/containerdata"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

type Page struct {
	database *partsdatabase.PartsDatabase
}

const keyParam string = "key"
const nameParam string = "name"
const descriptionParam string = "description"
const heightParam string = "height"
const widthParam string = "width"
const depthParam string = "depth"

func (p *Page) Path() string {
	return "/editcontainersubmit"
}

// Init : Load page template
func (p *Page) Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	p.database = db
}

// Show : Present the page
func (p *Page) Navigate(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Printf("Could not parse POST request to /editcontainersubmit. Error: %s", err)
		http.Redirect(res, req, "/containers=", http.StatusSeeOther)
		return
	}

	params := make(map[string]string)
	params["key"] = req.FormValue(keyParam)
	params["name"] = req.FormValue(nameParam)
	params["description"] = req.FormValue(descriptionParam)
	params["height"] = req.FormValue(heightParam)
	params["width"] = req.FormValue(widthParam)
	params["depth"] = req.FormValue(depthParam)

	container := containerdata.FromMap(params)

	p.database.UpdateContainer(container)

	http.Redirect(res, req, "/container?container="+params["key"], http.StatusSeeOther)
}
