package main

import (
	"log"
	"net/http"

	router "github.com/rubenmateus/ready-set-go/go-mux/web/router"
)

func main() {
	// Init Router
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":504", router))
}
