package main

import (
	"log"
	"net/http"
)

func main() {
	// Init Router
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":504", router))
}
