package newpart

import (
	"log"
	"net/http"

	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

const keyParam string = "key"
const nameParam string = "name"

var database *partsdatabase.PartsDatabase

// Init : Load page template
func Init(db *partsdatabase.PartsDatabase) {
	database = db
}

// Show : Present the page
func Show(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	key := params[keyParam]
	if len(key) == 0 {
		log.Print("No key sent as get query")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	name := params[nameParam]
	if len(name) == 0 {
		log.Print("No name sent as get query")
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	database.CreatePart(key[0], name[0])

	http.Redirect(res, req, "/part?part="+key[0], http.StatusSeeOther)
}
