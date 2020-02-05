package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Main Page</h1>")
}

func main() {
	http.HandleFunc("/", index)

	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
