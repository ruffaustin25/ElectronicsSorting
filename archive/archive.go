package archive

import (
	"log"
	"net/http"

	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

const keyParam string = "key"

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

	database.ArchivePart(key[0])

	http.Redirect(res, req, "/list", http.StatusSeeOther)
}
