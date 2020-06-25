package newpart

import (
	"net/http"

	"github.com/ruffaustin25/ElectronicsSorting/errorpage"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

type Page struct {
	database *partsdatabase.PartsDatabase
}

const keyParam string = "key"
const nameParam string = "name"

func (p *Page) Path() string {
	return "/newpart"
}

// Init : Load page template
func (p *Page) Init(layoutPath string, db *partsdatabase.PartsDatabase) {
	p.database = db
}

// Show : Present the page
func (p *Page) Navigate(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	key := params[keyParam]
	if len(key) == 0 {
		errorpage.DoErrorPage(res, req, "No key sent as get query", "/list")
		return
	}

	name := params[nameParam]
	if len(name) == 0 {
		errorpage.DoErrorPage(res, req, "No name sent as get query", "/list")
		return
	}

	err := p.database.CreatePart(key[0], name[0])
	if err != nil {
		errorpage.DoErrorPage(res, req, err.Error(), "/list")
		return
	}

	http.Redirect(res, req, "/part?part="+key[0], http.StatusSeeOther)
}
