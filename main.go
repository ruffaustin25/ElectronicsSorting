package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ruffaustin25/ElectronicsSorting/archive"
	"github.com/ruffaustin25/ElectronicsSorting/editpart"
	"github.com/ruffaustin25/ElectronicsSorting/editpartsubmit"
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

	index.Init(layoutPath)
	list.Init(layoutPath, db)
	part.Init(layoutPath, db)
	label.Init(db)
	newpart.Init(db)
	archive.Init(db)
	editpart.Init(layoutPath, db)
	editpartsubmit.Init(db)

	http.HandleFunc("/", index.Show)
	http.HandleFunc("/list", list.Show)
	http.HandleFunc("/part", part.Show)
	http.HandleFunc("/label", label.Download)
	http.HandleFunc("/newpart", newpart.Show)
	http.HandleFunc("/archive", archive.Show)
	http.HandleFunc("/editpart", editpart.Show)
	http.HandleFunc("/editpartsubmit", editpartsubmit.Show)
	http.Handle("/static/", http.StripPrefix("/static", staticFS))

	server := &http.Server{
		Addr:           ":2796",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
