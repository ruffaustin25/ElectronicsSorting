package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ruffaustin25/ElectronicsSorting/index"
	"github.com/ruffaustin25/ElectronicsSorting/label"
	"github.com/ruffaustin25/ElectronicsSorting/list"
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

	http.HandleFunc("/", index.Show)
	http.HandleFunc("/list", list.Show)
	http.HandleFunc("/part", part.Show)
	http.HandleFunc("/label", label.Download)
	http.Handle("/static/", http.StripPrefix("/static", staticFS))

	server := &http.Server{
		Addr:           ":2796",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
