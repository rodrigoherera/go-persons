package main

import (
	"fmt"
	"go-persons/db"
	rr "go-persons/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	httpPORT := os.Getenv("HTTP_PORT")

	fmt.Println("MariaDB : database initializing...")
	db.CreateMariaDB()
	fmt.Println("MariaDB : database is ready!")

	srv := &http.Server{
		Handler: rr.GetRouter(),
		Addr:    ":" + httpPORT,
	}
	fmt.Printf("HTTP : HTTP is ready on %v", httpPORT)
	log.Fatal(srv.ListenAndServe())
}
