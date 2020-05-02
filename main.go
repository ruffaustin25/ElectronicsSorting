package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const indexTemplatePath string = "./templates/index.html"

var indexTemplateText []byte

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, string(indexTemplateText))
}

func main() {
	fmt.Println("running server")

	file, err := os.Open(indexTemplatePath)
	if err != nil {
		log.Fatal(err)
	}

	indexTemplateText, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", index)

	server := &http.Server{
		Addr:           ":1234",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
