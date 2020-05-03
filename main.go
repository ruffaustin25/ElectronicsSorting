package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ruffaustin25/HouseManagement/examplepage"
	"github.com/ruffaustin25/HouseManagement/index"
	"github.com/ruffaustin25/HouseManagement/quotes"
	"github.com/ruffaustin25/HouseManagement/services"
)

const layoutPath string = "./templates/layout.html"

func main() {
	staticFS := http.FileServer(http.Dir("./static"))

	index.Init(layoutPath)
	quotes.Init(layoutPath)
	examplepage.Init(layoutPath)
	services.Init(layoutPath)

	http.HandleFunc("/", index.Show)
	http.HandleFunc("/quotes/", quotes.Show)
	http.HandleFunc("/services/", services.Show)
	http.HandleFunc("/examplepage/", examplepage.Show)
	http.Handle("/static/", http.StripPrefix("/static", staticFS))

	server := &http.Server{
		Addr:           ":1234",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
