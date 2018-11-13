package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var books []Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// like this until acess to db
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //Mock ID - TODO
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusCreated)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	//Mock implement db
	books = append(books, Book{"1", "1321-321", "Book One", &Author{"John", "Doe"}})
	books = append(books, Book{"2", "6667-321", "Book two", &Author{"Steve", "me"}})
	books = append(books, Book{"3", "23132-321", "Book three", &Author{"John", "Doe"}})

	// Route Handler / Endpoints
	router.HandleFunc("/api/books", GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/api/books", CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", DeleteBook).Methods("GET")

	http.ListenAndServe(":504", router)
}
