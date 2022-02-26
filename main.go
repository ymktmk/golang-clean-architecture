package main

import (
	"Golang-CleanArchitecture/infrastructure"
	"log"
	"net/http"
	"github.com/rs/cors"
)

func main() {
	// routing
	r := infrastructure.Routing()
	// cors enabled
	c := cors.AllowAll().Handler(r)
	// start the server
	log.Fatal(http.ListenAndServe(":8080", c))
}