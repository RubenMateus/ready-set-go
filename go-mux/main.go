package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	router := mux.NewRouter()

	// Route Handler / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("GET")

	http.ListenAndServe(":504", router)
}
