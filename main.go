package main

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/pion/mdns"
	"golang.org/x/net/ipv4"

	"github.com/ruffaustin25/ElectronicsSorting/buildconfig"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/archive"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/container"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/containers"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/editcontainer"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/editcontainersubmit"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/editpart"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/editpartsubmit"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/errorpage"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/index"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/label"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/list"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/newcontainer"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/newpart"
	"github.com/ruffaustin25/ElectronicsSorting/endpoint/part"
	"github.com/ruffaustin25/ElectronicsSorting/partsdatabase"
)

const layoutPath string = "./templates/layout.gohtml"

func main() {
	staticFS := http.FileServer(http.Dir("./static"))

	db := partsdatabase.NewPartsDatabase()

	endpoints := []endpoint.Endpoint{
		&index.Page{},
		&list.Page{},
		&part.Page{},
		&label.Page{},
		&newpart.Page{},
		&archive.Page{},
		&editpart.Page{},
		&editpartsubmit.Page{},
		&containers.Page{},
		&container.Page{},
		&newcontainer.Page{},
		&editcontainer.Page{},
		&editcontainersubmit.Page{},
		&errorpage.Page{},
	}

	for _, endpoint := range endpoints {
		endpoint.Init(layoutPath, db)
		http.HandleFunc(endpoint.Path(), endpoint.Navigate)
	}

	http.Handle("/static/", http.StripPrefix("/static", staticFS))

	server := &http.Server{
		Addr:           ":" + strconv.Itoa(buildconfig.AppPort),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	addr, err := net.ResolveUDPAddr("udp", mdns.DefaultAddress)
	if err != nil {
		panic(err)
	}

	l, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(err)
	}

	_, err = mdns.Server(ipv4.NewPacketConn(l), &mdns.Config{
		LocalNames: []string{buildconfig.BaseURL},
	})
	if err != nil {
		panic(err)
	}

	log.Fatal(server.ListenAndServe())
}
