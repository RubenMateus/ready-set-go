package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	mapper "github.com/PeteProgrammer/go-automapper"
	"github.com/gorilla/mux"
	"github.com/rubenmateus/ready-set-go/go-mux/services"
	. "github.com/rubenmateus/ready-set-go/go-mux/web/models"
)

var books = []Book{
	// Mocked
	Book{"1", "1321-321", "Book One", &Author{"John", "Doe"}},
	Book{"2", "6667-321", "Book two", &Author{"Steve", "me"}},
	Book{"3", "23132-321", "Book three", &Author{"John", "Doe"}},
}

func GetBooks(w http.ResponseWriter, r *http.Request) {

	var books []Book
	serviceBooks := services.GetBooks()
	mapper.MapLoose(serviceBooks, &books)

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
			w.WriteHeader(http.StatusAccepted)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}
