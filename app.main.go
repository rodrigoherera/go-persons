package main

import (
	"go-persons/db"
	rr "go-persons/routes"
	"log"
	"net/http"
)

func main() {
	db.CreateMariaDB()
	srv := &http.Server{
		Handler: rr.GetRouter(),
		Addr:    ":8080",
	}
	log.Fatal(srv.ListenAndServe())
}
