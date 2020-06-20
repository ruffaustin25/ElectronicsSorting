package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ruffaustin25/ElectronicsSorting/archive"
	"github.com/ruffaustin25/ElectronicsSorting/editpart"
	"github.com/ruffaustin25/ElectronicsSorting/editpartsubmit"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint"
	"github.com/ruffaustin25/ElectronicsSorting/index"
	"github.com/ruffaustin25/ElectronicsSorting/label"
	"github.com/ruffaustin25/ElectronicsSorting/list"
	"github.com/ruffaustin25/ElectronicsSorting/newpart"
	"github.com/ruffaustin25/ElectronicsSorting/part"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

const layoutPath string = "./templates/layout.gohtml"

func main() {
	staticFS := http.FileServer(http.Dir("./static"))

	db := partsdatabase.NewPartsDatabase()

	endpoints := []endpoint.Endpoint{
		index.Page{},
		list.Page{},
		part.Page{},
		label.Page{},
		newpart.Page{},
		archive.Page{},
		editpart.Page{},
		editpartsubmit.Page{},
	}

	for _, endpoint := range endpoints {
		endpoint.Init(layoutPath, db)
		http.HandleFunc(endpoint.Path(), endpoint.Navigate)
	}

	http.Handle("/static/", http.StripPrefix("/static", staticFS))

	server := &http.Server{
		Addr:           ":2796",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
