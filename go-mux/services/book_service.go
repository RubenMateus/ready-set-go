package services

import (
	mapper "github.com/PeteProgrammer/go-automapper"
	"github.com/rubenmateus/ready-set-go/go-mux/db"
	s "github.com/rubenmateus/ready-set-go/go-mux/services/models"
)

func GetBooks() []s.Book {

	var books []s.Book

	dbBooks := db.GetBooks()
	mapper.MapLoose(dbBooks, &books)

	return books
}
