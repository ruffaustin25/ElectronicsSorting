package editpartsubmit

import (
	"log"
	"net/http"

	"github.com/ruffaustin25/ElectronicsSorting/partdata"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

type Page struct {
	database *partsdatabase.PartsDatabase
}

const keyParam string = "key"
const nameParam string = "name"
const descriptionParam string = "description"
const containerParam string = "container"
const rowParam string = "row"
const columnParam string = "column"
const depthParam string = "depth"

func (p Page) Path() string {
	return "/editpartsubmit"
}

// Init : Load page template
func (p Page) Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	p.database = db
}

// Show : Present the page
func (p Page) Navigate(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Printf("Could not parse POST request to /editpartsubmit. Error: %s", err)
		http.Redirect(res, req, "/list=", http.StatusSeeOther)
		return
	}

	params := make(map[string]string)
	params["key"] = req.FormValue(keyParam)
	params["name"] = req.FormValue(nameParam)
	params["description"] = req.FormValue(descriptionParam)
	params["container"] = req.FormValue(containerParam)
	params["row"] = req.FormValue(rowParam)
	params["column"] = req.FormValue(columnParam)
	params["depth"] = req.FormValue(depthParam)

	part := partdata.FromMap(params)

	p.database.UpdatePart(part)

	http.Redirect(res, req, "/part?part="+params["key"], http.StatusSeeOther)
}
