package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ruffaustin25/ElectronicsSorting/index"
	"github.com/ruffaustin25/ElectronicsSorting/list"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

const layoutPath string = "./templates/layout.html"

func main() {
	staticFS := http.FileServer(http.Dir("./static"))

	db := partsdatabase.NewPartsDatabase()

	index.Init(layoutPath)
	list.Init(layoutPath, db)

	http.HandleFunc("/", index.Show)
	http.HandleFunc("/list", list.Show)
	http.Handle("/static/", http.StripPrefix("/static", staticFS))

	server := &http.Server{
		Addr:           ":2796",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
