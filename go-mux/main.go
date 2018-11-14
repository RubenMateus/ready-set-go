package main

import (
	"log"
	"net/http"

	r "github.com/rubenmateus/ready-set-go/go-mux/Router"
)

func main() {
	// Init Router
	router := r.NewRouter()

	log.Fatal(http.ListenAndServe(":504", router))
}
