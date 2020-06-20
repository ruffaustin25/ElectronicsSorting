package endpoint

import (
	"net/http"

	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

type Endpoint interface {
	Path() string
	Init(layoutPath string, db *partsdatabase.PartsDatabase)
	Navigate(res http.ResponseWriter, req *http.Request)
}
